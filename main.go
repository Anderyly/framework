package main

import (
	main2 "framework/cmd/serve"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "im-server",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(main2.Cmd())
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
