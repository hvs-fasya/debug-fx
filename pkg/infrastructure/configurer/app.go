package configurer

import "os"

type AppCfg struct {
	Env string
}

func ProvideAppCfg() *AppCfg {
	c := AppCfg{}
	c.Env = os.Getenv("ENV")
	return &c
}
