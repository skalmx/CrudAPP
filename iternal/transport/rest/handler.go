package rest

import (
	"CrudApp/iternal/domain"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type Lessons interface {
	Create(ctx context.Context, lesson domain.Lesson) error
	Delete(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (domain.Lesson, error)
	GetAll(ctx context.Context) ([]domain.Lesson, error)
	Update(ctx context.Context, id int64, input domain.UpdateLesson) error
}

type Handler struct {
	lessonsService Lessons
}

func NewHandler (lessons Lessons) *Handler { //ctor
	return &Handler{
		lessonsService: lessons,
	}
}

func (h *Handler) Init() *chi.Mux { 
	r := chi.NewRouter()
	r.Route("/lessons", func(r chi.Router){
		// r.Post("",)
		r.Get("/",h.getAllLessons)
		// r.Delete("/{id}",)
		// r.Put("/{id}",)
		// r.Get("/{id}",)
	})

	return r
} 

func (h *Handler) getAllLessons(w http.ResponseWriter, r *http.Request) {
	lessons, err := h.lessonsService.GetAll(context.TODO())
	if err != nil {
		log.Println("GetAllLessons func error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(lessons)
	if err != nil {
		log.Println("GetAllLessons func error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
