package module08

import (
	"net"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetenvWithFallback(key string, fallback string) string {
	temp := os.Getenv(key)
	if len(temp) == 0 {
		return fallback
	}
	return temp
}

func getIP(r *http.Request) string {
	var userIP string
	if len(r.Header.Get("CF-Connecting-IP")) > 1 {
		userIP = r.Header.Get("CF-Connecting-IP")
		log.Debug("CF-Connecting-IP ", net.ParseIP(userIP))
	} else if len(r.Header.Get("X-Forwarded-For")) > 1 {
		userIP = r.Header.Get("X-Forwarded-For")
		log.Debug("X-Forwarded-For ", net.ParseIP(userIP))
	} else if len(r.Header.Get("X-Real-IP")) > 1 {
		userIP = r.Header.Get("X-Real-IP")
		log.Debug("X-Real-IP ", net.ParseIP(userIP))
	} else {
		userIP = r.RemoteAddr
		if strings.Contains(userIP, ":") {
			log.Debug("RemoteAddr Host ", net.ParseIP(strings.Split(userIP, ":")[0]))
		} else {
			log.Debug(net.ParseIP(userIP))
		}
	}
	return userIP
}
