package main

import (
	"fmt"
	"io"
	"os"
)

type logWrite struct{}

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Eroor:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
