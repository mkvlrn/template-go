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
	tt := []testCase{{2, 2, 4}, {2, 3, 5}, {1000, 2000, 3000}}

	for _, tc := range tt {
		testName := fmt.Sprintf("%f + %f = %f", tc.a, tc.b, tc.want)
		t.Run(testName, func(t *testing.T) {
			got := math.Add(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("got %f, want %f", got, tc.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tt := []testCase{{2, 2, 0}, {2, 3, -1}, {1000, 2000, -1000}}

	for _, tc := range tt {
		testName := fmt.Sprintf("%f - %f = %f", tc.a, tc.b, tc.want)
		t.Run(testName, func(t *testing.T) {
			got := math.Subtract(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("got %f, want %f", got, tc.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tt := []testCase{{2, 2, 4}, {2, 3, 6}, {1000, 2000, 2000000}}

	for _, tc := range tt {
		testName := fmt.Sprintf("%f x %f = %f", tc.a, tc.b, tc.want)
		t.Run(testName, func(t *testing.T) {
			got := math.Multiply(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("got %f, want %f", got, tc.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tt := []testCase{{2, 2, 1}, {2, 4, 0.5}, {999, 333, 3}, {2, 0, 0}}

	for _, tc := range tt {
		testName := fmt.Sprintf("%f / %f = %f", tc.a, tc.b, tc.want)

		t.Run(testName, func(t *testing.T) {
			got, err := math.Divide(tc.a, tc.b)

			if tc.b == 0 {
				if !errors.Is(err, math.ErrDivideByZero) {
					t.Errorf("got err %s, want err %s", err, math.ErrDivideByZero)
				}

				return
			}

			if got != tc.want {
				t.Errorf("got %f, want %f", got, tc.want)
			}
		})
	}
}
