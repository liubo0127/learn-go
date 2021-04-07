package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	type args struct {
		name string
		language string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"english test1", args{"John", "english"}, "Hello, John"},
		{"spanish test1", args{"Marry", "spanish"}, "Hola, Marry"},
		{"spanish test2", args{"", "spanish"}, "Hola, World"},
		{"french test1", args{"Li", "french"}, "Bonjour, Li"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello(tt.args.name, tt.args.language)
			AssertCorrectMessage(t, got, tt.want)
		})
	}
}
