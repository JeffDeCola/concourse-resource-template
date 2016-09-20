package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	version struct {
		Ref string `json:"ref"`
	}
	inputJSON struct {
		Params  map[string]string `json:"params"`
		Source  map[string]string `json:"source"`
		Version version           `json:"version"`
	}
	checkOut []version
	metadata struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	ioOut struct {
		Version  version    `json:"version"`
		Metadata []metadata `json:"metadata"`
	}
)

func main() {
	var (
		input   inputJSON
		decoder = json.NewDecoder(os.Stdin)
		encoder = json.NewEncoder(os.Stdout)
	)

	if err := decoder.Decode(&input); err != nil {
		panic("Failed to decode stdin")
	}

	switch os.Args[1] {
	case "check":
		//TODO: do check
	case "in":
		//TODO: do in
	case "out":

		fmt.Fprintln(os.Stderr, input.Params["text"])
		fmt.Fprintln(os.Stderr, os.Args[2])
		encoder.Encode(ioOut{
			Version:  version{Ref: "12"},
			Metadata: []metadata{{Name: "text", Value: input.Params["text"]}},
		})

	}
}
