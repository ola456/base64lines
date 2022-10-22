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

	// Determine if input is from file or stdin
	var rawInput io.Reader
	filename := flag.Arg(0)
	if filename == "" || filename == "-" {
		rawInput = os.Stdin
	} else {
		r, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Sorry, I could not load input. Error: %v", err)
			os.Exit(1)
		}
		rawInput = r
	}

	// Encode alt. decode line by line
	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		if *decodeFlag {
			decoded, err := base64.StdEncoding.DecodeString(scanner.Text())

			if err != nil {
				fmt.Println("decode error:", err)
				return
			}

			fmt.Println(string(decoded))
		} else {
			encoded := base64.StdEncoding.EncodeToString([]byte(scanner.Text()))

			fmt.Println(encoded)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
