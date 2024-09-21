package log

import "go.uber.org/zap"

var Logger *zap.Logger

func ConfigureLogger() {
	var err error

	Logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
}
