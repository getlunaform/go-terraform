package goterraform

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAction_Run(t *testing.T) {
	tests := []struct {
		name    string
		a       *TerraformAction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.Run(); (err != nil) != tt.wantErr {
				t.Errorf("TerraformAction.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTruePtr(t *testing.T) {
	want := true
	assert.Equal(t, &want, TruePtr())
}
