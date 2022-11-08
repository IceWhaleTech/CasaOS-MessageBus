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
		baseURL, err := rootCmd.PersistentFlags().GetString(FlagBaseURL)
		if err != nil {
			panic(err)
		}

		sourceID, err := cmd.Flags().GetString(FlagSourceID)
		if err != nil {
			panic(err)
		}

		eventName, err := cmd.Flags().GetString(FlagEventName)
		if err != nil {
			panic(err)
		}

		wsURL := fmt.Sprintf("ws://%s%s/event_type/%s/%s/ws", strings.TrimRight(baseURL, "/"), basePath, sourceID, eventName)
		fmt.Printf("subscribed to %s\n", wsURL)

		ws, err := websocket.Dial(wsURL, "", origin)
		if err != nil {
			log.Fatal(err)
		}

		bufferSize, err := cmd.Flags().GetUint(FlagMessageBufferSize)
		if err != nil {
			log.Fatal(err)
		}

		for {
			msg := make([]byte, bufferSize)
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

	subscribeCmd.Flags().UintP(FlagMessageBufferSize, "m", 1024, "message buffer size in bytes")
	subscribeCmd.Flags().StringP(FlagSourceID, "s", "", "source id")
	subscribeCmd.Flags().StringP(FlagEventName, "n", "", "event name")

	if err := subscribeCmd.MarkFlagRequired(FlagSourceID); err != nil {
		panic(err)
	}

	if err := subscribeCmd.MarkFlagRequired(FlagEventName); err != nil {
		panic(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subscribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subscribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
