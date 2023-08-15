package main

import (
	"framework/cmd/serve"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "framework",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(serve.Cmd())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
