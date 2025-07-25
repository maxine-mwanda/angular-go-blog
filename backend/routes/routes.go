package routes

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"

	"angular-go-blog/controllers" // Replace with your actual module path

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Initialize controllers
	postCtrl := controllers.NewPostController(db)

	// API routes
	router.HandleFunc("/api/posts", postCtrl.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts/{slug}", postCtrl.GetPostBySlug).Methods("GET")
	//router.HandleFunc("/api/posts", postCtrl.CreatePost).Methods("POST")
	//router.HandleFunc("/api/posts", postCtrl.Update).Methods("POST")
	router.HandleFunc("/api/posts/{slug}", postCtrl.DeletePost).Methods("DELETE")

	// 2. Static file server configuration
	frontendPath := "/home/maxine/code/go/angular-go-blog/frontend"
	fs := http.FileServer(http.Dir(frontendPath))

	// 3. Explicit routes for static assets
	router.PathPrefix("/app/").Handler(fs)
	router.PathPrefix("/assets/").Handler(fs)
	router.Handle("/favicon.ico", fs)

	// 4. Special case for index.html
	router.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontendPath, "index.html"))
	})

	// 5. All other routes
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only serve index.html if the file doesn't exist
		if _, err := os.Stat(filepath.Join(frontendPath, r.URL.Path)); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(frontendPath, "index.html"))
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	return router
}
