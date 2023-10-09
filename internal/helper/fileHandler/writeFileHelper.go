package fileHandler

import (
	"crud-go-project/internal/constant"
	"crud-go-project/internal/helper/passwordHandler"
	"fmt"
	"os"
)

func WriteFileByLines(filename string, userNames []string) {
	lines := passwordHandler.FileStructureGenerator(constant.MIN_PASSWORD_VALUE, constant.MAX_PASSWORD_VALUE, userNames)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	for _, val := range lines {
		_, err := fmt.Fprintln(file, val)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
