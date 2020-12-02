package main

import (
	"os"

	"github.com/fadeev/mercury/cmd/mercuryd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
