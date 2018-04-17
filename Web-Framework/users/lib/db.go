package lib

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var settings = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "root",
	Password: "",
	Database: "membros",
}

// Sess is a variable what do connection on database
var Sess db.Database

func init() {
	var err error

	Sess, err = mysql.Open(settings)

	if err != nil {
		log.Fatal(err.Error())
	}
}
