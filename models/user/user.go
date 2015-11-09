package user

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNoIDSent     = errors.New("The user ID is empty.")
	ErrUserNotExist = errors.New("The user does not exist.")
	ErrNoEmailSend  = errors.New("The email no send.")
	ErrNoUsersFound = errors.New("No users found.")
)

const (
	sqlGetUserByID      = "SELECT id, nombre, apellidos, userName, email, contrasena, fechaCreacion, admin FROM usuario WHERE id = ? LIMIT 1"
	sqlUserExistByEmail = "SELECT true FROM usuario where email = ? LIMIT 1"
	sqlLastIdUser       = "SELECT id FROM usuario WHERE email = ? Limit 1"
)

type User struct {
	ID         int       `db:"id"`
	FirstName  string    `db:"nombre"`
	LastName   string    `db:"apellidos"`
	UserName   string    `db:"userName"`
	Email      string    `db:"email"`
	Password   string    `db:"contrasena" json:"-"`
	SignupDate time.Time `db:"fechaCreacion"`
	Admin      bool      `db:"admin"`
}

func Get(db *sqlx.DB, id int) (*User, error) {
	if id <= 0 {
		return nil, ErrNoIDSent
	}

	u := &User{}
	err := db.QueryRow(sqlGetUserByID, id).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.UserName,
		&u.Email,
		&u.Password,
		&u.SignupDate,
		&u.Admin,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotExist
		}
		panic(err)
	}

	return u, nil
}

func (u User) EmailExist(db *sqlx.DB) bool {

	if u.Email == "" {
		panic(ErrNoEmailSend)
	}

	row := db.QueryRow(sqlUserExistByEmail, u.Email)

	var exist bool
	if err := row.Scan(&exist); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		panic(err)
	}

	return exist

}

func (u *User) LoadID(db *sqlx.DB) error {

	row := db.QueryRow(sqlLastIdUser, u.Email)

	var num int
	if err := row.Scan(&num); err != nil {
		if err == sql.ErrNoRows {
			return ErrNoUsersFound
		}
		panic(err)
	}
	u.ID = num
	return nil
}
