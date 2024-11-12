/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"

	"github.com/ei-sugimoto/datapuppy/internal"
	"github.com/spf13/cobra"
)

// uricountCmd represents the uricount command
var uricountCmd = &cobra.Command{
	Use:   "uricount",
	Short: "Counts the number of accesses per URI.",
	Long:  `Counts the number of accesses per URI.`,
	Run: func(cmd *cobra.Command, args []string) {
		newLog := internal.Detail{
			ID:        "1",
			Time:      "2024-01-01T00:00:00Z",
			RemoteIP:  "127.0.0.1",
			Host:      "example.com",
			Method:    "GET",
			UserAgent: "Mozilla/5.0",
			Status:    200,
			Latency:   100,
		}
		logger := internal.NewLogger()
		err := logger.UpsertLogToFile(&newLog)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		logs, err := logger.GetLogsFromFile()
		if err != nil {
			slog.Error(err.Error())
			return
		}

		uriCount := internal.SummarizeRequestCountByURI(*logs)
		fmt.Println("🔍 URI, Counting...  🤖")
		for uri, count := range uriCount {
			fmt.Printf("uri:%s\n\tcount:%d\n", uri, count)
		}

	},
}

func init() {
	rootCmd.AddCommand(uricountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uricountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uricountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
