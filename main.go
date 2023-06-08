package main

import (
	"framework/cmd/serve"
	"framework/cmd/sql"
	_ "framework/docs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "framework",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// @title           framework
// @version         1.0
// @description     This is a gin framework

// @contact.name   Anderyly
// @contact.url    http://blog.aaayun.cc
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9090
// @BasePath  /api
func main() {
	rootCmd.AddCommand(sql.Cmd())
	rootCmd.AddCommand(serve.Cmd())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
