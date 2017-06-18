package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mattn/osaka"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			break
		}
		fmt.Print(osaka.Encode(string(b)))
	}
}
