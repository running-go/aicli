package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const CliVersion = "v0.1.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Know the installed version of aicli",
	Long:  `This command will help you to know the installed version of aicli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aicli version:", CliVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
