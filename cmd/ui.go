package cmd

import (
	"fmt"

	"github.com/segmentio/ksuid"
	"github.com/spf13/cobra"
)

// nolint: gochecknoinits
func init() {
	uiCmd := &cobra.Command{
		Use:   "ui",
		Short: "界面",
		Long:  `界面.`,
		Run: func(cmd *cobra.Command, args []string) {
			kid := ksuid.New()
			fmt.Println(kid.String())
			InitGUI(cmd).Run()
		},
	}
	rootCmd.AddCommand(uiCmd)
}
