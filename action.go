package goterraform

import (
	"os/exec"
	"strings"
	"bufio"
	"fmt"
	"sort"
)

type TerraformActionParams interface {
	Opts() map[string][]string
	OptsString() string
	OptsStringSlice() []string
}

type TerraformAction struct {
	Cmd    *exec.Cmd
	Dir    string
	out    *TerraformOutput
	action string
	bin    *TerraformCli
	opts   map[string]string
	logs   *OutputLog
	params TerraformActionParams
}

func (a *TerraformAction) Init() *TerraformAction {
	args := append([]string{a.action}, a.params.OptsStringSlice()...)
	fmt.Printf("%s\n", args)
	a.Cmd = exec.Command(a.bin.path, args...)
	if a.Dir != "" {
		a.Cmd.Dir = a.Dir
	}
	a.out = &TerraformOutput{}
	return a
}

// Run the terraform command
func (a *TerraformAction) Run() (err error) {
	return a.Cmd.Start()
}

func (a *TerraformAction) InitLogger(log *OutputLog) (err error) {
	a.logs = log

	// Configure stdout capture
	if a.out.Stdout, err = a.Cmd.StdoutPipe(); err != nil {
		return
	}
	scannerStdout := bufio.NewScanner(a.out.Stdout)
	go func() {
		for scannerStdout.Scan() {
			fmt.Print(
				a.logs.Stdout(scannerStdout.Text()).String() + "\n",
			)
		}
	}()

	// Configure stderr capture
	if a.out.Stderr, err = a.Cmd.StderrPipe(); err != nil {
		return
	}
	scannerStderr := bufio.NewScanner(a.out.Stderr)
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

func extractOptsStringSlice(p TerraformActionParams) (options []string) {
	opts := p.Opts()
	keys := mapStringSliceKeys(opts)
	sort.Strings(keys)

	outputs := make([]string, 0)
	for _, key := range keys {
		value := opts[key]
		sort.Strings(value)
		for _, val := range value {
			output := "-" + key
			if val != "" {
				switch key {
				case "var":
					outputs = append(outputs, output)
					outputs = append(outputs, "'"+val+"'")
					continue
				default:
					output = output + "=" + val
				}
			}
			outputs = append(outputs, output)
		}
	}
	return outputs
}

func extractOptsString(p TerraformActionParams) (options string) {
	return strings.Join(
		extractOptsStringSlice(p),
		" ",
	)
}

func mapStringSliceKeys(s map[string][]string) (keys []string) {
	keys = make([]string, len(s))

	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return
}
