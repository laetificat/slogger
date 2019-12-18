# sLogger
A simple Zap/Sentry logger.

This project makes it easy to log using Zap in combination with Sentry and Loggly.

## Usage
To use this project run `go get -u github.com/laetificat/slogger`. Configure the loggers first by setting the config:
```go
c := slogger.Config{
    Level:  "Info",
    Sentry: slogger.SentryZapOptions{
        Enabled: true,
        Dsn:     "https://somesentrycode.sentry.io",
    },
    Loggly: slogger.LogglyZapOptions{
        Enabled: true,
        Token:   "somelogglyid",
    },
}
slogger.SetConfig(c)

slogger.Error("This is my error message that will appear in Sentry.")
slogger.Info("This is my info message that will appear in Loggly.")
```

## Contributing
See [CONTRIBUTING.md](CONTRIBUTING.md)

## License
See [LICENSE.md](LICENSE.md)