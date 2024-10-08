package configuration

import "github.com/Netflix/go-env"

type DBConfig struct {
	Host string `env:"DB_HOST"`
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASS"`
	Name string `env:"DB_NAME"`
}

type FSConfig struct {
	UploadFolder string `env:"UPLOAD_FOLDER"`
}

type ServerConfig struct {
	HttpSocket string `env:"HTTP_SOCKET"`
	GrpcSocket string `env:"GRPC_SOCKET"`
}

type SMSConfig struct {
	ApiKey    string `env:"SMS_API_KEY"`
	ApiDomain string `env:"SMS_API_DOMAIN"`
}

type Config struct {
	DB     DBConfig
	Server ServerConfig
	SMS    SMSConfig
	FS     FSConfig
}

func Load() (*Config, error) {
	var cfg Config
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
