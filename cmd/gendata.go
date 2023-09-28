/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"datagen/appconfig"
	"datagen/config"
	"datagen/dbutils"
	"datagen/queryhelper"
	"datagen/types"
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// gendataCmd represents the gendata command
var gendataCmd = &cobra.Command{
	Use:   "gendata",
	Short: "Command to generate table as per provided configuration file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		debugEnabled, _ := cmd.Flags().GetBool("debug")

		if debugEnabled {
			log.SetLevel(log.DebugLevel)
			log.SetReportCaller(true)
		}

		log.Info("generate data invoked")
		aconfig := appconfig.GetConf()
		conname, _ := cmd.Flags().GetString("conname")
		tblname, _ := cmd.Flags().GetString("tblname")
		rcount, _ := cmd.Flags().GetInt("rcount")

		if rcount <= 0 {
			log.Fatal("rcount is zero, no need to proceed. exiting")
		}

		configfile, _ := cmd.Flags().GetString("configfile")
		configfileExtn := path.Ext(configfile)
		outFile, _ := cmd.Flags().GetString("outfile")
		inline, _ := cmd.Flags().GetBool("inline")

		batchSize := aconfig.BatchSize
		tuplesSize := aconfig.TupleSize
		//configFormat := aconfig.ConfigFormat

		var tconfig types.Table

		if strings.ToLower(configfileExtn) == "json" {
			tconfig, _ = config.ParseJsonFile(configfile)
		} else if strings.ToLower(configfileExtn) == "yaml" {
			tconfig, _ = config.ParseYamlFile(configfile)
		} else {
			log.Fatal("Invalid config format. Valid ones are [ 'json', 'yaml']")
		}

		log.Infof("About to generate insert statements for table: %s", tblname)
		tconfig = config.Enrich(conname, tconfig)
		recordsToGenerate := rcount
		totalContent := []string{}
		for recordsToGenerate > 0 {
			var recordItem string
			if recordsToGenerate > tuplesSize {
				recordItem, _ = queryhelper.GenerateFullInsertStatement(tconfig, tuplesSize)
				recordsToGenerate -= tuplesSize
			} else {
				recordItem, _ = queryhelper.GenerateFullInsertStatement(tconfig, recordsToGenerate)
				recordsToGenerate = 0
			}
			totalContent = append(totalContent, recordItem)
		}

		log.Infof("generated %d insert statements with max tupes %d in one statement.", len(totalContent), tuplesSize)

		if inline {
			dbutils.ExecuteInTransaction(conname, totalContent, batchSize)
		} else if !inline && outFile != "" {
			log.Info("about to write output file")
			config.WriteContent(outFile, strings.Join(totalContent, ";\n")+";\n")
		} else {
			log.Info(strings.Join(totalContent, ";\n"))
		}

		log.Infof("Completed inserting total rows %d", rcount)
	},
}

func init() {
	rootCmd.AddCommand(gendataCmd)
	gendataCmd.Flags().Int("rcount", 0, "Number of records to generate")
	gendataCmd.Flags().String("configfile", "", "input file name")
	gendataCmd.Flags().String("outfile", "", "output file name")
	gendataCmd.Flags().Bool("inline", false, "load directly into database")
}
