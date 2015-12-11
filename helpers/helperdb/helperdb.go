package helperdb

import "github.com/jmoiron/sqlx"

func GetDatabase() *sqlx.DB {
	return sqlx.MustConnect("mysql", "root:volcom..@tcp(localhost:3306)/db_injuv_backoffice?charset=utf8mb4,utf8&parseTime=true")
}
