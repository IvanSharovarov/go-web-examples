package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

// Todo list item
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData page data
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	// get absolute path
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.ParseFiles(wd + "/" + "layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":9000", nil)
}
