package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	header, err := in.ReadBytes('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading header: %v", err)
		os.Exit(1)
	}
	os.Stdout.Write(header)

	cmd := exec.Command("sort", os.Args[1:]...)
	cmd.Stdin = in
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "exec: %v", err)
		os.Exit(2)
	}
}
