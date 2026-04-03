package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kcm",
	Short: "KCM is a CLI for managing KDE Plasma configuration snapshots",
	Long:  `KDE Config Manager (KCM) allows you to backup, restore and manage multiple Plasma desktop profiles easily.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("KCM-CLI: Use 'kcm --help' for more information.")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
