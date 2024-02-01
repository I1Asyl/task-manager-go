package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/I1Asyl/task-manager-go/configuration"
	"github.com/I1Asyl/task-manager-go/database"
	"github.com/I1Asyl/task-manager-go/pkg/handler"
	"github.com/I1Asyl/task-manager-go/pkg/repositories"
	"github.com/I1Asyl/task-manager-go/pkg/services"

	log "github.com/sirupsen/logrus"
)

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
	config := configuration.New()

	db, err := database.NewConnection(config.Database)
	if err != nil {
		return err
	}

	repository := repositories.New(db)

	go handleShutdown(db)

	services := services.New(repository)

	h := handler.New(services)
	r := h.Assign()
	if err = r.Run(fmt.Sprintf(":%d", config.Server.Port)); err != nil {
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
