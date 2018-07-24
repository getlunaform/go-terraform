package goterraform

import (
	"io"
	"github.com/fatih/color"
	"fmt"
)

type TerraformOutput struct {
	Stderr io.ReadCloser
	Stdout io.ReadCloser
}

const (
	STDERR = "stderr"
	STDOUT = "stdout"
)

func NewOutputLogs() *OutputLog {
	return &OutputLog{
		Entries: make([]*OutputLogEntry, 0),
	}
}

func (ol *OutputLog) Stdout(message string) *OutputLogEntry {
	return ol.Append(&OutputLogEntry{
		Type:    STDOUT,
		Content: message,
	})
}

func (ol *OutputLog) Stderr(message string) *OutputLogEntry {
	return ol.Append(&OutputLogEntry{
		Type:    STDERR,
		Content: message,
	})
}

func (ol *OutputLog) Append(ole *OutputLogEntry) *OutputLogEntry {
	ol.Entries = append(ol.Entries, ole)
	return ole
}

func (ol *OutputLog) String() (output string) {
	output = ""
	for _, entry := range ol.Entries {
		output = output + entry.String() + "\n"
	}
	return
}

type OutputLog struct {
	Entries []*OutputLogEntry `json:"entries"`
}

type OutputLogEntry struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func (ole *OutputLogEntry) String() string {
	var prefix string
	if ole.Type == STDERR {
		prefix = color.RedString("[stderr]")
	} else if ole.Type == STDOUT {
		prefix = "[stdout]"
	}
	return fmt.Sprintf("%s %s", prefix, ole.Content)
}
