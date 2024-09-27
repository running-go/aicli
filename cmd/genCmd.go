package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gemini",
	Short: "A CLI tool to interact with the Gemini API",
	Long:  "A CLI tool to interact with the Gemini API. You can ask questions and get responses in text or image format.",
	Run: func(cmd *cobra.Command, args []string) {
		printLogo()
		err := cmd.Help()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	genCmd.AddCommand(searchCmd)
	genCmd.AddCommand(imageCmd)
}
