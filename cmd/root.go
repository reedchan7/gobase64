package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gobase64 [command] [flags]",
	Short: "A CLI for encoding or decoding by Base64",
	Long: "gobase64 is a CLI for encoding or decoding string by Base64",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
