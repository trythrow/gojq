package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
)

func main() {
	// Define flags
	noColor := flag.Bool("no-color", false, "Disable colored output")
	compact := flag.Bool("compact", false, "Print compact JSON (no indent)")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		var parsed interface{}
		if err := json.Unmarshal([]byte(line), &parsed); err != nil {
			// Not valid JSON, print raw
			fmt.Println(line)
			continue
		}

		if *noColor {
			// Print plain JSON
			var output []byte
			var err error

			if *compact {
				output, err = json.Marshal(parsed)
			} else {
				output, err = json.MarshalIndent(parsed, "", "  ")
			}

			if err != nil {
				fmt.Println(line)
				continue
			}

			fmt.Println(string(output))
		} else {
			// Use colorjson
			f := colorjson.NewFormatter()
			if !*compact {
				f.Indent = 2
			}

			s, err := f.Marshal(parsed)
			if err != nil {
				fmt.Println(line)
				continue
			}

			fmt.Println(string(s))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
