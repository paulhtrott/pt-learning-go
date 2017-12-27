package main

import (
	"flag"
	"fmt"
)

func main() {
	// Compile to a particular platform
	// GOOS=windows GOARCH=amd64 go build cli.go
	var cmd string

	flag.StringVar(&cmd, "cmd", cmd, `cmd can be either "hello" or "bye"`)
	flag.Parse()

	switch cmd {
	case "hello":
		fmt.Println("Hello!")
	case "bye":
		fmt.Println("bye!")
	default:
		flag.Usage()
	}
}
