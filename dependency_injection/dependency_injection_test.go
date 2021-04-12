package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
	}{
		{"test1", args{"Marry"}, "Hello, Marry"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			Greet(writer, tt.args.name)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Greet() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
