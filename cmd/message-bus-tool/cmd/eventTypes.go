/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/spf13/cobra"
)

// eventTypesCmd represents the eventTypes command
var eventTypesCmd = &cobra.Command{
	Use:   "event-types",
	Short: "list event types",
	Run: func(cmd *cobra.Command, args []string) {
		baseURL, err := rootCmd.PersistentFlags().GetString("base-url")
		if err != nil {
			panic(err)
		}
		client, err := codegen.NewClientWithResponses(strings.TrimRight(baseURL, "/") + basePath)
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, err := client.GetEventTypesWithResponse(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}

		if response.StatusCode() != http.StatusOK {
			fmt.Println("error")
			return
		}

		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 3, ' ', 0)

		if len(*response.JSON200) > 0 {
			fmt.Fprintln(w, "SOURCE ID\tEVENT NAME\tPROPERTY TYPES")
			fmt.Fprintln(w, "---------\t----------\t--------------")
		}

		for _, eventType := range *response.JSON200 {
			propertyTypes := make([]string, 0)
			for _, propertyType := range *eventType.PropertyTypeList {
				propertyTypes = append(propertyTypes, *propertyType.Name)
			}

			fmt.Fprintf(w, "%s\t%s\t%s\n", *eventType.SourceID, *eventType.Name, strings.Join(propertyTypes, ","))
		}

		w.Flush()
	},
}

func init() {
	listCmd.AddCommand(eventTypesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eventTypesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eventTypesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
