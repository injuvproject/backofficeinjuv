package activity

import "time"

type Activity struct {
	ID          int       `db:"id", json:"id"`
	Description string    `db:"descripcion", json:"descripcion"`
	ExpireDate  time.Time `db:"fechaTermino", json:"fecha Expiracion"`
	User        int       `db:"miembros_id", json:"Recurso"`
	Estate      int       `db:"estado_id", json:"Estado"`
}

type Estate struct {
}

type Comments struct {
}
