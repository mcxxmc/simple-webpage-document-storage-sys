package logging

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction(zap.AddCallerSkip(1))
	logger.Info("Logger initialized.")
}

// Sync synchronizes the logger; usually follows "defer"
func Sync() {
	logger.Sync()
	logger.Info("Logger synchronized.")
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
	logger.Info(msg)
}

func InfoInt(msg string, key string, val int) {
	logger.Info(msg, zap.Int(key, val))
}

// Error logs the error
func Error(err error, msg ...string) {
	logger.Error(processMsg(msg), zap.Error(err))
}

func Fatal(err error, msg ...string) {
	logger.Fatal(processMsg(msg), zap.Error(err))
}

// ConditionalLogError logs the error if the error is not nil
func ConditionalLogError(err error, msg ...string) {
	if err != nil {
		logger.Error(processMsg(msg), zap.Error(err))
	}
}
