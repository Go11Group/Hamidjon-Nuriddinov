package users

import (
	"time"
)

type User struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Birthday string    `json:"birthday"`
	CreateAt time.Time `json:"created_at"`
}
