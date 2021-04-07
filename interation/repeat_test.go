package interation

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	type args struct {
		character string
		cnt       int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"repeat test1", args{"a", 5}, "aaaaa"},
		{"repeat test2", args{"test", 3}, "testtesttest"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.character, tt.args.cnt); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	r := Repeat("a", 5)
	fmt.Println(r)
	// Output: aaaaa
}