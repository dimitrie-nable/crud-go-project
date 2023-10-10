package http

import (
	"crud-go-project/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer(usersAndPasswords []model.User, userNames []string) {
	router := gin.Default()

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
