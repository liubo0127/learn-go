package integers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Add test1", args{1, 3}, 4},
		{"Add test2", args{2, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func ExampleAdd(x, y int) {
//	sum := Add(1, 5)
//	fmt.Println(sum)
//	// Output: 6
//}