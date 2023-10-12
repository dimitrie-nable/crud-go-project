package passwordHandler

import (
	"crud-go-project/internal/constant"
	"crud-go-project/internal/model"
	"fmt"
	"sync"
	"time"
)

func TryPassword(user model.User, resultChannel chan model.User, wg *sync.WaitGroup) {
	for i := constant.MIN_PASSWORD_VALUE; i < constant.MAX_PASSWORD_VALUE; i++ {
		if i == user.Password {
			aux := model.User{
				UserName: user.UserName,
				Password: i,
			}
			resultChannel <- aux
			wg.Done()
			return
		}
	}
	fmt.Printf("Password of user %s could not be cracked", user.UserName)
	wg.Done()
	return

}

func (DefaultPasswordHandler) GetAllPasswords(users []model.User) []string {
	start := time.Now()
	resultChannel := make(chan model.User)
	wg := sync.WaitGroup{}
	for _, user := range users {
		wg.Add(1)
		go TryPassword(user, resultChannel, &wg)
	}
	go func() {
		wg.Wait()
		close(resultChannel)
	}()
	passwordsOfUsers := make([]string, 0)
	for guessed := range resultChannel {
		passwordsOfUsers = append(passwordsOfUsers, fmt.Sprintf("The password of user %v is %v", guessed.UserName, guessed.Password))
	}
	elapsed := time.Since(start)
	fmt.Printf("All passwords have been guessed in %s\n", elapsed)

	return passwordsOfUsers

}
