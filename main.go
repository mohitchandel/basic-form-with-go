package main

import (
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name     string
	Email    string
	Password string
}

func handelIndex(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		err = temp.Execute(w, "index")
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
		}
		return
	}

	userInfo := UserInfo{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	err = temp.Execute(w, struct {
		Success bool
		User    UserInfo
	}{true, userInfo})

	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func main() {

	// Handling Func
	http.HandleFunc("/", handelIndex)

	// Handling Assets
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.ListenAndServe(":8080", nil)
}
