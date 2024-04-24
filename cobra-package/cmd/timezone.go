/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// timezoneCmd represents the timezone command
var timezoneCmd = &cobra.Command{
	Use:   "timezone",
	Short: "Get the current time in given timezone",
	Long: `get the current time in given timezone:
given a timezone, this command will return the current time in that timezone.`,
	//Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timezone := args[0]
		dateFlag, _ := cmd.Flags().GetString("date")
		location, _ := time.LoadLocation(timezone)
		var date string

		if dateFlag != "" {
			date = time.Now().In(location).Format(dateFlag)
		} else {
			date = time.Now().In(location).Format(time.RFC3339)
		}

		fmt.Printf("Current date and time in %v is %v\n", timezone, date)
	},
}

func init() {
	rootCmd.AddCommand(timezoneCmd)
	rootCmd.PersistentFlags().StringP("date", "-d", "", "The timezone to get the current time in (format: (yyy-mm-dd)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timezoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timezoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
