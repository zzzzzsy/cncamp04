package module03

import (
	"log"
	"net/http"
)

type HttpServerLog struct {
	Request    http.Request
	StatusCode int
}

func httpServerLog(hs HttpServerLog) {
	log.Printf("Request is from %s\n", hs.Request.RemoteAddr)
	log.Printf("Response code is %d\n", hs.StatusCode)
}
