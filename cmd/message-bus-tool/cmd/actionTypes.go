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

// actionTypesCmd represents the actionTypes command
var actionTypesCmd = &cobra.Command{
	Use:   "action-types",
	Short: "list action types",
	Run: func(cmd *cobra.Command, args []string) {
		baseURL, err := rootCmd.PersistentFlags().GetString(FlagBaseURL)
		if err != nil {
			panic(err)
		}

		url := fmt.Sprintf("http://%s/%s", strings.TrimRight(baseURL, "/"), basePath)
		client, err := codegen.NewClientWithResponses(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, err := client.GetActionTypesWithResponse(ctx)
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
			fmt.Fprintln(w, "SOURCE ID\tACTION NAME\tPROPERTY TYPES")
			fmt.Fprintln(w, "---------\t----------\t--------------")
		}

		for _, actionType := range *response.JSON200 {
			propertyTypes := make([]string, 0)
			for _, propertyType := range *actionType.PropertyTypeList {
				propertyTypes = append(propertyTypes, *propertyType.Name)
			}

			fmt.Fprintf(w, "%s\t%s\t%s\n", *actionType.SourceID, *actionType.Name, strings.Join(propertyTypes, ","))
		}

		w.Flush()
	},
}

func init() {
	listCmd.AddCommand(actionTypesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// actionTypesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// actionTypesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
