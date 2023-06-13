package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "login.html"

	file, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error Parsing file ", err)
		return
	}
	err = file.ExecuteTemplate(w, fileName, nil)

	if err != nil {
		fmt.Println("error Executing Templates ", err)
		return
	}

}

var userDB = map[string]string{
	"alien": "password",
}

func lsubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "loging success..")

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, " Page Not Found ")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		login(w, r)
	case "/l-submit":
		lsubmit(w, r)
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8343", nil)

}
