package droplettemplate

import (
	"database/sql"
	"log"
	"time"

	"github.com/digitalocean/godo"
)

type DropletTemplate struct {
	Id       int
	Name     string
	Size     string
	Region   string
	Image    int
	ssh_keys string
}

type DropletTemplateInterface interface {
	Create() (DropletTemplate, error)
	Delete() (DropletTemplate, error)
	Get() (DropletTemplate, error)
	List() ([]DropletTemplate, error)
}

type DropletTemplateOp struct {
	DOClient *godo.Client
	Db       *sql.DB
}

var (
	dropletTemplateOp DropletTemplateOp
	qDropletTemplate  queryDropletTemplate
)

func Init(client *godo.Client, dbClient *sql.DB) {
	var err error
	dropletTemplateOp = DropletTemplateOp{
		DOClient: client,
		Db:       dbClient,
	}
	qDropletTemplate, err = initStmt(dropletTemplateOp.Db)
	if err != nil {
		log.Println("Error initializing prepared statements : ", err)
		panic(err)
	}
}

func (d *DropletTemplateOp) Create() {
	qDropletTemplate.createNewDropletTemplate.Exec(
		"TestDropletTemplate",
		"512mb",
		"sgp-1",
		1,
		"wtf",
		1,
		time.Now(),
	)
}

func (d *DropletTemplateOp) Delete() {
}

func (d *DropletTemplateOp) Get() {

}

func (d *DropletTemplateOp) List() {

}
