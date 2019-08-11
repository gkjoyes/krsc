package appcontext

import (
	"github.com/george-kj/krsc/conf"
	"github.com/george-kj/krsc/repo/backend"
)

// AppCtx represents the application context that hold references to backend, conf, logger etc.
type AppCtx struct {
	backend *backend.Backend
	conf    *conf.Vars
}

// Backend returns the backend instance which holds the various repo in the app.
func (c AppCtx) Backend() *backend.Backend {
	return c.backend
}

// Conf returns the conf instance.
func (c AppCtx) Conf() *conf.Vars {
	return c.conf
}

// New initializes the handlers package and sets the read-only instance appCtx.
func New(b *backend.Backend, c *conf.Vars) *AppCtx {
	return &AppCtx{backend: b, conf: c}
}
