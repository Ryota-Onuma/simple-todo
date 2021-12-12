package config
import "github.com/kelseyhightower/envconfig"

type (
	Config struct {
		Env          string `required:"true" envconfig:"ENV" default:"local"`
		Database Database
	}
	Database struct {
		URL       string `required:"true" envconfig:"PG_DATABASE_URL"`
	}
)

func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)

	return cfg, err
}