package models

import "github.com/marcelors/curso/usuarios/lib"

// User is representation the table "Users"
type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

// UserModel receive the table of database
var UserModel = lib.Sess.Collection("users")
