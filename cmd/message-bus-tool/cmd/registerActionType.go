/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/spf13/cobra"
)

// registerActionTypeCmd represents the registerActionType command
var registerActionTypeCmd = &cobra.Command{
	Use:   "action-type",
	Short: "register an action type",
	Run: func(cmd *cobra.Command, args []string) {
		baseURL, err := rootCmd.PersistentFlags().GetString(FlagBaseURL)
		if err != nil {
			panic(err)
		}

		url := fmt.Sprintf("http://%s%s", strings.TrimRight(baseURL, "/"), basePath)
		client, err := codegen.NewClientWithResponses(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, err := client.RegisterActionTypeWithResponse(ctx, codegen.RegisterActionTypeJSONRequestBody{})
		if err != nil {
			fmt.Println(err)
			return
		}

		if response.StatusCode() != http.StatusCreated {
			fmt.Println(response.Status())
			return
		}

		fmt.Println("success")
	},
}

func init() {
	registerCmd.AddCommand(registerActionTypeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerActionTypeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerActionTypeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
