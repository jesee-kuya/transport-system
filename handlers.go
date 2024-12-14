package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("public/authentication.html")
		if err != nil {
			fmt.Println(err)
		}
		tmpl.Execute(w, nil)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get form action (login or signup) and user type (school or parent)
	action := r.FormValue("action")     // "login" or "signup"
	userType := r.FormValue("userType") // "school" or "parent"

	if action == "" || userType == "" {
		http.Error(w, "Missing form action or user type", http.StatusBadRequest)
		return
	}

	switch action {
	case "login":
		handleLogin(w, r, userType)
	case "signup":
		handleSignup(w, r, userType)
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request, userType string) {
	db, err := sql.Open("sqlite", "./transport.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()
	userModel := &UserModel{DB: db}
	var email, password string

	if userType == "school" {
		email = r.FormValue("schoolEmail")
		password = r.FormValue("schoolPassword")
	} else if userType == "parent" {
		email = r.FormValue("parentEmail")
		password = r.FormValue("parentPassword")
	} else {
		http.Error(w, "Invalid user type", http.StatusBadRequest)
		return
	}
	exist, err := userModel.CheckCredentials(email, password, userType)
	if err != nil {
		http.Error(w, "Error checking creddentials", http.StatusInternalServerError)
		return
	}
	if exist {
		if userType == "school" {
			http.Redirect(w, r, "/schooldashboard", http.StatusSeeOther)
			return
		} else {
			http.Redirect(w, r, "/parentsdashboard", http.StatusSeeOther)
			return
		}
	} else {
		http.Error(w, "Wrong username or password", http.StatusBadRequest)
		return
	}
}

func handleSignup(w http.ResponseWriter, r *http.Request, userType string) {
	db, err := sql.Open("sqlite", "./transport.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()
	userModel := &UserModel{DB: db}

	if userType == "school" {
		schoolName := r.FormValue("schoolName")
		email := r.FormValue("signupSchoolEmail")
		password := r.FormValue("signupSchoolPassword")
		confirmPassword := r.FormValue("signupSchoolConfirmPassword")

		if password != confirmPassword {
			http.Error(w, "Passwords do not match", http.StatusBadRequest)
			return
		}
		success, err := userModel.InsertSchool(schoolName, email, password)
		if err != nil {
			fmt.Println(err)
			return
		}
		if success {
				http.Redirect(w, r, "/schooldashboard", http.StatusSeeOther)
				return
		}
		

		log.Printf("Signup Attempt: SchoolName=%s, Email=%s, Password=%s", schoolName, email, password)
		// fmt.Fprintf(w, "School Signup Successful for %s", schoolName)
	} else if userType == "parent" {
		fullName := r.FormValue("parentFullName")
		email := r.FormValue("parentSignupEmail")
		school := r.FormValue("parentSchool")
		childAdmissionNumber := r.FormValue("childAdmissionNumber")
		password := r.FormValue("signupParentPassword")
		confirmPassword := r.FormValue("signupParentConfirmPassword")

		if password != confirmPassword {
			http.Error(w, "Passwords do not match", http.StatusBadRequest)
			return
		}
		success, err := userModel.InsertParent(fullName, email, school, childAdmissionNumber, password)
		if err != nil {
			http.Error(w, "error inserting entry", http.StatusInternalServerError)
		}
		if success {
			http.Redirect(w, r, "/parentsdashboard", http.StatusSeeOther)
			return
	}

		log.Printf("Signup Attempt: FullName=%s, Email=%s, School=%s, ChildAdmission=%s", fullName, email, school, childAdmissionNumber)
		// fmt.Fprintf(w, "Parent Signup Successful for %s", fullName)
	} else {
		http.Error(w, "Invalid user type", http.StatusBadRequest)
		return
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, nil)
}
func handleDashboardParents(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/parentsdashboard.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, nil)
}
func handleDashboardSchool(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/schooldashboard.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, nil)
}
