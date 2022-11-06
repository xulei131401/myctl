package mysql

import (
	"github.com/spf13/cobra"
)

var (
	varBoolDSN bool

	Cmd = &cobra.Command{
		Use:   "mysql",
		Short: "mysql相关",
		RunE:  mysqlCommand,
		Args:  cobra.ArbitraryArgs,
	}
)

func init() {
	Cmd.Flags().BoolVar(&varBoolDSN, "dsn", false, "golang链接mysql dsn示例")
}
