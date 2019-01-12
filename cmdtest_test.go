package cmdtest

import (
	"testing"
)

func TestRun(t *testing.T) {
	cmd := Command("awk", "{print tolower($0)}")
	got, want, err := cmd.Run("input.txt", "output.txt")
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("Run(\"%s\") = %s, want: %s", "input.txt", got, want)
	}
}
