package cmdline

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRepl(t *testing.T) {
	type Expectation struct {
		Error      string
		TestCalled bool
		Out        string
	}
	tests := []struct {
		Name        string
		Expectation Expectation
		Input       string
	}{
		{
			Name:  "test command",
			Input: "test\n",
			Expectation: Expectation{
				TestCalled: true,
				Out:        "> \n> \n",
			},
		},
		{
			Name:  "command not found",
			Input: "unknown\n",
			Expectation: Expectation{
				Out: "> \nunknown command: unknown\n> \n",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var act Expectation

			commands := map[string]CommandFunc{
				"test": func(args []string) error {
					act.TestCalled = true
					return nil
				},
				"error": func(args []string) error {
					return fmt.Errorf("foo")
				},
			}

			out := bytes.NewBuffer(nil)
			stdout = out
			stdin = bytes.NewReader([]byte(test.Input))
			Repl(commands)
			act.Out = out.String()

			if diff := cmp.Diff(test.Expectation, act); diff != "" {
				t.Errorf("Repl() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
