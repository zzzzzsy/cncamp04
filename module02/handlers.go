package module02

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func healthz(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setResponseHeader(rw, r.Header)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Hello LiveRamp SRE"))
	logs := HttpServerLog{
		Request:    *r,
		StatusCode: http.StatusOK,
	}
	httpServerLog(logs)
}

func setResponseHeader(rw http.ResponseWriter, h http.Header) {
	for name, values := range h {
		for _, value := range values {
			log.Printf("Request header %s: %s\n", name, value)
			rw.Header().Set(name, value)
		}
	}

	if v, exist := os.LookupEnv("VERSION"); exist {
		rw.Header().Add("User-Version", v)
	} else {
		log.Printf("The env variable %s does not exist\n", v)
	}
}
