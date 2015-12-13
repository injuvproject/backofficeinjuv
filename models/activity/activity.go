package activity

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNoIDSent = errors.New("El identificador del usuario es Invalido.")
)

const (
	sqlInsertActivity  = "INSERT INTO tareas (titulo, descripcion, fechaTermino, miembros_id, estado, adjuntos_id, pioridad) VALUES (:titulo, :descripcion, :fechaTermino, :miembros_id, :estado, :adjuntos_id, :pioridad );"
	sqlInsertComments  = "INSERT INTO comentarios (fechaCreacion, comentarioscol, tareas_id) VALUES(:fechaCreacion, :comentarioscol, :tareas_id);"
	sqlGetActivityByID = "SELECT id, titulo, descripcion, fechaTermino, miembros_id, estado, adjuntos_id, pioridad, fechaCreacion, comentarioscol FROM tareas T INNER JOIN comentarios C ON T.id = C.areas_id WHERE T.miembros_id = ?;"
)

type Activity struct {
	ID          int        `db:"id", json:"id"`
	Title       string     `db:"titulo", json:"titulo"`
	Description string     `db:"descripcion", json:"descripcion"`
	ExpireDate  string     `db:"fechaTermino", json:"fecha Expiracion"`
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

/*func Get(db *sqlx.DB, id int) ([]*Activity, error) {
	var (
		Activitys                []*Activity
		CommentsGet          	[]*Comment
		commentCreation time.Time
		comment  string
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
			&a.ExpireDate,
			&a.User,
			&a.Estate,
			&a.Adjuntos,
			&a.Pioridad,
			&commentCreation,
			&comment,
		)

		a.Comments = &CommentsGet{
			FechaCreacion: commentCreation,
			Comentario:    comment,
		}

		if err != nil{
			if err = sql.ErrNoRows{
				return Activitys, nil
			}
			panic(err)
		}
		Activitys = append(Activitys, a)
	}

	if err := rows.Err(); err != nil{
		panic(err)
	}

	return Activitys, nil
}
*/
