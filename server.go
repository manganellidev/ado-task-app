package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Version string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/ endpoint hit")

		version, err := os.ReadFile("VERSION")
		if err != nil {
			log.Fatal("Could not read VERSION file:", err)
		}

		pageData := PageData{
			Version: string(version),
		}

		t := template.Must(template.ParseFiles("template.html"))
		err = t.Execute(w, pageData)
		if err != nil {
			log.Println("Error rendering template:", err)
			http.Error(w, "Could not render template", http.StatusInternalServerError)
		}
	})

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
