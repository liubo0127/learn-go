package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"John"}, "Hello, John"},
		{"test2", args{"Marry"}, "Hello, Marry"},
		{"test3", args{""}, "Hello, World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello(tt.args.name)
			AssertCorrectMessage(t, got, tt.want)
		})
	}
}
