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
	sqlInsertUser = "INSERT INTO usuario ( nombre, apellidos, userName, email, contrasena, fechaCreacion, admin) VALUES (:nombre, :apellidos, :userName, :email, :contrasena, :fechaCreacion, :admin)"

	sqlGetUserByID       = "SELECT id, nombre, apellidos, userName, email, contrasena, fechaCreacion, admin FROM usuario WHERE id = ? LIMIT 1"
	sqlUpdateUserDetails = "UPDATE usuario SET nombre = :nombre, apellidos = :apellidos, userName = :userName, email = :email, contrasena = :contrasena, admin = :admin, fechaCreacion = :fechaCreacion  WHERE ID = :id LIMIT 1"

	sqlGetAllUserRange = "SELECT id, nombre, apellidos, userName, email, contrasena, fechaCreacion, admin FROM usuario LIMIT ?,?"
	sqlGetAllUser      = "SELECT id, nombre, apellidos, userName, email, contrasena, fechaCreacion, admin FROM usuario"
	sqlUserCount       = "SELECT count(id) FROM usuario"
	sqlDeleteUserByID  = "DELETE FROM usuario WHERE ID=? LIMIT 1"

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

func GetAll(db *sqlx.DB) ([]*User, error) {
	var (
		users []*User
	)

	rows, err := db.Query(sqlGetAllUser)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		u := &User{}
		err := rows.Scan(
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
				return users, nil
			}
			panic(err)
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return users, nil
}

func Range(db *sqlx.DB, currentPage, perPage int) ([]*User, error) {
	var start, limit int

	start = (currentPage - 1) * perPage
	limit = perPage

	var (
		users []*User
	)

	rows, err := db.Query(sqlGetAllUserRange, start, limit)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		u := &User{}
		err := rows.Scan(
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
				return users, nil
			}
			panic(err)
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return users, nil
}

func CountAll(db *sqlx.DB) int {
	row := db.QueryRow(sqlUserCount)

	var total int
	if err := row.Scan(&total); err != nil {
		if err == sql.ErrNoRows {
			return 0
		}
		panic(err)
	}

	return total
}

func (u User) Create(db *sqlx.DB) {

	_, err := db.NamedExec(sqlInsertUser, &u)
	if err != nil {
		panic(err)
	}
}

func (u User) Save(db *sqlx.DB) {

	_, err := db.NamedExec(sqlUpdateUserDetails, &u)
	if err != nil {
		panic(err)
	}

}

func Delete(db *sqlx.DB, id int) error {
	if id <= 0 {
		return ErrNoIDSent
	}

	stmt, err := db.Prepare(sqlDeleteUserByID)
	result, err := stmt.Exec(id)

	if err != nil {
		panic(err)
	}

	n, _ := result.RowsAffected()

	if n != 1 {
		return ErrUserNotExist
	}

	return nil
}
