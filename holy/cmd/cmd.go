package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
	"github.com/xulei131401/holy-cmd/holy/internal/version"
)

const (
	codeFailure = 1
)

var rootCmd = &cobra.Command{
	Use:   "holy",
	Short: "holy工具自用",
	Long:  `常用命令进行封装打包，基于cobra，随时扩展`,
}

func init() {
	// 添加版本信息
	rootCmd.Version = fmt.Sprintf("%s %s/%s", version.BuildVersion, runtime.GOOS, runtime.GOARCH)

	// local选项
	//rootCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1", "主机")
	// 参数必填设置
	// _ = rootCmd.MarkFlagRequired("host")

	// 添加子命令
	//rootCmd.AddCommand(sub.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(aurora.Red(err.Error()))
		os.Exit(codeFailure)
	}
}
