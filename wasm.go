package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("make", "")
	if _, err := cmd.Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error running Makefile: ", err)
		os.Exit(1)
	}
}
