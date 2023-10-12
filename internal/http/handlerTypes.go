package http

import "crud-go-project/internal/model"

type FileHandler struct {
	fileHandler FileHandlerInterface
}

type PasswordHandler struct {
	passwordHandler PasswordHandlerInterface
}

type FileHandlerInterface interface {
	ReadFileToStruct(filename string) []model.User
}

type PasswordHandlerInterface interface {
	GetAllPasswords(users []model.User) []string
}

func NewFileHandler(fileHandler FileHandlerInterface) *FileHandler {
	return &FileHandler{
		fileHandler: fileHandler,
	}
}

func NewPasswordHandler(passwordHandler PasswordHandlerInterface) *PasswordHandler {
	return &PasswordHandler{
		passwordHandler: passwordHandler,
	}
}
