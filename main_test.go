package goterraform

import (
	"reflect"
	"testing"
)

func TestNewTerraformClient(t *testing.T) {
	tests := []struct {
		name string
		want *TerraformCli
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTerraformClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTerraformClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTerraformClientWithBinPath(t *testing.T) {
	type args struct {
		binPath string
	}
	tests := []struct {
		name string
		args args
		want *TerraformCli
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTerraformClientWithBinPath(tt.args.binPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTerraformClientWithBinPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTerraformCli_Plan(t *testing.T) {
	tests := []struct {
		name string
		t    *TerraformCli
		want *TerraformAction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Plan(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TerraformCli.Plan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTerraformCli_Apply(t *testing.T) {
	tests := []struct {
		name string
		t    *TerraformCli
		want *TerraformAction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Apply(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TerraformCli.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTerraformCli_fetchVersion(t *testing.T) {
	tests := []struct {
		name string
		t    *TerraformCli
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.fetchVersion()
		})
	}
}
