package droplettemplate

import (
	"database/sql"
)

type queryDropletTemplate struct {
	createNewDropletTemplate *sql.Stmt
	deleteDropletTemplate    *sql.Stmt
	getDropletTemplate       *sql.Stmt
	listDropletTemplates     *sql.Stmt
}

func initStmt(conn *sql.DB) (queryDropletTemplate, error) {
	var qDropletTemplate queryDropletTemplate
	var err error
	dbConn := conn

	qDropletTemplate.createNewDropletTemplate, err = dbConn.Prepare(`
		INSERT INTO droplet_template (
			name,
			size,
			region,
			image_id,
			ssh_key,
			status,
			create_time
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)
	`)
	if err != nil {
		panic(err)
	}

	qDropletTemplate.deleteDropletTemplate, err = dbConn.Prepare(`
		UPDATE 
			droplet_template
		SET 
			status = 0
		WHERE
			id = ?
	`)
	if err != nil {
		panic(err)
	}

	qDropletTemplate.getDropletTemplate, err = dbConn.Prepare(`
		SELECT *
		FROM 
			droplet_template
		WHERE
			id = ?
	`)
	if err != nil {
		panic(err)
	}
	return qDropletTemplate, nil
}
