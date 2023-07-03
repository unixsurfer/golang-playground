//
// main.go
// Copyright (C) 2023 pavlos <pavlos.parissis@gmail.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	var b_stdout, b_stderr bytes.Buffer
	cmd := exec.Command("./stdout.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	b_stdout.ReadFrom(stdout)
	b_stderr.ReadFrom(stderr)
	fmt.Printf("stdout:[%s]\n", b_stdout.String())
	fmt.Printf("stderr:[%s]\n", b_stderr.String())

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
