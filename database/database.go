package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	//Create new database file IF there's not existing database
	_, err := os.Stat("/etc/dodg/dodg.db")
	if err != nil {
		os.MkdirAll("/etc/dodg/", 0755)
		os.Create("/etc/dodg/dodg.db")
	}

	db, err := sql.Open("sqlite3", "/etc/dodg/dodg.db")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	err = InitPrepare(db)
	if err != nil {
		panic(err)
	}
	qPrepareTable.prepareDropletGroupTable.Exec()
	qPrepareTable.prepareDropletTemplateTable.Exec()
	return db
}
