package activity

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNoIDSent = errors.New("El identificador del usuario es Invalido.")
)

const (
	sqlInsertActivity            = "INSERT INTO tareas (titulo, descripcion, fechaTermino, miembros_id, estado, adjuntos_id, pioridad) VALUES (:titulo, :descripcion, :fechaTermino, :miembros_id, :estado, :adjuntos_id, :pioridad );"
	sqlInsertComments            = "INSERT INTO comentarios (fechaCreacion, comentarioscol, tareas_id) VALUES(:fechaCreacion, :comentarioscol, :tareas_id);"
	sqlGetActivityByID           = "SELECT id, titulo, descripcion, fechaTermino, miembros_id, estado, adjuntos_id, pioridad FROM tareas  WHERE miembros_id = ?;"
	sqlGetActivityImpedidasByID  = "SELECT id, titulo, descripcion, fechaTermino, miembros_id, estado, adjuntos_id, pioridad FROM tareas  WHERE miembros_id = ? && estado = 'impedida' ;"
	sqlGetActivityPendientesByID = "SELECT id, titulo, descripcion, fechaTermino, miembros_id, estado, adjuntos_id, pioridad FROM tareas  WHERE miembros_id = ? && estado = 'Pendientes' ;"
	sqlGetActivityEnProcesoByID  = "SELECT id, titulo, descripcion, fechaTermino, miembros_id, estado, adjuntos_id, pioridad FROM tareas  WHERE miembros_id = ? && estado = 'EnProceso' ;"
	sqlGetActivityTerminadosByID = "SELECT id, titulo, descripcion, fechaTermino, miembros_id, estado, adjuntos_id, pioridad FROM tareas  WHERE miembros_id = ? && estado = 'Terminados' ;"
)

type Activity struct {
	ID          int    `db:"id", json:"id"`
	Title       string `db:"titulo", json:"titulo"`
	Description string `db:"descripcion", json:"descripcion"`
	ExpireDate  string `db:"fechaTermino", json:"fecha Expiracion"`
	ExpireDates time.Time
	User        int        `db:"miembros_id", json:"Recurso"`
	Estate      string     `db:"estado", json:"Estado"`
	Adjuntos    int        `db:"adjuntos_id", json:"adjuntos"`
	Pioridad    string     `db:"pioridad", json:"pioridad"`
	Comments    []*Comment `db:"comentarios", json:"Comentarios"`
}

type Comment struct {
	IDTarea       int       `db:"tareas_id", json:"idTarea"`
	FechaCreacion time.Time `db:"fechaCreacion", json:"FechaCreacion"`
	Comentario    string    `db:"comentarioscol", json:"Comentario"`
}

func (a Activity) Create(db *sqlx.DB) {

	fmt.Println("\n\n\n", a)

	_, err := db.NamedExec(sqlInsertActivity, &a)
	if err != nil {
		panic(err)
	}
}

func (c Comment) NewComment(db *sqlx.DB) {

	_, err := db.NamedExec(sqlInsertComments, &c)
	if err != nil {
		panic(err)
	}
}

func Get(db *sqlx.DB, id int) ([]*Activity, error) {
	var (
		Activitys []*Activity
	)

	if id <= 0 {
		return nil, ErrNoIDSent
	}

	rows, err := db.Query(sqlGetActivityByID, id)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

		a := &Activity{}
		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Description,
			&a.ExpireDates,
			&a.User,
			&a.Estate,
			&a.Adjuntos,
			&a.Pioridad,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return Activitys, nil
			}
			panic(err)
		}
		Activitys = append(Activitys, a)
	}

	fmt.Println(Activitys[0])
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return Activitys, nil
}

func GetImpedidas(db *sqlx.DB, id int) ([]*Activity, error) {
	var (
		Activitys []*Activity
	)

	if id <= 0 {
		//execute sql sin id
		return nil, ErrNoIDSent
	}

	rows, err := db.Query(sqlGetActivityImpedidasByID, id)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

		a := &Activity{}
		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Description,
			&a.ExpireDates,
			&a.User,
			&a.Estate,
			&a.Adjuntos,
			&a.Pioridad,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return Activitys, nil
			}
			panic(err)
		}
		Activitys = append(Activitys, a)
	}

	fmt.Println(Activitys[0])
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return Activitys, nil
}

func GetPendintes(db *sqlx.DB, id int) ([]*Activity, error) {
	var (
		Activitys []*Activity
	)

	if id <= 0 {
		//execute sql sin id
		return nil, ErrNoIDSent
	}

	rows, err := db.Query(sqlGetActivityPendientesByID, id)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

		a := &Activity{}
		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Description,
			&a.ExpireDates,
			&a.User,
			&a.Estate,
			&a.Adjuntos,
			&a.Pioridad,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return Activitys, nil
			}
			panic(err)
		}
		Activitys = append(Activitys, a)
	}

	fmt.Println(Activitys[0])
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return Activitys, nil
}

func GetEnProceso(db *sqlx.DB, id int) ([]*Activity, error) {
	var (
		Activitys []*Activity
	)

	if id <= 0 {
		//execute sql sin id
		return nil, ErrNoIDSent
	}

	rows, err := db.Query(sqlGetActivityEnProcesoByID, id)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

		a := &Activity{}
		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Description,
			&a.ExpireDates,
			&a.User,
			&a.Estate,
			&a.Adjuntos,
			&a.Pioridad,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return Activitys, nil
			}
			panic(err)
		}
		Activitys = append(Activitys, a)
	}

	fmt.Println(Activitys[0])
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return Activitys, nil
}

func GetTerminados(db *sqlx.DB, id int) ([]*Activity, error) {
	var (
		Activitys []*Activity
	)

	if id <= 0 {
		//execute sql sin id
		return nil, ErrNoIDSent
	}

	rows, err := db.Query(sqlGetActivityTerminadosByID, id)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

		a := &Activity{}
		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Description,
			&a.ExpireDates,
			&a.User,
			&a.Estate,
			&a.Adjuntos,
			&a.Pioridad,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return Activitys, nil
			}
			panic(err)
		}
		Activitys = append(Activitys, a)
	}

	fmt.Println(Activitys[0])
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return Activitys, nil
}
