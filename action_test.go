package goterraform

import "testing"

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
