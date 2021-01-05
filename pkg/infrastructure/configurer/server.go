package configurer

import "github.com/BurntSushi/toml"

const serverCfgFileName = "server.toml"

type ServerCfg struct {
	Port    string
	Timeout Duration
}

func ProvideServerCfg() (c *ServerCfg, err error) {
	_, err = toml.DecodeFile(configsPath+serverCfgFileName, &c)
	return
}
