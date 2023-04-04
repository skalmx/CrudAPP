package service

import (
	"CrudApp/iternal/domain"
	"context"
)

type LessonsRepository interface {
	Create(ctx context.Context, lesson domain.Lesson) error
	Delete(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (domain.Lesson, error)
	GetAll(ctx context.Context) ([]domain.Lesson, error)
	Update(ctx context.Context, id int64, input domain.UpdateLesson) error
}

type Lessons struct {
	repository LessonsRepository
}

func NewLessons(repository LessonsRepository) * Lessons { //ctor
	return &Lessons{
		repository: repository,
	}
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
	return l.repository.GetAll(ctx)
}

func (l *Lessons) Update(ctx context.Context, id int64, input domain.UpdateLesson) error {
	// todo: realization for update
	return nil
}