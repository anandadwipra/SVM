package initialize

import (
	"errors"
	"fmt"

	"github.com/arduino/go-apt-client"
)

type Package_List struct {
	List    *[]apt.Package
	Version bool
}

// create New *Package_List
// Ex: NewPackage([]string{"nano","1.0.0"})
func NewPackage(package_lists ...map[string]string) (package_list Package_List) {
	for i, list := range package_lists {
		if i == 0 {
			package_list.List = &[]apt.Package{apt.Package{Name: list["Name"], Version: list["Version"]}}
		} else {
			*package_list.List = append(*package_list.List, apt.Package{Name: list["Name"], Version: list["Version"]})
		}
		//dor debuggin
		// fmt.Println(*package_list.List)
	}
	package_list.Version = false
	return
}

// Check installation status of list package
// Warning Sensitive char
func (package_list *Package_List) Check() (map[string][]*apt.Package, error) {
	var not_installed, installed []*apt.Package
	for _, package_check_req := range *package_list.List {
		fmt.Println("=>", package_check_req)
		search_ress, _ := apt.Search(package_check_req.Name)

		for _, search_res := range search_ress {
			if search_res.Name == package_check_req.Name {
				if search_res.Status == "installed" {
					fmt.Println("installed")
					installed = append(installed, search_res)
				} else {
					fmt.Println("NotInstalled")
					not_installed = append(not_installed, search_res)
				}
			}
		}

		if len(search_ress) == 0 {
			panic(errors.New("Cannot Find Package," + package_check_req.Name + " Please check Package name or your repository"))
			// return map[string][]string{"Installed": installed, "NotInstalled": not_installed},
			// not_installed = append(not_installed, package_check_req.Name)
		}
	}
	return map[string][]*apt.Package{"Installed": installed, "NotInstalled": not_installed}, nil
}

// Install package of list package
func (package_list *Package_List) Install() ([]string, error) {
	pacakge_needs, _ := package_list.Check()
	for _, package_need := range pacakge_needs["NotInstalled"] {
		fmt.Println(apt.Install(package_need))
	}
	// fmt.Println(package_need_install["NotInstalled"][0])
	return nil, nil
}

// Add Repository
// func (package_list *Package_List) AddRepo(){

// }
