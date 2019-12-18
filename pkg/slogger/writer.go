package slogger

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// globalZapCore is the global variable to hold an already initiated core
var globalZapCore *zapcore.Core

/*
Info logs informational messages, these are not really interesting and could be ignored.
*/
func Info(message string) {
	lg := zap.New(newZapCore())
	defer finishLogging(lg.Sync())

	lg.Info(message)
}

/*
Debug logs additional messages, these are interesting for debugging the application.
*/
func Debug(message string) {
	lg := zap.New(newZapCore())
	defer finishLogging(lg.Sync())

	lg.Debug(message)
}

/*
Warning logs things that went wrong but are not impacting the process, but it should be looked at.
*/
func Warning(message string) {
	lg := zap.New(newZapCore())
	defer finishLogging(lg.Sync())

	lg.Warn(message)
}

/*
Error logs errors that happened, this impacted the process and need to be looked at.
*/
func Error(message string) {
	lg := zap.New(newZapCore())
	defer finishLogging(lg.Sync())

	lg.Error(message)
}

/*
Fatal logs fatal errors, when this happens the process has stopped. This is worst case scenario.
*/
func Fatal(message string) {
	lg := zap.New(newZapCore())
	defer finishLogging(lg.Sync())

	lg.Fatal(message)
}

func newZapCore() zapcore.Core {
	if globalZapCore != nil {
		return *globalZapCore
	}

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	sentryErrors := zapcore.Lock(SentryZap)

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	jsonConfig := zap.NewProductionEncoderConfig()
	jsonConfig.MessageKey = "message"
	jsonConfig.TimeKey = "timestamp"
	jsonConfig.EncodeTime = EpochTimeEncoderInt64
	jsonEncoder := zapcore.NewJSONEncoder(jsonConfig)

	sentryHighPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return viper.GetBool("logger.sentry.enabled") && lvl >= zapcore.ErrorLevel
	})
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel || lvl == zapcore.WarnLevel
	})
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return viper.GetString("log.default_level") == "debug" && lvl == zapcore.DebugLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleDebugging, debugPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(jsonEncoder, sentryErrors, sentryHighPriority),
	)

	globalZapCore = &core

	return *globalZapCore
}

// finishLogging eats any errors returned from lg.Sync() to silence linters
func finishLogging(errs ...error) {}
