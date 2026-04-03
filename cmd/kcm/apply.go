package main

import (
	"fmt"
	"kcm-cli/internal/core"
	"os"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply [profile-name]",
	Short: "Apply a saved KDE Plasma configuration snapshot",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]
		if err := core.ApplySnapshot(profileName); err != nil {
			fmt.Fprintf(os.Stderr, "Error applying snapshot: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
