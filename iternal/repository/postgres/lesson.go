package postgres

import (
	"CrudApp/iternal/domain"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Lessons struct {
	db *sql.DB
}

func NewLessonsRepo(db *sql.DB) *Lessons { //ctor
	return &Lessons{db}
}

func (l *Lessons) Create(ctx context.Context, lesson domain.Lesson) error {
	_, err := l.db.Exec("INSERT INTO lessons (id, subject, classroom, teacher, starttime, grade) values ($1, $2, $3, $4, $5, $6)",
			lesson.ID, lesson.Subject, lesson.Classroom, lesson.Teacher, lesson.Starttime, lesson.Grade)

	return err
}

func (l *Lessons) Delete(ctx context.Context, id int64) error {
	_, err := l.db.Exec("DELETE FROM lessons where id = $1", id)

	return err
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
	values := make([]string, 0) // its slice for get all parameters user want to change
	arguments := make([]interface{}, 0) // its slice  with all arguments we need to change
	argNumber := 1 // counter of arguments 

	if input.Subject != nil {
		arguments = append(arguments, *input.Subject)
		values = append(values, fmt.Sprintf("subject=$%d", argNumber))
		argNumber++
	}
	if input.Classroom != nil {
		arguments = append(arguments, *input.Classroom)
		values = append(values, fmt.Sprintf("classroom=$%d", argNumber))
		argNumber++
	}
	if input.Teacher != nil {
		arguments = append(arguments, *input.Teacher)
		values = append(values, fmt.Sprintf("teacher=$%d", argNumber))
		argNumber++
	}
	if input.Starttime != nil {
		arguments = append(arguments, *input.Starttime)
		values = append(values, fmt.Sprintf("starttime=$%d", argNumber))
		argNumber++
	}
	if input.Grade != nil {
		arguments = append(arguments, *input.Grade)
		values = append(values, fmt.Sprintf("grade=$%d", argNumber))
		argNumber++
	}
	arguments = append(arguments, id)

	queryValues := strings.Join(values, ", ")
	query := fmt.Sprintf("UPDATE lessons SET %s WHERE id=$%d", queryValues, argNumber) // get a query like UPDATE {} SET {smth}="123" .... WHERE ID = id

	_, err := l.db.Exec(query, arguments...)
	return err
}