package user

import "time"

type User struct {
	ID         int       `db:"id"`
	FirstName  string    `db:"nombre"`
	LastName   string    `db:"apellido"`
	Email      string    `db:"email"`
	Password   string    `db:"contrasena" json:"-"`
	SignupDate time.Time `db:"fechaCreacion"`
}

//
