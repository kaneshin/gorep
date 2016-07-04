package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	flag.Parse()
	os.Exit(run(flag.Args()...))
}

func run(args ...string) int {
	if terminal.IsTerminal(syscall.Stdin) {
		return 1
	}

	if len(args) == 0 {
		return 1
	}

	re, err := regexp.Compile(args[0])
	if err != nil {
		return 1
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		loc := re.FindIndex([]byte(text))
		if loc == nil {
			continue
		}
		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return 0
}
