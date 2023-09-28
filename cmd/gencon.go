/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"datagen/appconfig"
	"datagen/config"
	"datagen/types"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// genconCmd represents the gencon command
var genconCmd = &cobra.Command{
	Use:   "genconf",
	Short: "Generate default config for a given table in database",
	Long:  "Generate default config for a given table in database either in JSON or YAML format",
	Run: func(cmd *cobra.Command, args []string) {

		debugEnabled, _ := cmd.Flags().GetBool("debug")

		if debugEnabled {
			log.SetLevel(log.DebugLevel)
			log.SetReportCaller(true)
		}

		log.Info("generate configuration invoked")
		tblname, _ := cmd.Flags().GetString("tblname")
		conname, _ := cmd.Flags().GetString("conname")
		// configFormat, _ := cmd.Flags().GetString("format")
		outFile, _ := cmd.Flags().GetString("out")
		aconfig := appconfig.GetConf()

		log.Infof("About to proble table : %s\n", tblname)
		var myTableInfo types.Table
		myTableInfo, _ = config.ProbeTableMetadata(conname, tblname)
		var result string
		switch aconfig.ConfigFormat {
		case "YAML", "yaml":
			yamlOut, _ := yaml.Marshal(myTableInfo)
			result = string(yamlOut)

		case "JSON", "json":
			foo_marshalled, _ := json.Marshal(myTableInfo)
			result, _ = config.PrettyString(string(foo_marshalled))

		default:
			fmt.Println("Invalid Option selected")
		}

		log.Infof("output file: %s", outFile)
		if outFile != "" {
			outFilePath := filepath.Join(aconfig.OutputDir, outFile)
			log.Infof("configuration will be written to %s", outFilePath)
			config.WriteContent(outFilePath, result)
		} else {
			log.Info("configuration will be written to console")
			fmt.Println(result)
		}

		log.Info("Complete!")
	},
}

func init() {
	rootCmd.AddCommand(genconCmd)
	//genconCmd.Flags().String("format", "JSON", "Config format to use JSON or YAML")
	genconCmd.Flags().String("out", "", "output file name")
}
