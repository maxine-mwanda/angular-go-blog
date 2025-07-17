package repositories

import (
	"database/sql"
	"log"

	"angular-go-blog/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	repo := &PostRepository{db: db}
	if err := repo.initialize(); err != nil {
		log.Fatal("Failed to initialize repository:", err)
	}
	return repo
}

func (r *PostRepository) initialize() error {
	if err := r.createTable(); err != nil {
		return err
	}
	return r.seedData()
}

func (r *PostRepository) createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		slug TEXT NOT NULL UNIQUE,
		excerpt TEXT,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := r.db.Exec(query)
	return err
}

func (r *PostRepository) seedData() error {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	if err != nil || count > 0 {
		return err
	}

	posts := []models.BlogPost{
		{
			Title:   "Getting Started with AngularJS",
			Slug:    "getting-started-with-angularjs",
			Excerpt: "Learn the basics of AngularJS and how to build your first application.",
			Content: "AngularJS is a structural framework for dynamic web apps...",
		},
		{
			Title:   "Building REST APIs with Golang",
			Slug:    "building-rest-apis-with-golang",
			Excerpt: "A comprehensive guide to creating RESTful APIs using Golang.",
			Content: "Golang, also known as Go, is a statically typed, compiled language...",
		},
	}

	for _, post := range posts {
		_, err := r.db.Exec(
			"INSERT INTO posts (title, slug, excerpt, content) VALUES (?, ?, ?, ?)",
			post.Title, post.Slug, post.Excerpt, post.Content,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostRepository) GetAll() ([]models.BlogPost, error) {
	rows, err := r.db.Query(`
		SELECT id, title, slug, excerpt, content, created_at 
		FROM posts 
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.BlogPost
	for rows.Next() {
		var post models.BlogPost
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Slug,
			&post.Excerpt,
			&post.Content,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) GetBySlug(slug string) (*models.BlogPost, error) {
	var post models.BlogPost
	err := r.db.QueryRow(`
		SELECT id, title, slug, excerpt, content, created_at 
		FROM posts 
		WHERE slug = ?
	`, slug).Scan(
		&post.ID,
		&post.Title,
		&post.Slug,
		&post.Excerpt,
		&post.Content,
		&post.CreatedAt,
	)
	return &post, err
}
