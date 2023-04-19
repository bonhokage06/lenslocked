package models

import (
	"github.com/bonhokage06/lenslocked/database"
	"github.com/bonhokage06/lenslocked/helpers"
	"github.com/bonhokage06/lenslocked/rand"
	"github.com/pocketbase/dbx"
)

type Session struct {
	UserId        int    `db:"user_id"`
	RememberToken string `db:"remember_token"`
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
func (s Session) Check() (bool, error) {
	var session Session
	err := database.Db.Select("*").From("sessions").Where(dbx.HashExp{"remember_token": helpers.Encode(s.RememberToken)}).One(&session)
	if err != nil {
		return false, err
	}
	return true, nil
}
