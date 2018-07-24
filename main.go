package goterraform

type TerraformCli struct {
	path             string
	version          string
	workingDirectory string
}

// NewTerraformClient return a struct which behaves like the cli terraform client.
//
// Between the unreliability of the internal interfaces in the terraform library and then
// need to communicate with providers, we'll wrap the terraform command in bash, rather
// than importing the `github.com/hashicorp/terraform` library and calling methods
// directly. See https://github.com/hashicorp/terraform/issues/12582 for more info.
func NewTerraformClient() *TerraformCli {
	return NewTerraformClientWithBinPath("terraform")
}

func NewTerraformClientWithBinPath(binPath string) *TerraformCli {
	cli := &TerraformCli{
		path: binPath,
	}
	cli.fetchVersion()
	return cli
}

func (t *TerraformCli) Plan(params *TerraformPlanParams) *TerraformAction {
	return &TerraformAction{
		action: "plan",
		bin:    t,
		params: params,
	}
}

func (t *TerraformCli) Apply() *TerraformAction {
	return &TerraformAction{
		action: "apply",
		bin:    t,
	}
}

func (t *TerraformCli) fetchVersion() {
	t.version = "dev"
}
func (client *TerraformCli) WithWorkingDirectory(workingDirectory string) (*TerraformCli) {
	client.workingDirectory = workingDirectory
	return client
}
