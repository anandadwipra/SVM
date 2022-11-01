package svm_cli

import (
	"fmt"

	svm_cloudinit "github.com/anandadwipra/SVM/internal/cloud_init"
)

func (profile Svm) Create(option Options) {
	fmt.Println("Running Create")
	profile.profile.Read()
	svm_cloudinit.New(profile.profile, option.name)

}
