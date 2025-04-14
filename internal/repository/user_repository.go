package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joaosp7/GoClassicTodo/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(input *domain.User) error {
	// First check if the table exists
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(36) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL,
			updatedAt TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		log.Printf("Error creating table: %v", err)
		return fmt.Errorf("error creating table: %v", err)
	}

	stmt, err := r.db.Prepare(`
		INSERT INTO users (id, name, email, password, createdAt, updatedAt) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	var user *domain.User
	user, err = domain.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		log.Printf("Error creating user domain object: %v", err)
		return fmt.Errorf("error creating user domain object: %v", err)
	}

	log.Printf("Attempting to insert user with values: ID=%s, Name=%s, Email=%s", user.ID, user.Name, user.Email)

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		log.Printf("Error executing insert: %v", err)
		return fmt.Errorf("error executing insert: %v", err)
	}

	log.Printf("Successfully created user with ID: %s", user.ID)
	return nil
}

func (r *UserRepository) FindById(id string) (*domain.User, error) {
	var user domain.User
	err := r.db.QueryRow(`
	SELECT id, name,password,  email, createdAt, updatedAt
	FROM users
	WHERE id = $1`, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
