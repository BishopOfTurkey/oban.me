package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type index struct {
	Time      time.Time
	TestField string
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmap := template.FuncMap{
		"formatAsTime": formatAsTime,
	}
	data := &index{
		Time: time.Now(),
	}
	err := renderView(w, "views/index.html", &fmap, data)
	if err != nil {
		log.Fatalf("Failed to parse HTML template: %v", err)
	}
}

func formatAsTime(t time.Time) string {
	return fmt.Sprintf("%v", t)
}
