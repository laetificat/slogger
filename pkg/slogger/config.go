package slogger

type Config struct {
	Level  string
	Sentry SentryZapOptions
	Loggly LogglyZapOptions
}

var cfg = Config{
	Level: "Info",
}

/*
SetConfig sets the cfg var to the given Config.
*/
func SetConfig(c Config) {
	cfg = c
}
