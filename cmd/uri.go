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
	Use:   "uri",
	Short: "get information about uri",
	Long:  `get information about uri",`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := internal.NewLogger()
		logs, err := logger.GetLogsFromFile()
		if err != nil {
			slog.Error(err.Error())
			return
		}

		uris := internal.SummarizeByURI(*logs)
		fmt.Println("🔍 URI, Counting...  🤖")
		for uri, val := range uris {
			fmt.Printf("URI: %s\nCount: %d\nMaxLatency: %d\n", uri, val.Count, val.MaxLatency)
			fmt.Println("====================================")
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
