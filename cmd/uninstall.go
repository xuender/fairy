/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
	"github.com/xuender/startup"
)

// nolint: gochecknoinits
func init() {
	uninstallCmd := &cobra.Command{
		Use:   "uninstall",
		Short: "取消开机自启动",
		Long:  `取消开机自启动.`,
		Run: func(cmd *cobra.Command, args []string) {
			base.Must(startup.Uninstall())
			logs.Info("Fairy uninstalled.")
		},
	}
	rootCmd.AddCommand(uninstallCmd)
}
