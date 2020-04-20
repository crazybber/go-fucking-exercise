package basic

import (
	"fmt"
	"testing"
)

func MaxFloat32InTwo(a, b float32) float32 {
	if a < b {
		return b
	}
	return a
}

func MaxFloat32(a, b, c float32) float32 {
	return MaxFloat32InTwo(MaxFloat32InTwo(a, b), c)
}

func TestMaxFloatTableDriven(t *testing.T) {
	var tests = []struct {
		a, b, c float32
		want    float32
	}{
		{0, 0.4, 1, 1},
		{1, 0, 0, 1},
		{12, -2.4, 3, 12},
		{-0.8, -0.4, -0.3, -0.3},
		{-1, 0, -1, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%f,%f", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := MaxFloat32(tt.a, tt.b, tt.c)
			if ans != tt.want {
				t.Errorf("got %f, want %f", ans, tt.want)
			}
		})
	}
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
