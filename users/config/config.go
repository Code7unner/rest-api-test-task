package config

import "github.com/caarlos0/env"

// Config all app variables are stored here
type Config struct {
	// Postgres db connection
	Username  string `env:"DB_USERNAME"`
	Password  string `env:"DB_PASSWORD"`
	Host      string `env:"DB_HOST" envDefault:"db"`
	Port      int    `env:"DB_PORT" envDefault:"5432"`
	DbName    string `env:"DB_NAME" envDefault:"test_db"`
	SslEnable bool   `env:"DB_SSL_ENABLE" envDefault:"false"`

	DebugMode        bool   `env:"DEBUG_MODE" envDefault:"false"`
	JWTSecret        string `env:"JWT_SECRET" envDefault:"secret"`
	ExpiresInMinutes int    `env:"EXPIRES_IN_MINUTES" envDefault:"600"`
	ServerAddress    int    `env:"SERVER_ADDRESS" envDefault:"8080"`
}

// New returns a new Config struct
func New() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
