/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "datagen",
	Short: "A cmd line utility to generate mock data for database table",
	Long:  "A cmd line utility to generate mock data for database table",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("conname", "", "Database Connection Name in datagen.json")
	rootCmd.PersistentFlags().String("tblname", "", "Table name to use for the action (generate config | generate data)")
	rootCmd.PersistentFlags().Bool("debug", false, "enable debug logs")
}
