package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/george-kj/krsc/app"
	"github.com/george-kj/krsc/conf"
	_ "github.com/lib/pq"
)

const (
	confPath = "conf/conf.toml"
	appName  = "Go Layout"
)

func main() {

	// Read app configuration.
	conf := conf.Read(appName, confPath)

	// Initialize app.
	r, err := app.Init(conf)
	if err != nil {
		log.Println("Error initializing app:- ", err)
		os.Exit(1)
	}

	srv := &http.Server{
		Addr:         conf.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * time.Duration(conf.GraceTimeout),
		Handler:      r,
	}

	log.Printf("%s: Started: Listening on: %v", appName, conf.Port)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
