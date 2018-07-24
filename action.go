package goterraform

import (
	"os/exec"
	"sort"
	"strings"
	"bufio"
	"fmt"
)

type TerraformActionI interface {
	Opts() map[string][]string
	OptsString() string
}

type TerraformAction struct {
	Cmd    *exec.Cmd
	out    *TerraformOutput
	action string
	bin    *TerraformCli
	opts   map[string]string
	logs   *OutputLog
}

func (a *TerraformAction) Init() *TerraformAction {
	a.Cmd = exec.Command(a.bin.path, a.action)
	a.out = &TerraformOutput{}
	return a
}

func (a *TerraformAction) Run() (err error) {

	return a.Cmd.Start()
}
func (a *TerraformAction) InitLogger(log *OutputLog) (err error) {
	a.logs = log

	if a.out.Stdout, err = a.Cmd.StdoutPipe(); err != nil {
		return
	}
	if a.out.Stderr, err = a.Cmd.StderrPipe(); err != nil {
		return
	}

	scannerStdout := bufio.NewScanner(a.out.Stdout)
	scannerStderr := bufio.NewScanner(a.out.Stderr)
	go func() {
		for scannerStdout.Scan() {
			fmt.Print(
				a.logs.Stdout(scannerStdout.Text()).String() + "\n",
			)
		}
	}()

	go func() {
		for scannerStderr.Scan() {
			fmt.Print(
				a.logs.Stderr(scannerStderr.Text()).String() + "\n",
			)
		}
	}()
	return
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

func extractOptsString(p TerraformActionI) (options string) {
	opts := p.Opts()
	outputs := make([]string, 0)
	for key, value := range opts {
		for _, val := range value {
			output := "-" + key
			if val != "" {
				switch key {
				case "var":
					output = output + " '" + val + "'"
				default:
					output = output + "=" + val
				}
			}
			outputs = append(outputs, output)
		}
	}
	sort.Strings(outputs)
	return strings.Join(outputs, " ")
}
