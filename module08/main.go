package module08

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/tylerb/graceful.v1"
)

func init() {
	// set default log level to info
	// you can change it in config.json
	setLogLevel("info")
}

func WebServer() {
	var config Config
	if c, err := config.initConfig(); err != nil {
		log.WithError(err).Error()
	} else {
		log.WithFields(log.Fields{
			"host":    c.ServerConfig.Host,
			"port":    c.ServerConfig.Port,
			"version": c.ServerConfig.Version,
		}).Debug("Server Configurations")

		log.Info("Log level is ", c.LogLevel)

		router := NewRouter(AllRoutes())

		// ListenAndServe is equivalent to http.Server.ListenAndServe with graceful shutdown enabled
		// timeout is the duration to wait until killing active requests and stopping the server.
		// If timeout is 0, the server never times out. It waits for all active requests to finish.
		srv := &graceful.Server{
			Timeout: 10 * time.Second,

			Server: &http.Server{
				Addr:    fmt.Sprintf("%s:%s", c.ServerConfig.Host, c.ServerConfig.Port),
				Handler: router,
			},
		}

		srv.ListenAndServe()
	}
}
