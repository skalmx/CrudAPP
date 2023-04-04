package postgres

import (
	"CrudApp/iternal/domain"
	"context"
	"database/sql"
)

type Lessons struct {
	db *sql.DB
}

func NewLessonsRepo(db *sql.DB) *Lessons { //ctor
	return &Lessons{db}
}

func (l *Lessons) Create(ctx context.Context, lesson domain.Lesson) error {
	// todo: realization for create
	return nil
}

func (l *Lessons) Delete(ctx context.Context, id int64) error {
	// todo: realization for delete
	return nil
}

func (l *Lessons) GetById(ctx context.Context, id int64) (domain.Lesson, error) {
	// todo: realization for getByID
	return domain.Lesson{}, nil
}

func (l *Lessons) GetAll(ctx context.Context) ([]domain.Lesson, error) {
	// todo: realization for getAll 
	rows, err := l.db.Query("SELECT id, subject, classroom, teacher, starttime, grade FROM lessons")
	if err != nil {
		return nil, err
	}

	lessons := make([]domain.Lesson,0)
	for rows.Next() {
		var lesson domain.Lesson
		if err := rows.Scan(&lesson.ID, &lesson.Subject, &lesson.Classroom, &lesson.Teacher, &lesson.Starttime, &lesson.Grade); err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, rows.Err()
}

func (l *Lessons) Update(ctx context.Context, id int64, input domain.UpdateLesson) error {
	// todo: realization for update
	return nil
}