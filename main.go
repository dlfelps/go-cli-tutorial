package main

import (
	"fmt"
	"os"

	"gocli-teacher/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
