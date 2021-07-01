package parser

import (
	"reflect"
	"testing"

	"github.com/liubo0127/learn-go/crawler/engine"
)

func TestParseCityList(t *testing.T) {
	type args struct {
		contents []byte
	}
	tests := []struct {
		name string
		args args
		want engine.ParserResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCityList(tt.args.contents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCityList() = %v, want %v", got, tt.want)
			}
		})
	}
}
