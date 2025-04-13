package bootstrap

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func NewMysqlDatabase(env *Env) *sqlx.DB {
	//dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	log.Printf("dbHost:%s, dbPort:%s\n", env.DBHost, env.DBPort)
	cfg := mysql.Config{
		User:                 env.DBUser,
		Passwd:               env.DBPass,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", env.DBHost, env.DBPort),
		DBName:               env.DBName,
		AllowNativePasswords: true,
	}

	db, err := sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("connect DB failed, err:", err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	return db
}

func CloseMysqlDbConnection(client *sqlx.DB) {
	err := client.Close()
	if err != nil {
		log.Fatal("close connect DB failed, err:", err)
	}

	log.Println("Connection to MysqlDB closed.")
}
