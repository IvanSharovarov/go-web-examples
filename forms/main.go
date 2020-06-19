package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type ContractDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	// get absolute path
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.ParseFiles(wd + "/forms/" + "form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContractDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":9000", nil)
}
