package activity

import "time"


const (
	sqlInsertUser := ""
)

type Activity struct {
	ID          int        `db:"id", json:"id"`
	Description string     `db:"descripcion", json:"descripcion"`
	ExpireDate  time.Time  `db:"fechaTermino", json:"fecha Expiracion"`
	User        int        `db:"miembros_id", json:"Recurso"`
	Estate      int        `db:"estado_id", json:"Estado"`
	Comment     []*Comments `db:"comentarios", json:"Comentarios"`
}

type Comments struct {
	ID            int       `db:"id", json:"id"`
	IDTarea       int       `db:"tareas_id", json:"idTarea"`
	FechaCreacion time.Time `db:"fechaCreacion", json:"FechaCreacion"`
	Comentarios   string    `db:"comentarioscol", json:"Comentario"`
}


func (a Activity) Create(db *sqlx.DB) {

	_, err := db.NamedExec(sqlInsertUser, &a)
	if err != nil {
		panic(err)
	}
}