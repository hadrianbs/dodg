package dropletgroup

import (
	"github.com/digitalocean/godo"
	"github.com/gorilla/mux"
)

type DropletGroupOp struct {
	ID              int64
	Name            string
	DropletTemplate string //TODO DropletTemplate Struct
}

type DropletGroupService interface {
	Create()
	New()
}

var doClient *godo.Client

func Init(router *mux.Router, client *godo.Client) {
	doClient = client
}
