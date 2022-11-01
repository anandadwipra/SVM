package file

import (
	"fmt"
	"os"
)

func CreateFile(path, filename string) {
	_, err := os.Stat(path + "/" + filename)
	if os.IsNotExist(err) {
		file, err := os.Create(path + "/" + filename)
		defer file.Close()
		if err != nil {
			fmt.Println("Error Bikin file:", err.Error())
		}
	}
	// fmt.Println(err)

}
