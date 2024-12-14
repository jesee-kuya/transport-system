package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "modernc.org/sqlite"
)

func port() int16 {
	var port int16 = 9000

	portStr, found := os.LookupEnv("PORT")
	if !found {
		return port
	}
	iport, err := strconv.Atoi(portStr)
	if err != nil {
		return port
	}
	return int16(iport)
}

func main() {
	port := port()
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
	http.HandleFunc("/boarding", handleDashboardParents)

	p := fmt.Sprintf(":%d", port)
	fmt.Println("Server started at ", p)
	log.Fatal(http.ListenAndServe(p, nil))
}
