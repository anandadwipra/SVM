package areader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(echo string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(echo)
	result, err := reader.ReadString('\n')
	result = strings.Replace(result, "\n", "", -1)
	return result, err
}
