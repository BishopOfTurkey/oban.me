package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"roadstoendpoint/middleware"

	"github.com/alexedwards/scs/v2"
	_ "github.com/mattn/go-sqlite3" //SQLite3 driver
)

var sessionManager *scs.SessionManager

var db *sql.DB

//Router f
func Router() {
	var err error
	db, err = sql.Open("sqlite3", "db/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sessionManager = middleware.NewSessionManager(db)

	in := http.NewServeMux() //Root router

	public := http.NewServeMux()
	authed := http.NewServeMux()

	in.Handle("/", public)
	in.Handle("/map/", authed)

	public.Handle("/auth", http.HandlerFunc(Auth))
	public.Handle("/auth/callback", http.HandlerFunc(AuthCallback))
	public.Handle("/", http.HandlerFunc(Index))

	authed.HandleFunc("/", Map)

	middlewares := middleware.Middlewares{middleware.Logger(), middleware.Sessions(sessionManager)}
	out := middleware.Chain(in, middlewares)

	fmt.Println("listening on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", out))
}
