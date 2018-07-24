package goterraform

import (
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewTerraformPlanParams(t *testing.T) {
	tests := []struct {
		name string
		want *TerraformPlanParams
	}{
		{
			name: "base",
			want: &TerraformPlanParams{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTerraformPlanParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTerraformPlanParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTerraformPlanParams_Opts(t *testing.T) {
	tests := []struct {
		name       string
		p          *TerraformPlanParams
		want       map[string][]string
		wantString string
	}{
		{
			name: "base",
			p:    &TerraformPlanParams{},
			want: map[string][]string{},
		},
		{
			name: "output",
			p: &TerraformPlanParams{
				Out: StringPtr("/mock/path"),
			},
			want: map[string][]string{
				"out": {"/mock/path"},
			},
			wantString: "-out=/mock/path",
		},
		{
			name: "1-target",
			p: &TerraformPlanParams{
				Target: StringSlicePtr([]string{"mock-resource1"}),
			},
			want: map[string][]string{
				"target": {"mock-resource1"},
			},
			wantString: "-target=mock-resource1",
		},
		{
			name: "2-target",
			p: &TerraformPlanParams{
				Target: StringSlicePtr([]string{"mock-resource1", "mock-resource2"}),
			},
			want: map[string][]string{
				"target": {"mock-resource1", "mock-resource2"},
			},
			wantString: "-target=mock-resource1 -target=mock-resource2",
		},
		{
			name: "true-bool-ptr",
			p: &TerraformPlanParams{
				Lock: BoolPtr(true),
			},
			want: map[string][]string{
				"lock": {"true"},
			},
			wantString: "-lock=true",
		},
		{
			name: "all",
			p: &TerraformPlanParams{
				Destroy:          true,
				DetailedExitcode: true,
				Input:            FalsePtr(),
				Lock:             FalsePtr(),
				LockTimeout:      1,
				ModuleDepth:      IntPtr(2),
				NoColor:          true,
				Out:              StringPtr("/mock/path"),
				Parallelism:      IntPtr(2),
				Refresh:          FalsePtr(),
				State:            StringPtr("mock-statefile.tfstate"),
				Target:           StringSlicePtr([]string{"mock-resource1"}),
				Var:              StringMapPtr(map[string]string{"foo": "bar", "hello": "world"}),
				VarFile:          StringSlicePtr([]string{"vars1", "vars2"}),
			},
			want: map[string][]string{
				"destroy":           {""},
				"detailed-exitcode": {""},
				"input":             {"false"},
				"lock":              {"false"},
				"lock-timeout":      {"1s"},
				"module-depth":      {"2"},
				"no-color":          {""},
				"out":               {"/mock/path"},
				"parallelism":       {"2"},
				"refresh":           {"false"},
				"state":             {"mock-statefile.tfstate"},
				"target":            {"mock-resource1"},
				"var":               {"foo=bar", "hello=world"},
				"var-file":          {"vars1", "vars2"},
			},
			wantString: "-destroy -detailed-exitcode -input=false -lock-timeout=1s -lock=false -module-depth=2 " +
				"-no-color -out=/mock/path -parallelism=2 -refresh=false -state=mock-statefile.tfstate " +
				"-target=mock-resource1 -var 'foo=bar' -var 'hello=world' -var-file=vars1 -var-file=vars2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.p.Opts())
			assert.Equal(t, tt.wantString, tt.p.OptsString())
		})
	}
}
