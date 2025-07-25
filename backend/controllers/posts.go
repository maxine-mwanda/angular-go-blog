package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"angular-go-blog/models"
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

func (c *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.BlogPost
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if post.Title == "" || post.Content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}
	newPost := c.repo.Create(&post)
	if newPost == nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, newPost)
}

// UpdatePost
func (c *PostController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]

	var updates models.BlogPost
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updates.Slug = slug
	updatedPost := c.repo.Update(&updates)
	if updatedPost == nil {
		http.NotFound(w, r)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedPost)
}

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
