package conv

import (
	
	"testing"
	"math"
)

func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 0},
		{input: 180, want: 82.22},
		{input: 50, want: 10},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}


func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 400, want: 126.85},
		{input: 10, want: -263.15},
		{input: 300, want: 26.85},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: -89.4, want: -128.92},
		{input: 94, want: 201.2},
		{input: -51.11, want: -60},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: -89.4, want: 183.75},
		{input: -263.15, want: 10},
		{input: 26.85, want: 300},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: -50, want: 227.59},
		{input: 300, want: 422.04},
		{input: 150, want: 338.71},
	}

	for _, tc := range tests {
		got := FarhenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 227.59, want: -50.01},
		{input: 422.04, want: 300},
		{input: 338.01, want: 148.75},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func withinTolerance(a, b, error float64) bool {
	if a == b {
	  return true
	}
  
	difference := math.Abs(a - b)
  
	if b == 0 {
	  return difference < error
	}
  
	return (difference/math.Abs(b)) < error
  }



