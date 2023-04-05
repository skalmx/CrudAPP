package rest

import (
	"CrudApp/iternal/domain"
	"context"
	"encoding/json"
	"io"
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
		r.Post("/", h.createLesson)
		r.Get("/", h.getAllLessons)
		// r.Delete("/{id}",)
		// r.Put("/{id}",)
		// r.Get("/{id}",)
	})

	return r
} 

func (h *Handler) getAllLessons(w http.ResponseWriter, r *http.Request) {
	lessons, err := h.lessonsService.GetAll(context.TODO())
	if err != nil {
		log.Println("GetAllLessons func error: error in services", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(lessons)
	if err != nil {
		log.Println("GetAllLessons func error:marshaling", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (h *Handler) createLesson(w http.ResponseWriter, r *http.Request) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Cant read request body in createLesson()", err)
		return
	}

	var lesson domain.Lesson
	if err = json.Unmarshal(req, &lesson); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Cant unmarshal in createLesson()", err)
		return
	}

	if err = h.lessonsService.Create(context.TODO(), lesson); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error in service or repo method in createLesson()", err)
		return
	}
}

func (h *Handler) updateLesson(w http.ResponseWriter, r *http.Request) {
	
}