package main

import (
	"strings"
	"testing"
)

func TestShape(t *testing.T) {
	type args struct {
		shape Shape
	}

	tests := []struct {
		name string // [rectangle, circle]
		args args
		want float64 // [rectangle area, rectangle perimeter, circle area, circle perimeter]
	}{
		{"rectangle area test1", args{&Rectangle{4, 5}}, 20},
		{"rectangle perimeter test1", args{&Rectangle{4, 5}}, 18},
		{"circle area test1", args{&Circle{4}}, 50.26548245743669},
		{"circle perimeter test1", args{&Circle{4}}, 25.132741228718345},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(strings.ToLower(tt.name), "area") {
				if got := tt.args.shape.Area(); got != tt.want {
					t.Errorf("Area() = %v, want %v", got, tt.want)
				}
			}
			if strings.Contains(strings.ToLower(tt.name), "perimeter") {
				if got := tt.args.shape.Perimeter(); got != tt.want {
					t.Errorf("Perimeter() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
