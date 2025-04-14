package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID string
	Name string 
	Password string
	Email string 
	CreatedAt time.Time
	UpdatedAt time.Time
}

func createHashPassword(password string) (string, error) {
bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

return string(bytes), err

}

func CheckPasswordHash (password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}


func NewUser (name, email, password string) (*User, error) {
	hashPassword, err := createHashPassword(password)

	if err!= nil {
		fmt.Println("Error for creating hash: ", err)
		return nil, err
	}

	user := User{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		Password: hashPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
	
}