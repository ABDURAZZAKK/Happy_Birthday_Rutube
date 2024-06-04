package app

import (
	"github.com/ABDURAZZAKK/Happy_Birthday_Rutube/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/t-tomalak/logrus-prefixed-formatter"
)

func SetLogger() {
	lvl := config.LOG_LEVEL
	// LOG_LEVEL not set, let's default to debug
	if lvl == "" {
		lvl = "debug"
	}
	// parse string, this is built-in feature of logrus
	ll, err := log.ParseLevel(lvl)
	if err != nil {
		ll = log.DebugLevel
	}
	// set global log level
	log.SetLevel(ll)
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&prefixed.TextFormatter{
		// DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}
