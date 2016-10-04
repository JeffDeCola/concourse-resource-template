package main

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/JeffDeCola/resource-template/actions"
)

func Test_dumpinput(t *testing.T) {
	type args struct {
		verb   string
		input  actions.InputJSON
		logger *log.Logger
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		dumpinput(tt.args.verb, tt.args.input, tt.args.logger)
	}
}

func Test_doOutput(t *testing.T) {
	type args struct {
		output  interface{}
		encoder *json.Encoder
		logger  *log.Logger
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := doOutput(tt.args.output, tt.args.encoder, tt.args.logger); (err != nil) != tt.wantErr {
			t.Errorf("%q. doOutput() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		main()
	}
}
