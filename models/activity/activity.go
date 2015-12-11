package activity

import (
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	sqlInsertActivity = "INSERT INTO tareas (titulo, descripcion, fechaTermino, miembros_id, estado, adjuntos_id, pioridad) VALUES (:titulo, :descripcion, :fechaTermino, :miembros_id, :estado, ?);"
	sqlInsertComments = ""
)

type Activity struct {
	ID          int         `db:"id", json:"id"`
	Title       string      `db:"titulo", json:"titulo"`
	Description string      `db:"descripcion", json:"descripcion"`
	ExpireDate  time.Time   `db:"fechaTermino", json:"fecha Expiracion"`
	User        int         `db:"miembros_id", json:"Recurso"`
	Estate      string      `db:"estado", json:"Estado"`
	Adjuntos    int         `db:"adjuntos_id", json:"adjuntos"`
	Pioridad    string      `db:"pioridad", json:"pioridad"`
	Comment     []*Comments `db:"comentarios", json:"Comentarios"`
}

type Comments struct {
	ID            int       `db:"id", json:"id"`
	IDTarea       int       `db:"tareas_id", json:"idTarea"`
	FechaCreacion time.Time `db:"fechaCreacion", json:"FechaCreacion"`
	Comentario    string    `db:"comentarioscol", json:"Comentario"`
}

func (a Activity) Create(db *sqlx.DB) {

	_, err := db.NamedExec(sqlInsertActivity, &a)
	if err != nil {
		panic(err)
	}
}
