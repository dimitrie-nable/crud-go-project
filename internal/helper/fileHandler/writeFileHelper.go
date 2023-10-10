package fileHandler

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func WriteFileByLines(fileName string, lines []string) {
	file, err := os.Create(fileName)
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

func AppendToFile(fileName, data string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, fs.ModeAppend)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	_, fErr := fmt.Fprintln(file, data)

	if fErr != nil {
		fmt.Println(fErr.Error())
		return
	}
}

func DeleteFile(fileName string) bool {
	err := os.Remove(fileName)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func DeleteLineFromFile(fileName, userName string) bool {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer file.Close()

	tempFileName := fileName + ".tmp"
	tempFile, err1 := os.Create(tempFileName)

	if err1 != nil {
		fmt.Println(err1.Error())
		return false
	}

	defer tempFile.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		raw := strings.Fields(line)

		if len(raw) == 2 {
			if raw[0] != userName {
				_, err2 := tempFile.WriteString(line + "\n")

				if err2 != nil {
					fmt.Println(err2.Error())
					return false
				}
			}

		}
	}

	errScan := scanner.Err()

	if errScan != nil {
		fmt.Println(err.Error())
		return false
	}

	file.Close()
	tempFile.Close()

	errScan = os.Rename(tempFileName, fileName)

	if errScan != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
