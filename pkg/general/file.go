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
			fmt.Printf("Failed to create file: %s", err.Error())
		}
	}
}
