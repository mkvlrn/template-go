package calculator_test

import (
	"testing"

	"github.com/mkvlrn/template-go/internal/calculator"
)

func TestSolve(t *testing.T) {
	tests := map[string]float64{
		"2 + 2":   4,
		"70 - 1":  69,
		"2.5 * 3": 7.5,
		"999 / 3": 333,
	}

	for name, want := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := calculator.Solve(name)
			if err != nil {
				t.Fatal("should not return an error")
			}

			if want != got {
				t.Fatalf("want %f, got %f", want, got)
			}
		})
	}
}

func TestSolve_Error(t *testing.T) {
	tests := map[string]struct {
		expression string
		message    string
	}{
		"divide by zero":     {"2 / 0", "cannot divide by zero"},
		"invalid expression": {"invalid", `"invalid" is an invalid expression`},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := calculator.Solve(tc.expression)
			if err == nil {
				t.Fatal("should return an error")
			}

			if tc.message != err.Error() {
				t.Fatalf("want err %q, got err %q", tc.message, err.Error())
			}
		})
	}
}
