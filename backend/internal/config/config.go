package config

type Config struct {
}
type Primary struct {
	Env string `koanf:"env" validate:"required"`
}
type ServercConfig struct {
	Port               string   `koanf:"port" validate:"required"`
	ReadTimeOut        int      `koanf:"read_timeout" validate:"required"`
	WriteTimeOut       int      `koanf:"write_timeout" validate:"required"`
	IdleTimeOut        int      `koanf:"idle_timeout" validate:"required"`
	CORSAllowedOrigins []string `koanf:"cors_allowed_origins" validate:"required"`
}
type DBconfig struct {
	Port            int    `koanf:"port" validate:"required"`
	Host            string `koanf:"host" validate:"required"`
	User            string `koanf:"user" validate:"required"`
	Passowrd        string `koanf:"passowrd"`
	Name            string `koanf:"name" validate:"required"`
	SSLMode         string `koanf:"ssl_mode" validate:"required"`
	MaxOpenConns    int    `koanf:"max_open_conns" validate:"required"`
	MaxIdleConns    int    `koanf:"max_idle_conns" validate:"required"`
	ConnMaxLifeTime int    `koanf:"conn_max_lifetime" validate:"required"`
	ConnMaxIdleTime int    `koanf:"conn_max_idletime" validate:"required"`
}
type AuthConfig struct {
	SecretKey string `koanf:"secret_key" validate:"required"`
}
