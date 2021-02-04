package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:     "decode",
	Short:   "Use Base64 algorithm to decode a string",
	Long:    "Use Base64 algorithm to decode a string",
	Example: "gobase64 decode hello-world",
	Args:    cobra.MaximumNArgs(1),
	RunE:    decode,
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}

func decode(_ *cobra.Command, args []string) error {
	decodedBytes, err := base64.StdEncoding.DecodeString(args[0])
	if err != nil {
		return err
	}

	fmt.Println(string(decodedBytes))
	return nil
}
