package cmdline

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	stdin  io.Reader = os.Stdin
	stdout io.Writer = os.Stdout
)

type CommandFunc func(args []string) error

// Repl reads commands from stdin and executes them
func Repl(commands map[string]CommandFunc) {
	fmt.Fprint(stdout, "> ")

	pr, pw := io.Pipe()
	go io.Copy(stdout, pr)
	scan := bufio.NewScanner(io.TeeReader(stdin, pw))
	for scan.Scan() {
		line := scan.Text()
		segs := strings.Split(line, " ")
		cmd, args := segs[0], segs[1:]
		if cmd == "" {
			fmt.Fprint(stdout, "> ")
			continue
		}

		if cmd == "exit" {
			return
		}

		if f, ok := commands[cmd]; ok {
			if err := f(args); err != nil {
				fmt.Fprintln(stdout, "error:", err.Error())
			}
		} else {
			fmt.Fprintln(stdout, "unknown command:", cmd)
		}

		fmt.Fprint(stdout, "> ")
	}
}
