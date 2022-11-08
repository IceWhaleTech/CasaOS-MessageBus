/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/websocket"
)

const origin = "http://localhost/"

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "subscribe to a websocket URL",
	Run: func(cmd *cobra.Command, args []string) {
		baseURL, err := rootCmd.PersistentFlags().GetString("base-url")
		if err != nil {
			panic(err)
		}

		sourceID, err := cmd.Flags().GetString("source-id")
		if err != nil {
			panic(err)
		}

		eventName, err := cmd.Flags().GetString("event-name")
		if err != nil {
			panic(err)
		}

		wsURL := fmt.Sprintf("ws://%s%s/event_type/%s/%s/ws", strings.TrimRight(baseURL, "/"), basePath, sourceID, eventName)
		fmt.Printf("subscribed to %s\n", wsURL)

		ws, err := websocket.Dial(wsURL, "", origin)
		if err != nil {
			log.Fatal(err)
		}

		buffer_size, err := cmd.Flags().GetUint("message-buffer-size")
		if err != nil {
			log.Fatal(err)
		}

		for {
			msg := make([]byte, buffer_size)
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

	subscribeCmd.Flags().UintP("message-buffer-size", "m", 1024, "message buffer size in bytes")
	subscribeCmd.Flags().StringP("source-id", "s", "", "source id")
	subscribeCmd.Flags().StringP("event-name", "n", "", "event name")

	subscribeCmd.MarkFlagRequired("source-id")
	subscribeCmd.MarkFlagRequired("event-name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subscribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subscribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
