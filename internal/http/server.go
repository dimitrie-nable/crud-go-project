package http

import (
	"crud-go-project/internal/db"
	"crud-go-project/internal/helper/fileHandler"
	"crud-go-project/internal/helper/passwordHandler"
	"crud-go-project/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer(usersAndPasswords []model.User, userNames []string) {

	fileHandlerDefault := fileHandler.DefaultFileHandler{}
	fileHandler := NewFileHandler(fileHandlerDefault)

	passwordHandlerDefault := passwordHandler.DefaultPasswordHandler{}
	passwordHandler := NewPasswordHandler(passwordHandlerDefault)

	dbInstance := db.ConnectToDatabase()

	router := gin.Default()

	router.GET("/userDatabase", GetUsersFromDatabase(dbInstance))

	router.POST("/userDatabase", AddUserToDatabase(dbInstance))

	router.DELETE("/userDatabase/:userName", DeleteUsersAndPasswordsFromDatabase(dbInstance))

	router.PUT("/userDatabase/:userName", UpdatePasswordOfUserInDatabase(dbInstance))

	router.GET("/dummyUser", fileHandler.GenerateUsersAndPasswords(usersAndPasswords, userNames))

	router.GET("/user", fileHandler.GetUsers())

	router.GET("/passwordsOfUsers", passwordHandler.GetPasswords(usersAndPasswords))

	router.POST("/user", fileHandler.CreateUserFile(usersAndPasswords))

	router.PUT("/user", fileHandler.AddUser(usersAndPasswords))

	router.DELETE("/user", fileHandler.DeleteUsersAndPasswordsFile())

	router.DELETE("/user/:userName", fileHandler.DeleteUsersAndPasswordsFromFile(usersAndPasswords))

	err := router.Run()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
