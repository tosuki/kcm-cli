package main

import (
	"fmt"
	"kcm-cli/internal/tui"
	"os"

	"github.com/spf13/cobra"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Start the interactive TUI to manage snapshots",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tui.StartUI(); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting TUI: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)
}
