package model
import (
	"time"
)

type Todo struct {
	ID         string
	Title			 string
	Content			string
	IsCompleted	bool
	User			*User
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
