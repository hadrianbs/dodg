package main

import (
	"database/sql"
	"net/http"
	"os"

	grace "gopkg.in/paytm/grace.v1"

	"github.com/gorilla/mux"

	"log"

	"github.com/hadrianbs/dodg/database"
	"github.com/hadrianbs/dodg/doclient"
	"github.com/hadrianbs/dodg/droplettemplate"

	"github.com/digitalocean/godo"
	"github.com/hadrianbs/dodg/config"
)

type AppInit struct {
	ServiceConfig config.ServiceConfig
	doClient      *godo.Client
	dbClient      *sql.DB
}

var (
	configuration *AppInit
)

func main() {
	configuration = new(AppInit)
	initConfig()
	initPackages(configuration)
	initRoutes()
}

//Initialize configuration, digital ocean client, and database connection
func initConfig() {
	sConfig := config.Init()
	doClient, err := doclient.InitClient(sConfig.PAT)
	if err != nil {
		log.Println("Error initializing digitalocean client : ", err)
		os.Exit(1)
	}
	dbClient := database.InitDB()
	configuration.ServiceConfig = sConfig
	configuration.doClient = doClient
	configuration.dbClient = dbClient
}

func initPackages(configuration *AppInit) {
	droplettemplate.Init(configuration.doClient, configuration.dbClient)
}

func initRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/ping", handlePing).Methods("GET")

	/*DropletTemplate Routes*/
	log.Fatal(grace.Serve(":9009", router))
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG!"))
}
