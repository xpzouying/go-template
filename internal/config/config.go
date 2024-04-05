package config

// Config is a struct that holds the configuration for the application.
type Config struct {
	// ListenAddr is the address to bind the local server to.
	ListenAddr string

	// LogLevel is the log level for the application.
	LogLevel string
}
