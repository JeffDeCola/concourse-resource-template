package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	metadata struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	checkoutputJSON []version
	inoutputJSON    struct {
		Version  version    `json:"version"`
		Metadata []metadata `json:"metadata"`
	}
	outoutputJSON struct {
		Version  version    `json:"version"`
		Metadata []metadata `json:"metadata"`
	}
)

func check(input inputJSON) {

	// PARSE THE JSON FILE /tmp/input.json
	var source1 = input.Source["source1"]
	var source2 = input.Source["source2"]
	var ref = input.Version.Ref
	fmt.Fprintln(os.Stderr, "source are")
	fmt.Fprintln(os.Stderr, source1, source2)
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "ref is")
	fmt.Fprintln(os.Stderr, ref)
	fmt.Fprintln(os.Stderr, "")

	// CHECK (THE RESOURCE VERSION(s)) *******************************************************************
	// Mimic a fetch and output the following versions for IN.

	ver1 := "123"
	ver2 := "3de"
	ver3 := "456"

	// OUTPUT **************************************************************************************
	output := &checkoutputJSON{
		{Ref: ver1},
		{Ref: ver2},
		{Ref: ver3},
	}

	fmt.Fprintln(os.Stderr, ".json output is:")
	b, _ := json.MarshalIndent(output, "", "  ")
	fmt.Fprintln(os.Stderr, string(b))
	fmt.Fprintln(os.Stderr, "")

	// Encode .json and send to stdout
	json.NewEncoder(os.Stdout).Encode(&output)

}

func in(input inputJSON) {

	// PARSE THE JSON FILE /tmp/input.json
	var source1 = input.Source["source1"]
	var source2 = input.Source["source2"]
	var param1 = input.Params["param1"]
	var param2 = input.Params["param2"]
	var ref = input.Version.Ref
	fmt.Fprintln(os.Stderr, "source are")
	fmt.Fprintln(os.Stderr, source1, source2)
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "params are")
	fmt.Fprintln(os.Stderr, param1, param2)
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "ref is")
	fmt.Fprintln(os.Stderr, ref)
	fmt.Fprintln(os.Stderr, "")

	// IN (FETCH THE RESOURCE) *************************************************************************
	// Mimic a fetch and place a fetched.json file in the working directory that contains the following.

	jsonfile := "Hi everone, This is a file I made"

	// Create a fake fetched file
	filewrite, err := os.Create("fetch.json")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer filewrite.Close()
	fmt.Fprintf(filewrite, jsonfile)

	//ls -lat $WORKING_DIR
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Fprintln(os.Stderr, f.Name())
	}
	fmt.Fprintln(os.Stderr, "")

	// Cat the file
	fmt.Fprintln(os.Stderr, "Cat fetch.json")
	file, err := os.Open("fetch.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bb, err := ioutil.ReadAll(file)
	fmt.Print(string(bb))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "")

	var monkeyname = "Larry"

	// OUTPUT **************************************************************************************
	output := &inoutputJSON{
		Version: version{Ref: ref},
		Metadata: []metadata{
			{Name: "nameofmonkey", Value: monkeyname},
			{Name: "author", Value: "Jeff DeCola"},
		},
	}

	fmt.Fprintln(os.Stderr, ".json output is:")
	b, _ := json.MarshalIndent(output, "", "  ")
	fmt.Fprintln(os.Stderr, string(b))
	fmt.Fprintln(os.Stderr, "")

	// Encode .json and send to stdout
	json.NewEncoder(os.Stdout).Encode(&output)

}

func out(input inputJSON) {

	// PARSE THE JSON FILE /tmp/input.json
	var param1 = input.Params["param1"]
	var param2 = input.Params["param2"]
	var source1 = input.Source["source1"]
	var source2 = input.Source["source2"]
	fmt.Fprintln(os.Stderr, "params are")
	fmt.Fprintln(os.Stderr, param1, param2)
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "source are")
	fmt.Fprintln(os.Stderr, source1, source2)
	fmt.Fprintln(os.Stderr, "")

	// OUT (UPDATE THE RESOURCE) *************************************************************************
	// Mimic an out.

	var monkeyname = "Henry"
	var ref = "123"

	// OUTPUT **************************************************************************************
	output := &outoutputJSON{
		Version: version{Ref: ref},
		Metadata: []metadata{
			{Name: "nameofmonkey", Value: monkeyname},
			{Name: "author", Value: "Jeff DeCola"},
		},
	}

	fmt.Fprintln(os.Stderr, ".json output is:")
	b, _ := json.MarshalIndent(output, "", "  ")
	fmt.Fprintln(os.Stderr, string(b))
	fmt.Fprintln(os.Stderr, "")

	// Encode .json and send to stdout
	json.NewEncoder(os.Stdout).Encode(&output)

}

func main() {

	// Decode the .json from stdin and place in a .json struct.
	var input inputJSON
	var decoder = json.NewDecoder(os.Stdin)

	err := decoder.Decode(&input)

	// Test if error reading stdin .json format
	if err != nil {
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

	// Get the working directory from arg $2 (Only for IN and OUT)
	if os.Args[1] != "check" {
		var workingdir = os.Args[2]
		fmt.Fprintln(os.Stderr, "WORKING_DIR = ", workingdir)
		fmt.Fprintln(os.Stderr, "List whats in the working directory")
		//ls -lat $WORKING_DIR
		files, _ := ioutil.ReadDir("./")
		for _, f := range files {
			fmt.Fprintln(os.Stderr, f.Name())
		}
		fmt.Fprintln(os.Stderr, "")
	}

	// List whats in the input stdin .json
	// MashalIndent makes it print nicely
	fmt.Fprintln(os.Stderr, "This is the input stdin .json format:")
	b, _ := json.MarshalIndent(input, "", "  ")
	fmt.Fprintln(os.Stderr, string(b))
	fmt.Fprintln(os.Stderr, "")

	switch os.Args[1] {
	case "check":
		check(input)
	case "in":
		in(input)
	case "out":
		out(input)
	}

}
