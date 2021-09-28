package logging

import "go.uber.org/zap"

// InitLog returns a logger
func InitLog() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

// SetGlobal sets the global logging
func SetGlobal(logger *zap.Logger) func() {
	return zap.ReplaceGlobals(logger)
}

// processes the passed in messages
func processMsg(msg []string) string {
	write := ""
	if len(msg) > 0 {
		for _, s := range msg {
			write = write + s
		}
	}
	return write
}

// Info logs the info
func Info(msg string) {
	zap.L().Info(msg)
}

// Error logs the error
func Error(err error, msg ...string) {
	zap.L().Error(processMsg(msg), zap.Error(err))
}

func Fatal(err error, msg ...string) {
	zap.L().Fatal(processMsg(msg), zap.Error(err))
}

// ConditionalLogError logs the error if the error is not nil
func ConditionalLogError(err error, msg ...string) {
	if err != nil {
		zap.L().Error(processMsg(msg), zap.Error(err))
	}
}
