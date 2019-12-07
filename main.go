package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: godl VERSION")
		os.Exit(1)
	}

	version := strings.TrimLeft(os.Args[1], "go")

	must(func() error {
		cmd := exec.Command("go", "get", "golang.org/dl/go"+version)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
		return nil
	}())

	must(func() error {
		cmd := exec.Command("go"+version, "download")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
		return nil
	}())
}

func must(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
