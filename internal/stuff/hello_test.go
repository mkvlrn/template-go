package stuff_test

import (
	"testing"

	"github.com/mkvlrn/template-go/internal/stuff"
)

func TestHello(t *testing.T) {
	got := stuff.Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got: %q, want %q", got, want)
	}
}
