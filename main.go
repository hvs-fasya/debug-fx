package main

import (
	"go.uber.org/fx"

	"github.com/hvs-fasya/debug-fx/pkg/infrastructure/configurer"
	"github.com/hvs-fasya/debug-fx/pkg/infrastructure/logger"
	"github.com/hvs-fasya/debug-fx/pkg/server"
)

func main() {
	app := fx.New(
		fx.NopLogger,
		configurer.Constructors,
		fx.Provide(logger.NewLogger),
		fx.Invoke(server.Run),
	)
	app.Run()
}
