package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"angular-go-blog/config"
	"angular-go-blog/db"
	"angular-go-blog/routes"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer database.Close()

	router := routes.SetupRoutes(database)

	// Handle AngularJS routes
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../frontend/index.html")
	})
	log.Printf("Serving static files from %s", "../../frontend")
	// In main.go
	absPath := "/home/maxine/code/go/angular-go-blog/frontend"
	if _, err := os.Stat(filepath.Join(absPath, "index.html")); err == nil {
		log.Println("index.html FOUND at correct location")
	} else {
		log.Fatal("index.html still missing:", err)
	}

	log.Printf("Server running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
