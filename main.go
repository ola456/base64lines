package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	decodeFlag = flag.Bool("d", false, "Decode. Encode is the default.")
)

func main() {
	flag.Parse()

	// determine if input is from file or stdin
	var rawInput io.Reader
	filename := flag.Arg(0)
	if filename == "" || filename == "-" {
		rawInput = os.Stdin
	} else {
		r, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Sorry, I could not load input. Error: %v", err)
			os.Exit(1)
		}
		rawInput = r
	}

	// encode alt. decode line by line
	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		line := scanner.Text()

		if *decodeFlag {
			// add padding if necessary
			switch len(line) % 4 {
			case 2:
				line += "=="
			case 3:
				line += "="
			}

			decoded, err := base64.StdEncoding.DecodeString(line)

			if err != nil {
				fmt.Fprintf(os.Stderr, "decode error for %s: %s\n", line, err)
				continue
			}

			fmt.Println(string(decoded))
		} else {
			encoded := base64.StdEncoding.EncodeToString([]byte(line))

			fmt.Println(encoded)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
