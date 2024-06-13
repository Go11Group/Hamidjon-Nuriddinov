package module

import "time"

type User struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Birthday string    `json:"birthday"`
	CreateAt time.Time `json:"created_at"`
}

type Course struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Lesson struct {
	Title     string    `json:"title"`
	Content   string    `json:"conetnt"`
	CreatedAt time.Time `son:"created_at"`
}

type Enrollment struct {
}
