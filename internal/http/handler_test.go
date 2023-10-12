package http

import (
	"crud-go-project/internal/constant"
	"crud-go-project/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

var usersAndPasswords = []model.User{{UserName: constant.TEST_USER1, Password: constant.TEST_PASSWORD1}, {UserName: constant.TEST_USER2, Password: constant.TEST_PASSWORD2}}

type MockFileHandler struct {
	fileHandlerMock FileHandlerInterface
}

type MockPasswordHandler struct {
	passwordHandlerMock PasswordHandlerInterface
}

func (MockFileHandler) ReadFileToStruct(filename string) []model.User {
	return usersAndPasswords
}

func (MockPasswordHandler) GetAllPasswords(users []model.User) []string {
	passwordsOfUsers := make([]string, 0)
	for _, guessed := range users {
		passwordsOfUsers = append(passwordsOfUsers, fmt.Sprintf("The password of user %v is %v", guessed.UserName, guessed.Password))
	}
	return passwordsOfUsers
}

func TestGetUsersHandler(t *testing.T) {

	mockFileHandler := MockFileHandler{}
	handler := NewFileHandler(mockFileHandler)

	r := gin.Default()
	r.GET("/users", handler.GetUsers())

	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestGetPasswordsHandler(t *testing.T) {

	mockPasswordHandler := MockPasswordHandler{}
	handler := NewPasswordHandler(mockPasswordHandler)

	r := gin.Default()
	r.GET("/passwordsOfUsers", handler.GetPasswords(usersAndPasswords))

	req, err := http.NewRequest(http.MethodGet, "/passwordsOfUsers", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
