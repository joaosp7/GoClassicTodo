package dto

import (
	"testing"
	"time"

	"github.com/joaosp7/GoClassicTodo/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserDto(t *testing.T){
	tests := []struct {
		userRequestdto UserRequestDto
		wantErr bool
		errMsg string
	}{
		{
			userRequestdto: UserRequestDto{Name: "john doe", Email: "johndoe@gmail.com", Password: "strongPassword"},
			wantErr: false,
		},
		{
			userRequestdto: UserRequestDto{Name: "john doe", Email: "johndoe@gmail.com", Password: "strong"},
			wantErr: false,
		},
		
	}

for _,tt := range tests {
	t.Run("It Should be able to transform Request Object", func (t *testing.T){
		user, err := ToUser(tt.userRequestdto)

		if (tt.wantErr) {
			assert.Error(t, err, "Error done")
			assert.Nil(t, user, "The return for the user should be empty")
			return
		}

		assert.NoError(t, err, "Assert there is no error in the flux")
		assert.NotNil(t, user, "Assert user is not nil")
		assert.NotEmpty(t, user.ID, "Assert there is a ID for the user")


	} )
}
}

func TestToResponseOutput(t *testing.T) {
	tests := []struct {
		user domain.User
		wantErr bool
		msgError string

	}{
		{
			user: domain.User{ID: "123", Name: "john doe", Password: "strong", Email: "john@doe.com", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			wantErr: false,
		},
	}

	for _,tt := range tests {
		t.Run("It Should be able to return response output", func(t *testing.T){
			output := ToResponseOutput(&tt.user)

			assert.NotEmpty(t, output)
			assert.NotNil(t, output)
			assert.Equal(t, tt.user.ID, output.ID)
			assert.Equal(t, tt.user.Name, output.Name)
			assert.Equal(t, tt.user.Email, output.Email)
		})
	}

}