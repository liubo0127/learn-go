package main

import (
	"reflect"
	"testing"
)

func TestArraySum(t *testing.T) {
	type args struct {
		numbers [5]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Array test1", args{[5]int{1,2,3,4,5}}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArraySum(tt.args.numbers); got != tt.want {
				t.Errorf("ArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Slice test1", args{[]int{1,2,3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceSum(tt.args.numbers); got != tt.want {
				t.Errorf("SliceSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceSumAll(t *testing.T) {
	type args struct {
		numbers [][]int
	}
	tests := []struct{
		name string
		args args
		want []int
	}{
		{"SliceSumAll test1", args{[][]int{{1,2,3}, {4,5,6}}}, []int{6, 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceSumAll(tt.args.numbers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceNumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}