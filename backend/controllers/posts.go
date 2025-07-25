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

/*func (c *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post string
	// Decode the JSON request body into the post struct
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.repo.Create(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) //status 201
	json.NewEncoder(w).Encode(post)
}*/

/*func (r *PostController) Update (w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	post := r.repo.Update(slug)

	// Decode the JSON request body into the post struct
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	post.Slug = slug // Ensure the slug is set from the URL parameter

	if err := c.repo.Update(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) //status 200
	json.NewEncoder(w).Encode(post)
}*/

func (c *PostController) DeletePost(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	err := c.repo.Delete(slug)

	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent) //status 204
}
