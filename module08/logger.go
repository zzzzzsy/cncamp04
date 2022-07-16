package module08

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HttpServerLog struct {
	Request    http.Request
	StatusCode int
}

func httpServerLog(hs HttpServerLog) {
	log.Infof("Request is from %s\n", hs.Request.RemoteAddr)
	log.Infof("Response code is %d\n", hs.StatusCode)
}
