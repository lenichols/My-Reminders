package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Reminder struct {
	Item string
	Done bool
}

type PageData struct {
	Title     string
	Reminders []Reminder
}

func reminders(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		Title: "Reminders List",
		Reminders: []Reminder{
			{Item: "Pay Capital One Bill", Done: true},
			{Item: "Student Loan Payment Due", Done: false},
			{Item: "Setup House Cleaning", Done: true},
		},
	}

	tmpl.Execute(w, data)

}

func main() {

	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("/reminders", reminders)
	log.Fatal(http.ListenAndServe(":9091", mux))

}
