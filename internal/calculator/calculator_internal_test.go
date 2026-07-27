package calculator

import "testing"

func Test_parseInputValidExpression(t *testing.T) {
	want := &expression{a: 2, b: 2, op: "+"}

	got, err := parseInput("2 + 2")
	if err != nil {
		t.Fatal("should not return an error")
	}

	if *want != *got {
		t.Errorf("want %v, got %v", *want, *got)
	}
}

func Test_parseInputInvalidExpression(t *testing.T) {
	want := `"invalid" is an invalid expression`

	_, err := parseInput("invalid")
	if err == nil {
		t.Fatal("should return an error")
	}

	if want != err.Error() {
		t.Errorf(`should error with %q, not %q`, want, err)
	}
}
