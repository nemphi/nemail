package nemail

import (
	"github.com/emersion/go-smtp"
)

type App struct {
	cfg         *Config
	smtpServer  *smtp.Server
	smtpBackend *smtp.Backend
	db          *DB
}

func New(options ...Option) (*App, error) {
	app := &App{}
	for _, opt := range options {
		err := opt(app)
		if err != nil {
			return nil, err
		}
	}
	return app, nil
}

func (app *App) Start() error {
	return nil
}

func (app *App) Stop() error {
	return nil
}
