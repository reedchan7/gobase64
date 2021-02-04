package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encodeCmd)
}

var encodeCmd = &cobra.Command{
	Use:     "encode",
	Short:   "Use Base64 algorithm to encode a string",
	Long:    "Use Base64 algorithm to encode a string",
	Example: "gobase64 encode hello-world",
	Args:    cobra.MaximumNArgs(1),
	Run:     encode,
}

func encode(_ *cobra.Command, args []string) {
	encodedStr := base64.StdEncoding.EncodeToString([]byte(args[0]))
	fmt.Println(encodedStr)
}
