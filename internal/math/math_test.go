package math_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mkvlrn/template-go/internal/math"
)

type testCase struct {
	a    float64
	b    float64
	want float64
}

func TestAdd(t *testing.T) {
	tc := []testCase{{2, 2, 4}, {2, 3, 5}, {1000, 2000, 3000}}

	for _, tt := range tc {
		testName := fmt.Sprintf("%f + %f = %f", tt.a, tt.b, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := math.Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tc := []testCase{{2, 2, 0}, {2, 3, -1}, {1000, 2000, -1000}}

	for _, tt := range tc {
		testName := fmt.Sprintf("%f - %f = %f", tt.a, tt.b, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := math.Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tc := []testCase{{2, 2, 4}, {2, 3, 6}, {1000, 2000, 2000000}}

	for _, tt := range tc {
		testName := fmt.Sprintf("%f x %f = %f", tt.a, tt.b, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := math.Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tc := []testCase{{2, 2, 1}, {2, 4, 0.5}, {999, 333, 3}, {2, 0, 0}}

	for _, tt := range tc {
		testName := fmt.Sprintf("%f / %f = %f", tt.a, tt.b, tt.want)

		t.Run(testName, func(t *testing.T) {
			got, err := math.Divide(tt.a, tt.b)
			if err != nil {
				if !errors.Is(err, math.ErrDivideByZero) {
					t.Errorf("got err %s, want err %s", err, math.ErrDivideByZero)
				}
			} else {
				if got != tt.want {
					t.Errorf("got %f, want %f", got, tt.want)
				}
			}
		})
	}
}
