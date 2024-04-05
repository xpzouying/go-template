package cli

import (
	"github.com/xpzouying/go-cmd-project-template/internal/config"
	"github.com/xpzouying/go-cmd-project-template/log"
)

func InitConfigAndComponents() (*config.Config, error) {

	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	if err := initLogger(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadConfig() (*config.Config, error) {

	cfg := &config.Config{
		ListenAddr: listenAddr,
		LogLevel:   logLevel,
	}

	return cfg, nil
}

func initLogger(cfg *config.Config) error {
	if err := log.InitGlobalLogger(cfg.LogLevel); err != nil {
		return err
	}
	return nil
}
