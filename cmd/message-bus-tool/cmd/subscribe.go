/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/net/websocket"
)

const origin = "http://localhost/"

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "subscribe to a websocket URL",
	Run: func(cmd *cobra.Command, args []string) {
		url := cmd.Flag("url").Value.String()

		ws, err := websocket.Dial(url, "", origin)
		if err != nil {
			log.Fatal(err)
		}

		for {
			msg := make([]byte, 512)
			var n int
			if n, err = ws.Read(msg); err != nil {
				log.Fatal(err)
			}
			log.Printf("%s\n", msg[:n])
		}
	},
}

func init() {
	rootCmd.AddCommand(subscribeCmd)

	subscribeCmd.Flags().StringP("url", "u", "", "websocket URL")
	subscribeCmd.Flags().UintP("message-buffer-size", "m", 1024, "message buffer size in bytes")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subscribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subscribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
