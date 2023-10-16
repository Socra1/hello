package hello

import "testing"

func HelloText(t *testing.T) {
	want := "hello world"
	if got := Hello(); got != want {
		t.Errorf("Hello()=%q", got)
	}
}
