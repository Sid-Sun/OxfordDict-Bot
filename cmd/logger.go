package main

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func initLogger(env string) {
	var err error

	if env == dev {
		logger, err = zap.NewDevelopmentConfig().Build()
	} else {
		logger, err = zap.NewProductionConfig().Build()
	}

	if err != nil {
		panic(err)
	}
}
