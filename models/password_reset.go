package models

import (
	"fmt"
	"time"

	"github.com/bonhokage06/lenslocked/database"
	"github.com/bonhokage06/lenslocked/helpers"
	"github.com/bonhokage06/lenslocked/rand"
	"github.com/pocketbase/dbx"
)

type PasswordReset struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	//token is only set when passwordreset is being created
	TokenHash string    `json:"token_hash"`
	ExpiresAt time.Time `json:"expires_at"`
}
type PasswordResetUser struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	ExpiresAt    time.Time
}

// create a function that will create a password reset token
func (pr *PasswordReset) Create() error {
	//create a token
	token, err := rand.PassworToken()
	if err != nil {
		return err
	}
	pr.TokenHash = token
	//set the expiration time
	pr.ExpiresAt = time.Now().Add(time.Minute * 30)
	//save the token to the database
	query := database.Db.NewQuery("INSERT INTO password_resets (user_id,token_hash,expires_at) VALUES ({:user_id},{:token_hash},{:expires_at}) ON CONFLICT (user_id) DO UPDATE SET token_hash={:token_hash},expires_at={:expires_at}")
	query.Bind(dbx.Params{
		"user_id":    pr.UserId,
		"token_hash": helpers.Encode(pr.TokenHash),
		"expires_at": pr.ExpiresAt,
	})
	_, err = query.Execute()
	if err != nil {
		return err
	}
	return nil
}

// check token if valid
func (pr *PasswordReset) Check() PasswordResetUser {
	//check if the token exists in the database
	var passwordResetUser PasswordResetUser
	query := database.Db.NewQuery("SELECT users.email as email,users.password_hash as password_hash,password_resets.expires_at as expires_at FROM password_resets JOIN users ON password_resets.user_id=users.id WHERE password_resets.token_hash = {:token_hash} LIMIT 1")
	query.Bind(dbx.Params{
		"token_hash": helpers.Encode(pr.TokenHash),
	})
	err := query.One(&passwordResetUser)
	if err != nil {
		return passwordResetUser
	}
	//check if the token has expired
	if time.Now().After(passwordResetUser.ExpiresAt) {
		return passwordResetUser
	}
	fmt.Println(passwordResetUser)
	return passwordResetUser
}

func (pr *PasswordReset) Delete() error {
	_, err := database.Db.Delete("password_resets", dbx.HashExp{"token_hash": helpers.Encode(pr.TokenHash)}).Execute()
	if err != nil {
		return err
	}
	return nil
}
