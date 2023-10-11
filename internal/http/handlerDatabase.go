package http

import (
	"crud-go-project/internal/model"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsersFromDatabase(dbInstance *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var users []model.User

		rows, err := dbInstance.Query("SELECT * FROM user")

		if err != nil {
			fmt.Println(err.Error())
		}
		for rows.Next() {
			var user model.User
			if err := rows.Scan(&user.UserName, &user.Password); err != nil {
				return
			}
			users = append(users, user)
		}
		c.IndentedJSON(http.StatusOK, users)
	}
	return gin.HandlerFunc(fn)
}

func AddUserToDatabase(dbInstance *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var newUser model.User

		if err := c.BindJSON(&newUser); err != nil {
			fmt.Println(err.Error())
			return
		}
		_, err := dbInstance.Exec("INSERT INTO user (username, password) VALUES (?, ?)", newUser.UserName,
			newUser.Password)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.IndentedJSON(http.StatusCreated, newUser)
	}
	return gin.HandlerFunc(fn)
}

func DeleteUsersAndPasswordsFromDatabase(dbInstance *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userName := c.Param("userName")
		_, err := dbInstance.Exec("DELETE FROM user WHERE username = ?", userName)
		if err == nil {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusInternalServerError)
		}
	}
	return gin.HandlerFunc(fn)
}

func UpdatePasswordOfUserInDatabase(dbInstance *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userName := c.Param("userName")
		var existingUser model.User

		if err := c.BindJSON(&existingUser); err != nil {
			fmt.Println(err.Error())
			return
		}
		_, err := dbInstance.Exec("UPDATE user SET password = ? WHERE username = ?", existingUser.Password, userName)
		if err == nil {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusInternalServerError)
		}
	}
	return gin.HandlerFunc(fn)
}
