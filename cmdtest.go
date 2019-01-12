package cmdtest

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Cmd keeps the arguments for exec.Command and some options.
type Cmd struct {
	name      string
	args      []string
	TrimSpace bool
}

// Command returns the Cmd struct for reusing the arguments of exec.Command.
func Command(name string, arg ...string) *Cmd {
	return &Cmd{
		name:      name,
		args:      arg,
		TrimSpace: true,
	}
}

// Run starts the command with inputName as Stdin and returns
// the Stdout as string, along with the content of wantOutputName.
func (c *Cmd) Run(inputName, wantOutputName string) (got, want string, err error) {
	got, err = c.runToGetOutput(inputName)
	if err != nil {
		return
	}
	want, err = c.readWantOutput(wantOutputName)
	return
}

func (c *Cmd) runToGetOutput(inputName string) (got string, err error) {
	input, err := os.Open(inputName)
	if err != nil {
		return
	}
	defer func() { _ = input.Close() }()

	var buffer bytes.Buffer

	// Run the command with the input
	cmd := exec.Command(c.name, c.args...)
	cmd.Stdin = input
	cmd.Stdout = &buffer

	err = cmd.Run()

	got = buffer.String()
	if c.TrimSpace {
		got = strings.TrimSpace(got)
	}
	return
}

func (c *Cmd) readWantOutput(wantOutputName string) (want string, err error) {
	content, err := ioutil.ReadFile(wantOutputName)
	if err != nil {
		return
	}

	want = string(content)
	if c.TrimSpace {
		want = strings.TrimSpace(want)
	}
	return
}
