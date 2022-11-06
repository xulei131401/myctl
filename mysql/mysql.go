package mysql

import (
	"fmt"
	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
)

func mysqlCommand(_ *cobra.Command, _ []string) (err error) {
	defer func() {
		if err == nil {
			fmt.Println(aurora.Green("Done."))
		}
	}()

	if varBoolDSN {
		dsn := "root:xxxxx@tcp(127.0.0.1:3306)/myboot?charset=utf8mb4&parseTime=True"
		fmt.Println(dsn)
	}

	return nil
}
