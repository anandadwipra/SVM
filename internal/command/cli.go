package svm_cli

import (
	"flag"
	"fmt"
	"os"

	svm_profile "github.com/anandadwipra/SVM/internal/profile"
	"github.com/posener/complete/v2"
)

// type Config struct {
// 	ssh_path, workdir string
// }

// the option should be in its own internal package?
type Options struct {
	cmd                                                       *complete.Command
	name, base_image, network1, network2, network3, user, url string
	cpu, mem, storage                                         int
	all, off                                                  bool
	create, delete, initialize, profile, list, img            *flag.FlagSet
	config                                                    *svm_profile.Profile
}

type Svm struct {
	profile svm_profile.Profile
}

func (data Options) desc() {
	fmt.Println(`Usage: vm [OPTIONS] argument ... 
    Available Options:
      create
      delete
      init 
      img
      list
      profile 
    Get Help:
      vm [OPTIONS] --help`)
	os.Exit(1)
}

// Get svm Command and Argument
func (data *Options) Get() {
	var svm Svm
	data.auto_complete()
	// fmt.Println("Masuk")
	if len(os.Args) < 2 {
		// fmt.Println("Info ni boss").
		data.cmd.Complete("command")
		data.desc()
	}
	switch os.Args[1] {
	// Create Command
	case "create":
		data.Create()
		data.create.Parse(os.Args[2:])
		if data.name == "" {
			fmt.Println("Usage of create:")
			data.create.PrintDefaults()
		} else {
			svm.Create(*data)
		}

		// Delete Command
	case "delete":
		data.Delete()
		data.delete.Parse(os.Args[2:])
		if data.name == "" {
			fmt.Println("Usage of delete:")
			data.delete.PrintDefaults()
		} else {
			svm.Delete()
		}

		// Init Command
	case "init":
		data.config = &svm.profile
		data.Initialize()
		data.initialize.Parse(os.Args[2:])
		fmt.Println(data.config)
		// data.initialize.Parse(os.Args[2:])
		if data.config.Ssh_path == "-" {
			Initialize(&svm.profile)
		} else if data.config.Workdir != "/data" || data.config.Default_user != "ubuntu" && data.config.Ssh_path == "-" {
			fmt.Println("Usage of delete:")
			data.initialize.PrintDefaults()
		} else {
			data.config.Write("default")
		}

		// List Command
	case "list":
		data.List()
		data.list.Parse(os.Args[2:])
		if data.config.Ssh_path == "" {
			fmt.Println("Usage of delete:")
			data.list.PrintDefaults()
		} else {
			svm.List()
		}

		// Profile Command
	case "profile":
		data.Profile()
		data.profile.Parse(os.Args[2:])
		if data.config.Ssh_path == "" {
			fmt.Println("Usage of delete:")
			data.profile.PrintDefaults()
		} else {
			svm.Profile()
		}
	case "img":
		data.Img()
		data.img.Parse(os.Args[2:])
		if data.url == "" {
			fmt.Println("Usage of image:")
			data.img.PrintDefaults()
		} else {
			svm.Img(*data)
		}
	default:
		// fmt.Println("Default ni boss")
		data.desc()
	}
}

// Parse Create Argument to variable
func (data *Options) Create() {
	data.create = flag.NewFlagSet("create", flag.ExitOnError)
	data.create.StringVar(&data.name, "name", "", "Virtual machine name")
	data.create.StringVar(&data.base_image, "image", "ubuntu20.img", "Define Virtual machine base image")
	data.create.StringVar(&data.network1, "network[n]", "default", "Define Virtual machine network interface")
	data.create.IntVar(&data.cpu, "cpu", 1, "Virtual machine vcpus")
	data.create.IntVar(&data.mem, "mem", 1024, "Virtual machine Memory")
	data.create.IntVar(&data.storage, "storage", 20, "Virtual machine Storage in GB")
}

// Parse Delete Argument to variable
func (data *Options) Delete() {
	data.delete = flag.NewFlagSet("delete", flag.ExitOnError)
	data.delete.StringVar(&data.name, "name", "", "Virtual machine name")
}

// Parse Init Argument to variable
func (data *Options) Initialize() {
	data.initialize = flag.NewFlagSet("init", flag.ExitOnError)
	data.initialize.StringVar(&data.config.Ssh_path, "ssh_path", "-", "Default ssh for machine (Required)")
	data.initialize.StringVar(&data.config.Workdir, "workdir", "/data", "VM Directory")
	data.initialize.StringVar(&data.config.Default_user, "default_user", "ubuntu", "Default user for VM")

}

// Parse List Argument to variable
func (data *Options) List() {
	data.list = flag.NewFlagSet("list", flag.ExitOnError)

	// fmt.Printf("VM User Config:\n\tWorkdir : %s\n\tSSH: %s\n", data.config.workdir, data.config.ssh_path)
}

// Parse Profile Argument to variable
func (data *Options) Profile() {
	data.profile = flag.NewFlagSet("profile", flag.ExitOnError)
}

// Parse Img Argument to variable
func (data *Options) Img() {
	data.img = flag.NewFlagSet("profile", flag.ExitOnError)
	data.img.StringVar(&data.url, "url", "", "Url to download image")
}

// func (data *Options) Switch() {

// }
