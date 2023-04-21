package models

import (
	"github.com/bonhokage06/lenslocked/database"
	"github.com/bonhokage06/lenslocked/helpers"
	"github.com/pocketbase/dbx"
)

type User struct {
	Id    int    `db:"id"`
	Email string `db:"email"`
	Hash  string `db:"password_hash"`
}

func (u User) Get() ([]User, error) {
	var users []User
	err := database.Db.Select("*").From("users").All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (u User) GetUserByEmail() (User, error) {
	var user User
	err := database.Db.Select("*").From("users").Where(dbx.HashExp{"email": u.Email}).One(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// write a function that create a user
func (u User) Create() error {
	hash, err := helpers.HashAndSalt(u.Hash)
	if err != nil {
		return err
	}

	_, err = database.Db.Insert("users", dbx.Params{
		"email":         u.Email,
		"password_hash": hash,
	}).Execute()
	if err != nil {
		return err
	}
	return nil
}

// write a function that signin user
func (u User) Authenticate() (bool, User) {
	var user User
	err := database.Db.Select("*").From("users").Where(dbx.HashExp{"email": u.Email}).One(&user)
	if err != nil {
		return false, User{}
	}
	isValid := helpers.ComparePasswords(user.Hash, u.Hash)
	return isValid, user
}
