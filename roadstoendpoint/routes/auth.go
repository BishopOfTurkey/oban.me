package routes

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

type authView struct {
	AuthLink string
}

//Auth s
func Auth(w http.ResponseWriter, r *http.Request) {
	conf := stravaConf()
	conf.RedirectURL = fmt.Sprintf("http://%v/auth/callback", r.Host)

	t := template.Must(template.New("auth.html").ParseFiles("views/auth.html"))

	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	stateCheckStr := base64.RawURLEncoding.EncodeToString(b)
	sessionManager.Put(r.Context(), "oauth_state", stateCheckStr)

	data := authView{
		AuthLink: conf.AuthCodeURL(stateCheckStr, oauth2.AccessTypeOnline),
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Fatalf("Failed to parse auth page: %v", err)
	}
}

type Athlete struct {
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	ID        float64 `json:"id"`
	UpdatedAt string  `json:"updated_at"`
}

//AuthCallback s
func AuthCallback(w http.ResponseWriter, r *http.Request) {
	conf := stravaConf()

	if r.FormValue("state") != sessionManager.GetString(r.Context(), "oauth_state") {
		w.Write([]byte("Oauth session state failed to match"))
		w.WriteHeader(400)
		return
	}

	code := r.URL.Query().Get("code")

	tok, err := conf.Exchange(context.TODO(), code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(context.TODO(), tok)
	resp, err := client.Get("https://www.strava.com/api/v3/athlete")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	var athlete Athlete
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&athlete); err != io.EOF && err != nil {
		log.Fatalf("Error decoding: %v", err)
	}

	err = createAndAddUser(&athlete, *tok)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(400)
		return
	}

	w.Header().Add("content-type", "application/json")

	out, _ := json.Marshal(athlete)

	io.Copy(w, bytes.NewReader(out))
}

type stravaCreds struct {
	ClientID     string
	ClientSecret string
}

func stravaConf() oauth2.Config {
	file, err := os.Open("./strava_creds.json")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	dec := json.NewDecoder(file)
	var creds stravaCreds
	if err := dec.Decode(&creds); err != nil {
		log.Fatalf("Error decoding strava config file: %v", err)
	}

	return oauth2.Config{
		ClientID:     creds.ClientID,
		ClientSecret: creds.ClientSecret,
		Scopes:       []string{"activity:read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.strava.com/oauth/authorize",
			TokenURL: "https://www.strava.com/oauth/token",
		},
	}
}
