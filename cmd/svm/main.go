package main

import (
	svm_cli "github.com/anandadwipra/SVM/internal/command"
)

func main() {
	// type str_str map[string]string
	// packageList := []map[string]string{
	// 	str_str{"Name": "nano", "Version": "1.0.0"},
	// 	str_str{"Name": "vim", "Version": "2.0.0"},
	// 	// str_str{"Name": "xvim", "Version": "2.0.0"},
	// }
	// package_dep := initialize.NewPackage(packageList...)
	// // package_dep.Install()
	// data, _ := package_dep.Check()
	// fmt.Println(data)

	var data svm_cli.Options

	data.Get()

}
