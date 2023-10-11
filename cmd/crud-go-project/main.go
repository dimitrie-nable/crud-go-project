package main

import (
	"crud-go-project/internal/constant"
	"crud-go-project/internal/helper/fileHandler"
	"crud-go-project/internal/http"
	"crud-go-project/internal/model"
)

var userNames = make([]string, 0)
var usersAndPasswords = make([]model.User, 0)

func init() {
	userNames = fileHandler.ReadByLine(constant.USERNAMES_PATH)
	usersAndPasswords = fileHandler.ReadFileToStruct(constant.USERS_AND_PASSWORDS_PATH)
}

func main() {

	http.StartServer(usersAndPasswords, userNames)

}
