package app

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func Launch() {
	logger := logger()
	fx.New(Module,
		fx.Provide(func() *zap.Logger { return logger }),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		})).Run()
}
