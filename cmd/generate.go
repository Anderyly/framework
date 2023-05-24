package main

import "framework/cmd/generate"

func main() {
	if err := generate.Cmd().Execute(); err != nil {
		panic(err)
	}
}
