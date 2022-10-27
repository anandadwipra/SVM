package initialize

import (
	"fmt"
	"os"
)

// Create Working directory
func CheckDir(path_workdir *string) {
	var err error
	if string((*path_workdir)[len(*path_workdir)-1]) != "/" {
		*path_workdir = *path_workdir + "/"
		fmt.Println(*path_workdir)
	}

	for _, subpath := range []string{"vms", "isos"} {
		err = os.MkdirAll(*path_workdir+subpath, 0750)

		if err != nil {
			if err.Error() == "mkdir /data: permission denied" {
				panic("Permission denied, Please run with sudoer user or change your workdir")
			} else {
				fmt.Println(err.Error())
			}
		}
	}
	fmt.Println("Working Directory Ready")
}
