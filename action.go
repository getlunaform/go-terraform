package goterraform

import (
	"os/exec"
	"time"
)

type TerraformAction struct {
	cmd    *exec.Cmd
	out    *TerraformOutput
	action string
	bin    *TerraformCli
	opts   map[string]string
}

func (a *TerraformAction) Run() (err error) {
	a.cmd = exec.Command(a.bin.path + " " + a.action)
	time.Sleep(1000 * time.Millisecond)
	if a.out.Stderr, err = a.cmd.StderrPipe(); err != nil {
		return
	}
	if a.out.Stdout, err = a.cmd.StdoutPipe(); err != nil {
		return
	}

	return a.cmd.Start()
}

func BoolPtr(a bool) *bool {
	return &a
}

func TruePtr() *bool {
	return BoolPtr(true)
}

func FalsePtr() *bool {
	return BoolPtr(false)
}

func StringPtr(a string) *string {
	return &a
}

func IntPtr(a int) *int {
	return &a
}

func StringSlicePtr(a []string) *[]string {
	return &a
}

func StringMapPtr(a map[string]string) *map[string]string {
	return &a
}
