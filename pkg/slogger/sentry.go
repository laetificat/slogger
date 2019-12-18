package slogger

import (
	"encoding/json"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap/zapcore"
)

type SentryZapOptions struct {
	Enabled bool
	Dsn     string
}

type SentryZap struct {
	Message string
	Options SentryZapOptions
}

var sentryZap SentryZap

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
func (s SentryZap) Sync() error {
	if sentryZap.Message != "" {
		hub := sentry.CurrentHub()
		if hub == nil || hub.Client() == nil || hub.Client().Options().Dsn != sentryZap.Options.Dsn {
			if err := sentry.Init(sentry.ClientOptions{
				Dsn: sentryZap.Options.Dsn,
			}); err != nil {
				return err
			}
		}

		defer sentry.Flush(0)

		event := sentry.NewEvent()
		if err := json.Unmarshal([]byte(sentryZap.Message), &event); err != nil {
			panic(err)
		}

		sentry.CaptureEvent(event)
	}

	return nil
}

/*
Write writes the given bytes to SentryZap.Message as a string.
*/
func (s SentryZap) Write(data []byte) (n int, err error) {
	sentryZap.Message = string(data)
	return 0, nil
}
