package goterraform

import "strconv"

type TerraformInitParams struct {
	Backend       *bool
	BackendConfig string
	ForceCopy     bool
	FromModule    string
	Get           *bool
	GetPlugins    *bool

	Input       *bool
	Lock        *bool
	LockTimeout int
	NoColor     bool

	PluginDir     string
	Reconfigure   bool
	Upgrade       *bool
	VerifyPlugins *bool
}

func NewTerraformInitParams() *TerraformInitParams {
	return &TerraformInitParams{}
}

func (p *TerraformInitParams) Opts() map[string][]string {
	opts := make(map[string][]string)

	if p.Backend != nil && *p.Backend == false {
		opts["backend"] = []string{"false"}
	}

	if p.BackendConfig != "" {
		opts["backend-config"] = []string{p.BackendConfig}
	}

	if p.ForceCopy {
		opts["force-copy"] = []string{""}
	}

	if p.FromModule != "" {
		opts["from-module"] = []string{p.FromModule}
	}

	if p.Get != nil && *p.Get == false {
		opts["get"] = []string{"false"}
	}

	if p.GetPlugins != nil && *p.GetPlugins == false {
		opts["get-plugins"] = []string{"false"}
	}

	if p.Input != nil && *p.Input == false {
		opts["input"] = []string{"false"}
	}

	if p.Lock != nil {
		if *p.Lock == true {
			opts["lock"] = []string{"true"}
		} else {
			opts["lock"] = []string{"false"}
		}
	}

	if p.LockTimeout != 0 {
		opts["lock-timeout"] = []string{strconv.Itoa(p.LockTimeout) + "s"}
	}

	if p.NoColor {
		opts["no-color"] = []string{""}
	}

	if p.PluginDir != "" {
		opts["plugin-dir"] = []string{p.PluginDir}
	}

	if p.Reconfigure {
		opts["reconfigure"] = []string{""}
	}

	if p.Upgrade != nil && *p.Upgrade == true {
		opts["upgrade"] = []string{"true"}
	}

	if p.VerifyPlugins != nil && *p.VerifyPlugins == false {
		opts["verify-plugins"] = []string{"false"}
	}

	return opts
}

func (p *TerraformInitParams) OptsString() string {
	return extractOptsString(p)
}

func (p *TerraformInitParams) OptsStringSlice() []string {
	return extractOptsStringSlice(p)
}
