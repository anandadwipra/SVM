package svm_cli

import (
	"fmt"
	"os"
	"os/exec"
)

func (profile Svm) Img(option Options) {
	profile.profile.Read()
	path := fmt.Sprintf("%s/isos", profile.profile.Workdir)
	os.Chdir(path)
	command := fmt.Sprintf("wget %s", option.url)
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Run()

}
