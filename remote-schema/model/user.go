package model
import (
	"time"
)

type User struct {
	ID         string
	name			 string
	last_seen			time.Time
	created_at		time.Time
}