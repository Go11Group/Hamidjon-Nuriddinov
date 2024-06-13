package lessons

import "time"

type Lesson struct {
	Title     string    `json:"title"`
	Content   string    `json:"conetnt"`
	CreatedAt time.Time `son:"created_at"`
}
