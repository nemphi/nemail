package nemail

import (
	"log"

	"github.com/emersion/go-smtp"
)

type App struct {
	cfg         *Config
	smtpServer  *smtp.Server
	smtpBackend *Backend
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

func (app *App) Start() {
	b := &Backend{}
	s := smtp.NewServer(b)
	app.smtpBackend = b
	app.smtpServer = s

	err := s.ListenAndServeTLS()
	if err != nil {
		log.Fatal(err)
	}

}

func (app *App) Stop() error {
	return app.smtpServer.Close()
}
