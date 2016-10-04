package actions

import (
	"log"
	"reflect"
	"testing"
)

func Test_getversions(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := getversions(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. getversions() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestCheck(t *testing.T) {
	type args struct {
		input  InputJSON
		logger *log.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    checkOutputJSON
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Check(tt.args.input, tt.args.logger)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Check() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Check() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestIn(t *testing.T) {
	type args struct {
		input  InputJSON
		logger *log.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    inOutputJSON
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := In(tt.args.input, tt.args.logger)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. In() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. In() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestOut(t *testing.T) {
	type args struct {
		input  InputJSON
		logger *log.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    outOutputJSON
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Out(tt.args.input, tt.args.logger)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Out() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Out() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
