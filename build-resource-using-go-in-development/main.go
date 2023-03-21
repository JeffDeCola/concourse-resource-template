// concourse-resource-template main.go

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/JeffDeCola/concourse-resource-template/actions"
)

// dumpinput ...
func dumpinput(verb string, input actions.InputJSON, logger *log.Logger) {

	// Get the working directory from arg $2 (Only for IN and OUT)
	if verb != "check" {
		var workingdir = os.Args[2]
		logger.Print("WORKING_DIR = ", workingdir)
		logger.Print("List whats in the working directory:")
		//ls -lat $WORKING_DIR
		files, _ := ioutil.ReadDir(workingdir)
		for _, f := range files {
			logger.Print(f.Name())
		}
	}

	// List whats in the input stdin .json
	// MashalIndent makes it print nicely
	logger.Print("This is the input stdin .json format:")
	b, _ := json.MarshalIndent(input, "", "  ")
	logger.Print(string(b))

}

// doOutput ...
func doOutput(output interface{}, encoder *json.Encoder, logger *log.Logger) error {

	logger.Print(".json output is:")
	b, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return err
	}
	logger.Print(string(b))

	// encode output to stdout
	return encoder.Encode(output)

}

func main() {

	var (
		input   actions.InputJSON
		decoder = json.NewDecoder(os.Stdin)
		encoder = json.NewEncoder(os.Stdout)
		arg1    = os.Args[1]
		logger  = log.New(os.Stderr, "resource:", log.Lshortfile)
	)

	// Concourse passes .json on stdin
	if err := decoder.Decode(&input); err != nil {
		logger.Fatalf("Failed to decode to stdin: %s", err)
	}

	// arg1 is check, in, or out.
	switch arg1 {

	case "check":
		logger.Print("CHECK (THE RESOURCE VERSION(s))")
		dumpinput(arg1, input, logger)

		output, err := actions.Check(input, logger)
		if err != nil {
			logger.Fatalf("Input missing a field: %s", err)
		}

		if err = doOutput(output, encoder, logger); err != nil {
			logger.Fatalf("Failed to encode to stdout: %s", err)
		}

	case "in":
		logger.Print("IN (FETCH THE RESOURCE)")
		dumpinput(arg1, input, logger)

		output, err := actions.In(input, logger)
		if err != nil {
			logger.Fatalf("Input missing a field: %s", err)
		}

		if err := doOutput(output, encoder, logger); err != nil {
			logger.Fatalf("Failed to encode to stdout: %s", err)
		}

	case "out":
		logger.Print("OUT (UPDATE THE RESOURCE)")
		dumpinput(arg1, input, logger)

		output, err := actions.Out(input, logger)
		if err != nil {
			logger.Fatalf("Input missing a field: %s", err)
		}

		if err := doOutput(output, encoder, logger); err != nil {
			logger.Fatalf("Failed to encode to stdout: %s", err)
		}
	}

}
