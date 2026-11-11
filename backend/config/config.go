package config

type (
	AppConfig struct {
		DataBase   *DatabaseConfig
		CorsConfig *CorsConfig
		Http       *HttpConfig
	}
	DatabaseConfig struct {
		Host         string `env:"DATABASE_HOST,required"`
		Port         string `env:"DATABASE_PORT,required"`
		UserName     string `env:"DATABASE_USERNAME,required"`
		Password     string `env:"DATABASE_PASSWORD,required"`
		Name         string `env:"DATABASE_NAME,required"`
		SSLMode      string `env:"DATABASE_SSL_MODE"            envDefault:"disable"`
		Timezone     string `env:"DATABASE_TIMEZONE"             envDefault:"Asia/Yangon"`
		MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS"     envDefault:"10"`
		MaxIdleConns int    `env:"DB_MAX_IDLE_CONNS"     envDefault:"2"`
	}
	CorsConfig struct {
		AllowOrigins string `env:"CORS_ALLOW_ORIGINS,required"`
		AllowMethods string `env:"CORS_ALLOW_METHODS,required"`
	}
	HttpConfig struct {
		Port     int    `env:"HTTP_PORT,required"`
		Timezone string `env:"TZ" envDefault:"Asia/Yangon"`
	}
)
