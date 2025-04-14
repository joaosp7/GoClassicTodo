package domain

import "time"

type Todo struct {
	ID string
	Title string
	Description string 
	DueDate time.Time
	CreateAt time.Time
	UpdatedAt time.Time
	UserID string 

}