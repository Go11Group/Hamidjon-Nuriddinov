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
	EnrollmentId   string    `json:"enrollment_id"`
	UserId         string    `json:"user_id"`
	CourseId       string    `json:"course_id"`
	EnrollmentDate time.Time `json:"enrollment_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
