package cd

const ConfigFileName = "server-cd.yml"

type Config struct {
	Name         string `yaml:"name"`
	DeployBranch string `yaml:"deploy-branch"`
	RunCommand   string `yaml:"run-command"`
}