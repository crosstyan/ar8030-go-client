package log

import "go.uber.org/zap"

var logger *zap.Logger
var sugar *zap.SugaredLogger

func Sugar() *zap.SugaredLogger {
	if sugar == nil {
		logger = Logger()
		sugar = logger.Sugar()
		return sugar
	}
	return sugar
}

func Logger() *zap.Logger {
	if logger == nil {
		logger, _ = zap.NewDevelopment()
		sugar = logger.Sugar()
		return logger
	}
	return logger
}
