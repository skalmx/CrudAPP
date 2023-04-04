package domain

import (
	"errors"
	"time"
)

type Lesson struct {
	ID          int64     `json:"id"`
	Subject     string    `json:"subject"`
	Classroom   string    `json:"classroom"`
	Teacher     string    `json:"teacher"`
	LessonStart time.Time `json:"start_time"`
	Grade 		string 	  `json:"grade"`
}
type UpdateLesson struct {
	Subject     string    `json:"subject"`
	Classroom   string    `json:"classroom"`
	Teacher     string    `json:"teacher"`
	LessonStart time.Time `json:"start_time"`
	Grade 		string 	  `json:"grade"`
}

var (
	ErrLessonNotExist = errors.New("there is no such lesson")
)
