package fileHandler

import (
	"bufio"
	"crud-go-project/internal/model"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (DefaultFileHandler) ReadByLine(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	value := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value = append(value, scanner.Text())
	}
	return value
}

func (DefaultFileHandler) ReadFileToStruct(filename string) []model.User {
	value := make([]model.User, 0)
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		raw := strings.Fields(scanner.Text())
		passwordNumber, _ := strconv.Atoi(raw[1])
		aux := model.User{
			UserName: raw[0],
			Password: passwordNumber,
		}
		value = append(value, aux)
	}
	return value
}
