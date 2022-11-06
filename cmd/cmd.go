package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
	"github.com/xulei131401/myctl/internal/version"
	"github.com/xulei131401/myctl/mysql"
)

const (
	codeFailure = 1
)

var rootCmd = &cobra.Command{
	Use:   "myctl",
	Short: "myctl工具自用",
	Long:  `常用命令进行封装打包，基于cobra，随时扩展`,
}

func init() {
	rootCmd.Version = fmt.Sprintf("%s %s/%s", version.BuildVersion, runtime.GOOS, runtime.GOARCH)
	rootCmd.AddCommand(mysql.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(aurora.Red(err.Error()))
		os.Exit(codeFailure)
	}
}
