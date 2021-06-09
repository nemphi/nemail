package nemail

type Option func(*App) error

func WithConfig(cfg *Config) Option {
	return func(app *App) error {
		app.cfg = cfg
		return nil
	}
}

func WithDB(db *DB) Option {
	return func(app *App) error {
		app.db = db
		return nil
	}
}
