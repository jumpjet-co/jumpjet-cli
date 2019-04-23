package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "jumpjet",
	Short: "Jumpjet cli tool",
	Long: `Jumpjet cli tool`,
}

func Execute() {
	initSubCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
