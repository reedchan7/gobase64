package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const defaultVersion = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gobase64",
	Long:  "Print the version number of gobase64",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gobase64 version:", getVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func getVersion() string {
	return defaultVersion
}
