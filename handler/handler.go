package handler

import (
	"github.com/george-kj/krsc/conf"
	"github.com/george-kj/krsc/handler/appcontext"
	"github.com/george-kj/krsc/repo/backend"
)

// appctx holds the read-only instance of the application context.
var appctx *appcontext.AppCtx

// Init initializes the handler package by setting the application context.
func Init(b *backend.Backend, c *conf.Vars) {
	appctx = appcontext.New(b, c)
}
