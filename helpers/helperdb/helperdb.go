package helperdb

import "github.com/jmoiron/sqlx"

func GetDatabase() *sqlx.DB {
	return sqlx.MustConnect("mysql", "root:volcom..@tcp(localhost:3306)/DB_INJUV_Backoffice?charset=utf8mb4,utf8&parseTime=true")
}
