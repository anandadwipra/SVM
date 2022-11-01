package svm_profile

type Profile struct {
	Name         string `mapstructure:"name"`
	Ssh_path     string `mapstructure:"ssh_path"`
	Workdir      string `mapstructure:"workdir"`
	Default_user string `mapstructure:"default_user"`
	Uuid         []byte `mapstructure:"uuid"`
}
