package services

import (
	"fmt"

	"github.com/joaosp7/GoClassicTodo/internal/domain"
	"github.com/joaosp7/GoClassicTodo/internal/dto"
)

type UserService struct {
	repository domain.UserRepository
}

func NewUserService (repository domain.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateAccount(input dto.UserRequestDto) (*dto.UserResponseDto) {	
	user,err := dto.ToUser(input)
	fmt.Println("service go : ",user)
	if (err!=nil) {
		return nil
	}

	err = s.repository.Create(user) 
	if err!=nil {
		return nil
	}

	userResponse := dto.ToResponseOutput(user)

	return &userResponse

}

