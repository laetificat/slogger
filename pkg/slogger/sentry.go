package slogger

import (
	"encoding/json"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type sentryZapStruct struct {
	Message string
}

// SentryZap is the exported variable for the struct.
var SentryZap sentryZapStruct

/*
EpochTimeEncoderInt64 is a time encoder for Zap that encodes time to int64 instead of float64.
*/
func EpochTimeEncoderInt64(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	nanos := t.UnixNano()
	sec := float64(nanos) / float64(time.Second)
	enc.AppendInt64(int64(sec))
}

/*
Sync sends the message that resides in SentryZap.Message to Senty if it's not empty.
*/
func (s sentryZapStruct) Sync() error {
	hub := sentry.CurrentHub()
	if hub == nil || hub.Client() == nil || hub.Client().Options().Dsn != viper.GetString("logger.sentry.dsn") {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: viper.GetString("logger.sentry.dsn"),
		}); err != nil {
			return err
		}
	}

	if SentryZap.Message != "" {
		defer sentry.Flush(0)

		event := sentry.NewEvent()
		if err := json.Unmarshal([]byte(SentryZap.Message), &event); err != nil {
			panic(err)
		}

		sentry.CaptureEvent(event)
	}

	return nil
}

/*
Write writes the given bytes to SentryZap.Message as a string.
*/
func (s sentryZapStruct) Write(data []byte) (n int, err error) {
	SentryZap.Message = string(data)
	return 0, nil
}
