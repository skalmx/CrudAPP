package rest

import (
	"CrudApp/iternal/domain"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

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
		r.Route("/{lessonsId}", func(r chi.Router) {
			// r.Delete("/{id}",)
			r.Put("/", h.updateLesson)
			// r.Get("/{id}",)
		})	
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
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) updateLesson(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("cant get id in updateLesson()", err)
		return
	}

	req, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Cant read request body in updateLesson()", err)
		return
	}

	var lessonInp *domain.UpdateLesson
	if err = json.Unmarshal(req, &lessonInp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Cant unmarshal in updateLesson()", err)
		return
	}

	err = h.lessonsService.Update(context.TODO(), id, *lessonInp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error in a service update method", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func getIdFromRequest (r *http.Request) (int64, error) {
	vars := chi.URLParam(r, "lessonsId")
	id, err := strconv.ParseInt(vars, 10, 64)
	if err != nil {
		return 0, err
	}
	if id <= 0 {
		return 0, errors.New("id can be only > 0")
	}
	return id, nil
}