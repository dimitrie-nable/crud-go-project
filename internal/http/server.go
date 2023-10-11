package http

import (
	"crud-go-project/internal/db"
	"crud-go-project/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer(usersAndPasswords []model.User, userNames []string) {
	dbInstance := db.ConnectToDatabase()

	router := gin.Default()

	router.GET("/userDatabase", GetUsersFromDatabase(dbInstance))

	router.POST("/userDatabase", AddUserToDatabase(dbInstance))

	router.DELETE("/userDatabase/:userName", DeleteUsersAndPasswordsFromDatabase(dbInstance))

	router.PUT("/userDatabase/:userName", UpdatePasswordOfUserInDatabase(dbInstance))

	router.GET("/dummyUser", GenerateUsersAndPasswords(usersAndPasswords, userNames))

	router.GET("/user", GetUsers(usersAndPasswords))

	router.GET("/passwordsOfUsers", GetPasswords(usersAndPasswords))

	router.POST("/user", CreateUserFile(usersAndPasswords))

	router.PUT("/user", AddUser(usersAndPasswords))

	router.DELETE("/user", DeleteUsersAndPasswordsFile())

	router.DELETE("/user/:userName", DeleteUsersAndPasswordsFromFile(usersAndPasswords))

	err := router.Run()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
