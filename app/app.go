package app

import (
	"io"
	"log"
	"os"

	"github.com/george-kj/krsc/conf"
	"github.com/george-kj/krsc/handler"
	"github.com/george-kj/krsc/repo/backend"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// Init will initialize our application.
func Init(conf *conf.Vars) (*mux.Router, error) {

	// Initialize logging file.
	f, err := os.OpenFile(conf.Log.Error, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, errors.Errorf("Failed to open log file %s: %v", conf.Log.Error, err.Error())
	}
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	// Establish connection with postgres database.
	db, err := backend.Connect(
		conf.Postgres.Host,
		conf.Postgres.User,
		conf.Postgres.Pass,
		conf.Postgres.DB,
		conf.Postgres.Port,
	)
	if err != nil {
		return nil, err
	}

	// Init handlers package.
	repo := backend.Init(db)
	handler.Init(repo, conf)

	return Routes(), nil
}
