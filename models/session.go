package models

import (
	"github.com/bonhokage06/lenslocked/database"
	"github.com/bonhokage06/lenslocked/helpers"
	"github.com/bonhokage06/lenslocked/rand"
	"github.com/pocketbase/dbx"
)

type Session struct {
	Email         string `db:"email"`
	UserId        int    `db:"user_id"`
	RememberToken string `db:"remember_token"`
}
type UserSession struct {
	Email  string `db:"email"`
	UserId int    `db:"user_id"`
}

// create a session
func (s Session) Create() (Session, error) {
	//delete old sessions
	_, err := database.Db.Delete("sessions", dbx.HashExp{"user_id": s.UserId}).Execute()
	if err != nil {
		return Session{}, err
	}
	rememberToken, err := rand.SessionToken()
	if err != nil {
		return Session{}, err
	}
	_, err = database.Db.Insert("sessions", dbx.Params{
		"user_id":        s.UserId,
		"remember_token": helpers.Encode(rememberToken),
	}).Execute()
	if err != nil {
		return Session{}, err
	}
	return Session{
		RememberToken: rememberToken,
	}, nil
}

// delete a session
func (s Session) Delete() error {
	_, err := database.Db.Delete("sessions", dbx.HashExp{"remember_token": s.RememberToken}).Execute()
	if err != nil {
		return err
	}
	return nil
}

// check if a session is valid
func (s Session) Check() (UserSession, error) {
	var userSession UserSession
	query := database.Db.NewQuery("SELECT email,user_id FROM users JOIN sessions ON users.id=sessions.user_id WHERE remember_token = {:remember_token}")
	query.Bind(dbx.Params{"remember_token": helpers.Encode(s.RememberToken)})
	err := query.One(&userSession)
	if err != nil {
		return userSession, err
	}
	return userSession, nil
}
