package main

import (
	"crud-go-project/internal/constant"
	"crud-go-project/internal/helper/fileHandler"
	"crud-go-project/internal/helper/passwordHandler"
	"crud-go-project/internal/model"
	"fmt"
)

var userNames = make([]string, 0)
var users = make([]model.User, 0)

func init() {
	userNames = fileHandler.ReadByLine(constant.USERNAMES_PATH)
	users = fileHandler.ReadFileToStruct(constant.USERS_AND_PASSWORDS_PATH)
}

func main() {

	var selection int
	runApp := true

	for runApp {
		fmt.Println("Introduce the number of the action that you want to take:")
		fmt.Println("1. Write a file with some dummy data with users and passwords")
		fmt.Println("2. Read from the file with users and passwords")
		fmt.Println("3. Crack the passwords of the users")
		fmt.Println("4. Close the application")

		fmt.Scan(&selection)

		isValid := selection >= 1

		if isValid {
			fmt.Println("Your choice is valid!")
			switch selection {
			case 1:
				fileHandler.WriteFileByLines(constant.USERS_AND_PASSWORDS_PATH, userNames)
			case 2:
				fmt.Println(fileHandler.ReadFileToStruct(constant.USERS_AND_PASSWORDS_PATH))
			case 3:
				passwordHandler.GetAllPasswords(users)
			case 4:
				runApp = false
			}
		} else {
			fmt.Println("Bad selection!")
		}
	}

}
