package routes

import (
	"log"
	"net/http"
	"time"
)

type usersMap struct {
	Time      time.Time
	TestField string
}

//Map d
func Map(w http.ResponseWriter, r *http.Request) {
	sessionManager.Put(r.Context(), "test", "test1233456")

	test := sessionManager.GetString(r.Context(), "test")

	data := &usersMap{
		TestField: test,
	}
	err := renderView(w, "views/map.html", nil, data)
	if err != nil {
		log.Fatalf("Failed to parse HTML template: %v", err)
	}
}
