package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
	"github.com/xuender/oils/oss"
)

// nolint: gochecknoinits
func init() {
	installCmd := &cobra.Command{
		Use:   "install",
		Short: "设置开机自启动",
		Long:  `设置开机自启动，实时监听分组目录.`,
		Run: func(cmd *cobra.Command, args []string) {
			str := base.Must1(oss.ExecOut("crontab", "-l"))
			if strings.Contains(str, "fairy") {
				logs.Info("Fairy already installed.")

				return
			}

			command := fmt.Sprintf("@reboot %s\n", os.Args[0])

			if str != "" {
				command = str + command
			}

			cron := exec.Command("crontab")
			inpipe := base.Must1(cron.StdinPipe())

			logs.Debug(command)
			base.Must1(inpipe.Write([]byte(command)))
			base.Must(cron.Start())
			base.Must(inpipe.Close())
			base.Must(cron.Wait())
			logs.Info("Fairy installed.")
		},
	}
	rootCmd.AddCommand(installCmd)
}
