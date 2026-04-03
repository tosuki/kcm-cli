package main

import (
	"fmt"
	"kcm-cli/internal/core"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved KDE Plasma configuration snapshots",
	Run: func(cmd *cobra.Command, args []string) {
		profiles, err := core.ListProfiles()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing profiles: %v\n", err)
			os.Exit(1)
		}

		if len(profiles) == 0 {
			fmt.Println("No snapshots found. Use 'kcm save [name]' to create one.")
			return
		}

		fmt.Printf("%-20s | %-20s | %-15s | %-15s\n", "NAME", "CREATED AT", "THEME", "PLASMA")
		fmt.Println("--------------------------------------------------------------------------------")
		for _, p := range profiles {
			fmt.Printf("%-20s | %-20s | %-15s | %-15s\n",
				p.Name,
				p.CreatedAt.Format("2006-01-02 15:04"),
				p.GlobalTheme,
				p.PlasmaVersion)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
