package svm_cli

import (
	"fmt"
	"os/exec"

	svm_cloudinit "github.com/anandadwipra/SVM/internal/cloud_init"
)

func (profile Svm) Create(option Options) {
	fmt.Println("Running Create")
	profile.profile.Read()
	pathname := svm_cloudinit.New(profile.profile, option.name)
	command := fmt.Sprintf("genisoimage -ouput %s/%s-cidata-iso -vloid cidata -joiler -rock", pathname, option.name)
	cmd := exec.Command("/bin/sh", string(command))
	cmd.Run()
}
