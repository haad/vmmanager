package log

import (
	"log"

	"go.uber.org/zap"
)

var Slog *zap.SugaredLogger
var logger *zap.Logger

func InitLogger(debug bool) {
	var err error
	var logger *zap.Logger

	if debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	Slog = logger.Sugar()

	defer logger.Sync()
}
