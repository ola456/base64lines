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
		if *decodeFlag {
			base64Str := scanner.Text()

			// add padding if necessary
			switch len(base64Str) % 4 {
			case 2:
				base64Str += "=="
			case 3:
				base64Str += "="
			}

			decoded, err := base64.StdEncoding.DecodeString(base64Str)

			if err != nil {
				fmt.Fprintf(os.Stderr, "decode error for %s: %s\n", base64Str, err)
				return
			}

			fmt.Println(string(decoded))
		} else {
			encoded := base64.StdEncoding.EncodeToString([]byte(scanner.Text()))

			fmt.Println(encoded)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
