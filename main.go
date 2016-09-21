package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	)

	if err := decoder.Decode(&input); err != nil {
		panic("Failed to decode stdin")
	}

	switch os.Args[1] {
	case "check":
		fmt.Fprintln(os.Stderr, "CHECK (THE RESOURCE VERSION(s))")
		fmt.Fprintln(os.Stderr, "")
	case "in":
		fmt.Fprintln(os.Stderr, "IN (FETCH THE RESOURCE)")
		fmt.Fprintln(os.Stderr, "")
	case "out":
		fmt.Fprintln(os.Stderr, "OUT (UPDATE THE RESOURCE)")
		fmt.Fprintln(os.Stderr, "")
	}

	// Get the working directory from arg $2
	var workingdir = os.Args[2]
	fmt.Fprintln(os.Stderr, "WORKING_DIR = ", workingdir)
	fmt.Fprintln(os.Stderr, "List whats in the working directory")
	//ls -lat $WORKING_DIR
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Fprintln(os.Stderr, f.Name())
	}
	fmt.Fprintln(os.Stderr, "")

	// List whats in the input stdin json
	b, _ := json.MarshalIndent(input, "", "  ")
	fmt.Fprintln(os.Stderr, string(b))
	fmt.Fprintln(os.Stderr, "")




	switch os.Args[1] {

	// CHECK (THE RESOURCE VERSION(s)) *******************************************************************
	// Mimic a fetch and output the following versions for IN.
	case "check":

		//TODO: do CHECK

	// IN (FETCH THE RESOURCE) *************************************************************************
	// Mimic a fetch and place a fetched.json file in the working directory that contains the following.
	case "in":

		//TODO: do IN

	// OUT (UPDATE THE RESOURCE) *************************************************************************
	// Mimic an out.
	case "out":

		json.NewEncoder(os.Stdout).Encode(ioOut{
			Version: version{Ref: "12"},
			// Metadata: []metadata{{Name: "text", Value: input.Params["text"]}},
		})

	}
}
