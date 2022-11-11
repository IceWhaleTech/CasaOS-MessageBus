/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// registerEventTypeCmd represents the registerEventType command
var registerEventTypeCmd = &cobra.Command{
	Use:   "event-type",
	Short: "register an event type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("registerEventType called")
	},
}

func init() {
	registerCmd.AddCommand(registerEventTypeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerEventTypeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerEventTypeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
