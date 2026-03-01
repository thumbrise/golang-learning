package env

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

var ErrRead = "read config"

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

func (*Loader) MustLoad(cfg interface{}) {
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("%s: %s", ErrRead, err)
	}
}
