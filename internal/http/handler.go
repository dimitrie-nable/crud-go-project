package http

import (
	"crud-go-project/internal/constant"
	"crud-go-project/internal/helper/fileHandler"
	"crud-go-project/internal/helper/passwordHandler"
	"crud-go-project/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *FileHandler) GetUsers() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		usersAndPasswords := h.fileHandler.ReadFileToStruct(constant.USERS_AND_PASSWORDS_PATH)
		fmt.Println(usersAndPasswords)
		c.IndentedJSON(http.StatusOK, usersAndPasswords)
	}
	return gin.HandlerFunc(fn)
}

func (h *PasswordHandler) GetPasswords(usersAndPasswords []model.User) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		passwords := h.passwordHandler.GetAllPasswords(usersAndPasswords)
		fmt.Println(passwords)
		c.IndentedJSON(http.StatusOK, passwords)
	}
	return gin.HandlerFunc(fn)
}

func (h *FileHandler) CreateUserFile(usersAndPasswords []model.User) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var createUsersAndPasswords []model.User

		if err := c.BindJSON(&createUsersAndPasswords); err != nil {
			fmt.Println(err.Error())
			return
		}
		usersAndPasswords = nil
		fmt.Println(createUsersAndPasswords)
		for _, val := range createUsersAndPasswords {
			usersAndPasswords = append(usersAndPasswords, val)
		}

		var usersAndPasswordsString []string

		for _, val := range usersAndPasswords {
			usersAndPasswordsString = append(usersAndPasswordsString, val.UserName+" "+strconv.Itoa(val.Password))
		}
		fileHandler.WriteFileByLines(constant.USERS_AND_PASSWORDS_PATH, usersAndPasswordsString)
		c.IndentedJSON(http.StatusCreated, usersAndPasswords)
	}
	return gin.HandlerFunc(fn)
}

func (h *FileHandler) AddUser(users []model.User) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var newUser model.User

		if err := c.BindJSON(&newUser); err != nil {
			fmt.Println(err.Error())
			return
		}
		users = append(users, newUser)
		newUserString := newUser.UserName + " " + strconv.Itoa(newUser.Password)
		fileHandler.AppendToFile(constant.USERS_AND_PASSWORDS_PATH, newUserString)
		c.IndentedJSON(http.StatusCreated, newUser)
	}
	return gin.HandlerFunc(fn)
}

func (h *FileHandler) GenerateUsersAndPasswords(usersAndPasswords []model.User, userNames []string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		usersAndPasswordsString := passwordHandler.FileStructureGenerator(constant.MIN_PASSWORD_VALUE, constant.MAX_PASSWORD_VALUE, userNames)
		fileHandler.WriteFileByLines(constant.USERS_AND_PASSWORDS_PATH, usersAndPasswordsString)
		usersAndPasswords = h.fileHandler.ReadFileToStruct(constant.USERS_AND_PASSWORDS_PATH)
		c.IndentedJSON(http.StatusOK, usersAndPasswords)
	}
	return gin.HandlerFunc(fn)
}

func (h *FileHandler) DeleteUsersAndPasswordsFile() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		if fileHandler.DeleteFile(constant.USERS_AND_PASSWORDS_PATH) {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusInternalServerError)
		}
	}
	return gin.HandlerFunc(fn)
}

func (h *FileHandler) DeleteUsersAndPasswordsFromFile(usersAndPasswords []model.User) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userName := c.Param("userName")
		if fileHandler.DeleteLineFromFile(constant.USERS_AND_PASSWORDS_PATH, userName) {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusInternalServerError)
		}
	}
	return gin.HandlerFunc(fn)
}
