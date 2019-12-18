# sLogger
A simple Zap/Sentry logger.

This project makes it easy to log using Zap in combination with Sentry, this project uses Viper and the following items 
should be set:
```toml
[log]
    default_level = "info"                                          # Default log level, info, debug, error, fatal (string)
  
    [sentry]
        enabled = false                                             # Enable sending to Sentry (bool)
        dsn = "https://randomstringhere@sentry.io/andsomeidhere"    # DSN link (string)
```

## Usage
To use this project run `go get -u github.com/laetificat/slogger` in your project and use `slogger.Info("Test message")`.

## Contributing
See [CONTRIBUTING.md](CONTRIBUTING.md)

## License
See [LICENSE.md](LICENSE.md)