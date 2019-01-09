package auth

import (
	proto "../../../proto/auth"
	"../_env"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"upper.io/db.v3"
)

// Interface zur DB
var users db.Collection
var paginationDefault uint

func InitEnvironment(env *environment.Env) {
	users = env.DB.Collection("user")
	paginationDefault = 23
}

// Passwort Username mit gespeicherten Werten vergleichen
func login(username string, password string) (proto.User, error) {
	var user proto.User
	res := users.Find(db.Cond{"username": username})
	err := res.One(&user)
	if err != nil {
		return proto.User{}, err
	}
	if checkPasswordHash(password, user.Password) {
		// falsches kennwort
		return proto.User{}, errors.New("incorrect password")
	}
	return user, nil
}

// Passwort hash erzeugen
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

// Passwort hash prüfen
// true => alles io, false => stimmt nicht
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}
