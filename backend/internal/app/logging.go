package app

import "go.uber.org/zap"

func logger() *zap.Logger {
	log, _ := zap.NewDevelopment()
	return log
}
