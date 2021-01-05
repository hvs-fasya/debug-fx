package logger

import (
	"context"
	"strings"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/hvs-fasya/debug-fx/pkg/infrastructure/configurer"
)

type Logger interface {
	Error(string)
	Fatal(string)
	Info(string)
	Debug(string)
	Named(string) Logger
	Sync() error
}

type appLogger struct {
	logger *zap.Logger
}

func NewLogger(lc fx.Lifecycle, appCfg *configurer.AppCfg) Logger {
	var cfg = zap.NewProductionConfig()
	cfg.DisableCaller = true
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05")
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	var level = zapcore.InfoLevel
	if appCfg.Env != "prod" {
		level = zapcore.DebugLevel
		cfg.Encoding = "console"
	}
	cfg.Level = zap.NewAtomicLevelAt(level)
	var l, _ = cfg.Build()
	l.Info("log level " + strings.ToUpper(level.String()))

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			l.Sync()
			return nil
		},
	})
	return &appLogger{logger: l}
}

func (l *appLogger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *appLogger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *appLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *appLogger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *appLogger) Named(name string) Logger {
	named := appLogger{logger: l.logger.Named(name)}
	return &named
}

func (l *appLogger) Sync() error {
	return l.logger.Sync()
}
