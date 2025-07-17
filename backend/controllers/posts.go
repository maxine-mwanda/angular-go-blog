package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"angular-go-blog/repositories"

	"github.com/gorilla/mux"
)

type PostController struct {
	repo *repositories.PostRepository
}

func NewPostController(db *sql.DB) *PostController {
	return &PostController{
		repo: repositories.NewPostRepository(db),
	}
}

func (c *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := c.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}

func (c *PostController) GetPostBySlug(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	post, err := c.repo.GetBySlug(slug)

	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, post)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
