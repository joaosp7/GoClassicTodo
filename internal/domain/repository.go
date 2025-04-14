package domain

type UserRepository interface {
	Create(user *User) error
	FindById(id string) (*User , error)
}

type TodoRepository interface {
	Create() error
	FindById(id string)
}

