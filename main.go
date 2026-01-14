package main

import (
	"os"

	"github.com/73bits/pw/cmd"
)

func main() {
	if len(os.Args) < 2 {
		cmd.Usage()
		return
	}

	switch os.Args[1] {
	case "init":
		cmd.Init()
	case "add":
		cmd.Add()
	case "get":
		cmd.Get()
	default:
		cmd.Usage()
	}
}
