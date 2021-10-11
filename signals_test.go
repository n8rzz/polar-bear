package main

import (
	"testing"
)

func TestIsFastAboveSlowForLength(t *testing.T) {
	type args struct {
		slow   []float64
		fast   []float64
		length int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"fast is above slow for #length", args{slow: []float64{-0.5, -0.5}, fast: []float64{1, 1}, length: 2}, true},
		{"fast is above slow for less than #length", args{slow: []float64{-0.5, 0}, fast: []float64{1, -1}, length: 2}, false},
		{"fast is above then equal to slow", args{slow: []float64{-0.5, 0}, fast: []float64{1, 0}, length: 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFastAboveSlowForLength(tt.args.slow, tt.args.fast, tt.args.length); got != tt.want {
				t.Errorf("IsFastAboveSlowForLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFastAboveSlow(t *testing.T) {
	type args struct {
		slow []float64
		fast []float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"fast is above slow", args{slow: []float64{-0.5}, fast: []float64{0}}, true},
		{"slow is above fast", args{slow: []float64{-0.5}, fast: []float64{-1}}, false},
		{"fast is equal to slow", args{slow: []float64{-0.5}, fast: []float64{-1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFastAboveSlow(tt.args.slow, tt.args.fast); got != tt.want {
				t.Errorf("IsFastAboveSlow() = %v, want %v", got, tt.want)
			}
		})
	}
}
