package routes

import (
	"errors"

	"golang.org/x/oauth2"
)

type User struct {
	UserID      string
	Firstname   string
	Lastname    string
	StravaID    float64
	StravaToken oauth2.Token
	UpdatedAt   string
}

func createAndAddUser(athlete *Athlete, token oauth2.Token) error {
	if _, found := getUser(athlete.ID); found {
		return errors.New("User already exists")
	}
	id, err := getNewUserID()
	if err != nil {
		return err
	}
	user := User{
		UserID:      id,
		Firstname:   athlete.Firstname,
		Lastname:    athlete.Lastname,
		StravaID:    athlete.ID,
		StravaToken: token,
		UpdatedAt:   athlete.UpdatedAt,
	}
	q := "INSERT INTO users (UserID, Firstname, Lastname, StravaID, UpdatedAt) VALUES (?,?,?,?,?)"
	_, err = db.Exec(q, user.UserID, user.Firstname, user.Lastname, user.StravaID, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func getUser(stravaID float64) (*User, bool) {
	row := db.QueryRow("SELECT * FROM users WHERE StravaID = ?", stravaID)

	var user User

	err := row.Scan(&user)
	if err != nil {
		return nil, false
	}

	return &user, true
}

func getNewUserID() (string, error) {
	row := db.QueryRow("SELECT word FROM dictionary LIMIT 1")
	var word string
	err := row.Scan(&word)
	if err != nil {
		return "", err
	}
	_, err = db.Exec("DELETE FROM dictionary WHERE word = ?", word)
	if err != nil {
		return "", err
	}
	return word, nil
}
