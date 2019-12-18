package slogger

import (
	"github.com/segmentio/go-loggly"
)

type LogglyZapOptions struct {
	Enabled bool
	Token   string
}

type LogglyZap struct {
	Message string
	Client  *loggly.Client
	Options LogglyZapOptions
}

var logglyZap LogglyZap

/*
Sync sends the message that resides in LogglyZap.Message to Loggly if it's not empty.
*/
func (l LogglyZap) Sync() error {
	if logglyZap.Options.Enabled && logglyZap.Message != "" {
		if logglyZap.Client == nil {
			logglyZap.Client = loggly.New(logglyZap.Options.Token)
		}

		if err := logglyZap.Client.Info(logglyZap.Message); err != nil {
			return err
		}

		if err := logglyZap.Client.Flush(); err != nil {
			return err
		}
	}

	return nil
}

/*
Write writes the given bytes to LogglyZap.Message as a string.
*/
func (l LogglyZap) Write(data []byte) (n int, err error) {
	logglyZap.Message = string(data)
	return 0, nil
}
