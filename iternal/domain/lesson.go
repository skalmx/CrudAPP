package domain

import (
	"errors"
	"time"
)

type Lesson struct { // dto
	ID          int64     `json:"id"`
	Subject     string    `json:"subject"`
	Classroom   string    `json:"classroom"`
	Teacher     string    `json:"teacher"`
	Starttime   time.Time `json:"start_time"`
	Grade 		string 	  `json:"grade"`
}
type UpdateLesson struct { // dto
	Subject     *string    `json:"subject"`
	Classroom   *string    `json:"classroom"`
	Teacher     *string    `json:"teacher"`
	Starttime   *time.Time `json:"start_time"`
	Grade 		*string 	`json:"grade"`
}

var (
	ErrLessonNotExist = errors.New("there is no such lesson")
)
