package configurer

import "go.uber.org/fx"

const configsPath = "configs/"

var (
	Constructors = fx.Provide(
		ProvideAppCfg,
		ProvideServerCfg,
	)
)
