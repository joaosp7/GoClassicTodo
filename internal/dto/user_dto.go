package dto

import "github.com/joaosp7/GoClassicTodo/internal/domain"

type UserRequestDto struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserResponseDto struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}




func ToUser(input UserRequestDto) (*domain.User, error) {
	user,err := domain.NewUser(input.Name, input.Email, input.Password)
	if err!= nil {
		return nil, err
	}
	return user, nil
}

func ToResponseOutput (user *domain.User) UserResponseDto {
	return UserResponseDto{ID: user.ID, Name: user.Name, Email: user.Email}
}