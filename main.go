package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "./transport.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	userModel := &UserModel{DB: db}
	userModel.InitTable()
	staticDir := "./static"

	// Serve static files for any path **except** "/" and "/auth"
	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/auth", handleAuth)
	http.HandleFunc("/schooldashboard", handleDashboardSchool)
	http.HandleFunc("/parentsdashboard", handleDashboardParents)
	
	log.Println("Server is running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}