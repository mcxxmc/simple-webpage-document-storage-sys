package logging

import "go.uber.org/zap"

// the default logger
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

// Info logs the info
func Info(msg string, sss ...SS) {
	logger.Info(msg, convert(sss)...)
}

// Error logs the error
func Error(err error, sss ...SS) {
	fields := append([]zap.Field{zap.Error(err)}, convert(sss)...)
	logger.Error("", fields...)
}

// Fatal logs the error and stops the program
func Fatal(err error, sss ...SS) {
	fields := append([]zap.Field{zap.Error(err)}, convert(sss)...)
	logger.Fatal("", fields...)
}

// ConditionalLogError logs the error if the error is not nil
func ConditionalLogError(err error, sss ...SS) {
	if err != nil {
		Error(err, sss...)
	}
}
