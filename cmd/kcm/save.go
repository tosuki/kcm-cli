package main

import (
	"fmt"
	"kcm-cli/internal/core"
	"os"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save [profile-name]",
	Short: "Save current KDE Plasma configuration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]
		if err := core.SaveSnapshot(profileName); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving snapshot: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
