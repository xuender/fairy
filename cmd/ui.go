package cmd

import (
	"github.com/spf13/cobra"
)

// nolint: gochecknoinits
func init() {
	uiCmd := &cobra.Command{
		Use:   "ui",
		Short: "界面",
		Long:  `界面.`,
		Run: func(cmd *cobra.Command, args []string) {
			InitGUI(cmd).Run()
		},
	}
	rootCmd.AddCommand(uiCmd)
}
