// Copyright (c) 2023 Julian Müller (ChaoticByte)

package main

import (
	"fmt"
	"os"
)

var Version = "development"

func main() {
	if len(os.Args) > 1 {
		ui := Cli{}
		ui.Run()
	} else {
		fmt.Println("No argument provided. Exiting.")
	}
}
