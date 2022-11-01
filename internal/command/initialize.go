package svm_cli

import (
	"fmt"

	svm_profile "github.com/anandadwipra/SVM/internal/profile"
	"github.com/anandadwipra/SVM/pkg/initialize"
	areader "github.com/anandadwipra/SVM/pkg/input"
)

func Initialize(profile *svm_profile.Profile) {
	fmt.Println("Running Initialize")
	profile.Name = "Default"
	profile.Workdir, _ = areader.Input("Chose Workdir : ")
	profile.Ssh_path, _ = areader.Input("Define your SSH Pub Key(PATH) : ")
	profile.Default_user, _ = areader.Input("VM Default user : ")
	// fmt.Println("Workdir", profile.Workdir, "Created")
	// fmt.Println("SSH", profile.Ssh_path, "Added")
	// fmt.Println("User", profile.Default_user, "Added")
	profile.Write("Default Configuration")
	initialize.CheckDir(&profile.Workdir)

}
