package database

type DatabaseConfig struct {
	Engine   string `env:"DB_ENGINE,default=postgres"`
	Host     string `env:"DB_HOST,default=localhost"`
	Name     string `env:"DB_NAME,default=tr_db"`
	Port     string `env:"DB_PORT,default=5432"`
	User     string `env:"DB_USER,default=root"`
	Password string `env:"DB_PASSWORD,default=pwd"`
	Options  string `env:"DB_OPTIONS"`
}