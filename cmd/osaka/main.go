package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mattn/osaka"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(osaka.Encode(scanner.Text()))
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(1)
	}
}
