package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/I1Asyl/task-manager-go/database"
	"github.com/I1Asyl/task-manager-go/docs"
	"github.com/I1Asyl/task-manager-go/pkg/handler"
	"github.com/I1Asyl/task-manager-go/pkg/repositories"
	"github.com/I1Asyl/task-manager-go/pkg/services"
	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

//	@title			Task manager
//	@version		1.0
//	@description	Task manager for a developing company.

//	@contact.name	Task manager
//	@contact.email	altayyerassyl@gmail.com

// @securityDefinitions.apikey.in header
// @securityDefinitions.apikey.name Authorization

//	@host		localhost:8080
//	@BasePath	/

func init() {
	log.SetOutput(os.Stdout)
}

func main() {

	log.Infof("Starting server...")

	if err := run(); err != nil {
		log.Fatalf("Cannot start server, err: %v", err)
	}

}

func run() error {
	godotenv.Load(".env")
	docs.SwaggerInfo.BasePath = "/"

	db, err := database.NewConnection(os.Getenv("DSN"))
	if err != nil {
		return err
	}
	repository := repositories.New(db)

	go handleShutdown(db)

	services := services.New(repository)

	h := handler.New(services)
	r := h.Assign()
	if err = r.Run(fmt.Sprintf(":%d", 8000)); err != nil {
		return err
	}

	return nil
}

func handleShutdown(db *sql.DB) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	// handle ctrl+c event here
	// for example, close database
	log.Warn("Closing DB connection before complete shutdown")

	if err := db.Close(); err != nil {
		log.Errorf("error while closing the connection to the database: %v", err)
	}

	os.Exit(0)
}
