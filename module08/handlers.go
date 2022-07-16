package module08

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

func healthz(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defaultHandler(rw, r, "Hello LiveRamp SRE")
}

func setResponseHeader(rw http.ResponseWriter, r *http.Request) {
	if v, exist := os.LookupEnv("VERSION"); exist {
		rw.Header().Add("User-Version", v)
	} else {
		log.Infof("The env variable %s does not exist\n", v)
	}

	for name, values := range r.Header {
		for _, value := range values {
			log.Debugf("Request header %s: %s\n", name, value)
			rw.Header().Set(name, value)
		}
	}

	if log.GetLevel() == log.DebugLevel {
		for name, values := range rw.Header() {
			for _, value := range values {
				log.Debugf("Response header %s: %s\n", name, value)
			}
		}
	}
}

func ip(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defaultHandler(rw, r, getIP(r))
}

func defaultHandler(rw http.ResponseWriter, r *http.Request, msg string) {
	setResponseHeader(rw, r)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(msg))
	logs := HttpServerLog{
		Request:    *r,
		StatusCode: http.StatusOK,
	}
	httpServerLog(logs)
}
