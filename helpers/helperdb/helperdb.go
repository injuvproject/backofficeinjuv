package helperdb

import "github.com/jmoiron/sqlx"

func GetDatabase() *sqlx.DB {
	return sqlx.MustConnect("mysql", "root:root@tcp(localhost:3306)/go-vipex?charset=utf8mb4,utf8&parseTime=true")
}
