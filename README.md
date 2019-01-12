# Package cmdtest

Package cmdtest helps to test competitive programming solutions.

```go
package main

import (
	"testing"

	"github.com/yeonghoey/cmdtest"
)

func TestSamples(t *testing.T) {
	cmd := cmdtest.Command("go", "run", ".")

	samples := []struct {
		input  string
		output string
	}{
		{"input1.txt", "output1.txt"},
	}

	for _, sa := range samples {
		got, want, err := cmd.Run(sa.input, sa.output)

		if err != nil {
			t.Errorf("cmd.Run failed: %v", err)
		}
		if got != want {
			t.Errorf("cmd.Run(%s) = %s, want %s",
				sa.input, got, want)
		}
	}
}
```
