package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func DbConn() *sql.DB {
	dsn := os.Getenv("DB_USERNAME") + ":" +
		os.Getenv("DB_PASSWORD") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	dBConnection, err := sql.Open("mysql", dsn)

	log.Infof("Connect to the database server")
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v, %v", dsn, err)
	}

	err = dBConnection.Ping()
	if err != nil {
		fmt.Println("Ping to database failed!!")
	}

	dBConnection.SetMaxOpenConns(10)
	dBConnection.SetMaxIdleConns(5)
	dBConnection.SetConnMaxLifetime(time.Second * 10)

	return dBConnection
}

// CloseStmt after run stmt
func CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}

// CloseRows when select
func CloseRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}
