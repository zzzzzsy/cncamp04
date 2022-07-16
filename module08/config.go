package module08

import (
	"os"

	"github.com/creamdog/gonfig"
	log "github.com/sirupsen/logrus"
)

const (
	DEFAULT_CONF     = "module08/conf/config.json" // run local only. please provide env variable CONFIG_PATH in your dockerfile
	TIMESTAMP_FORMAT = "2006-01-02 15:04:05"
)

type Config struct {
	ServerConfig ServerConfig
	LogLevel     string
}

type ServerConfig struct {
	Port    string
	Host    string
	Version string
}

func (c *Config) initConfig() (*Config, error) {
	cp := GetenvWithFallback("CONFIG_PATH", DEFAULT_CONF)
	f, err := os.Open(cp)
	if err != nil {
		return c, err
	}
	defer f.Close()

	cfg, err := gonfig.FromJson(f)
	if err != nil {
		return c, err
	}

	var sc ServerConfig
	if err := cfg.GetAs("server", &sc); err != nil {
		return c, err
	}
	os.Setenv("VERSION", sc.Version)

	ll, err := cfg.GetString("loglevel", "info")

	if err != nil {
		return c, err
	}

	setLogLevel(ll)

	c = &Config{
		ServerConfig: sc,
		LogLevel:     ll,
	}

	return c, nil
}

func setLogLevel(lvl string) {
	ll, err := log.ParseLevel(lvl)
	if err != nil {
		ll = log.InfoLevel
	}
	old := log.GetLevel()
	// set global log level
	if ll != old {
		log.SetLevel(ll)
		log.Debugf("Switch log level from %s to %s", old, ll)
	}

	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = TIMESTAMP_FORMAT
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
}
