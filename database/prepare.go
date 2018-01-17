package database

import (
	"database/sql"
)

type queryPrepareTable struct {
	prepareDropletTemplateTable *sql.Stmt
	prepareDropletGroupTable    *sql.Stmt
}

var qPrepareTable queryPrepareTable

func InitPrepare(conn *sql.DB) error {
	var err error
	dbConn := conn

	qPrepareTable.prepareDropletTemplateTable, err = dbConn.Prepare(`
		CREATE TABLE IF NOT EXISTS
		droplet_template (
			id INTEGER PRIMARY KEY,
			name TEXT,
			size TEXT,
			region TEXT,
			image_id INT,
			ssh_key TEXT,
			status INTEGER,
			create_time DATETIME
		)
	`)

	if err != nil {
		return err
	}

	qPrepareTable.prepareDropletGroupTable, err = dbConn.Prepare(`
		CREATE TABLE IF NOT EXISTS
		droplet_group (
			id INTEGER PRIMARY KEY,
			droplet_template_id INTEGER,
			name TEXT
		)	
	`)

	if err != nil {
		return err
	}

	return nil
}
