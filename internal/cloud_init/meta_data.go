package svm_cloudinit

import (
	"fmt"
	"os"
	"strings"

	svm_profile "github.com/anandadwipra/SVM/internal/profile"
	file "github.com/anandadwipra/SVM/pkg/general"
)

func init() {
	// fmt.Println("Running Cloud Init")
}

func metadata(path, hostname string) {
	file.CreateFile(path, "meta-data")
	file, err := os.OpenFile(path+"/meta-data", os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	data := string(fmt.Sprintf("local-hostname: %s\n", hostname))
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	file.Sync()
}

func userdata(path, hostname, user, sshpath string) {
	file.CreateFile(path, "user-data")

	file, err := os.OpenFile(path+"/user-data", os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	pub, errpub := os.ReadFile(sshpath)
	if errpub != nil {
		fmt.Println(err.Error())
	}
	pubstr := strings.Replace(string(pub), "\n", "", -1)
	data := fmt.Sprintf(`#cloud-config
users:
  - name: %s
    ssh-authorized-keys:
      - %s 
    sudo: ['ALL=(ALL) NOPASSWD:ALL'] 
    groups: sudo
    shell: /bin/bash
runcmd:
  - echo "AllowUsers %s" >> /etc/ssh/sshd_config
  - restart ssh`, user, pubstr, user)
	// fmt.Print(string(pub))

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	file.Sync()
}

func New(profile svm_profile.Profile, hostname string) {
	pathname := fmt.Sprintf("%s/vms/%s", profile.Workdir, hostname)
	os.MkdirAll(pathname, os.ModePerm)
	metadata(pathname, hostname)
	userdata(pathname, hostname, profile.Default_user, profile.Ssh_path)
}
