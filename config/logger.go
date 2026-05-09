package config

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger() {
	env := GetEnv("APP_ENV")

	// output ke terminal
	Logger.SetOutput(os.Stdout)

	// format khusus terminal (simple & readable)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,  // warna di terminal
	})

	// level log
	switch strings.ToLower(env) {
	case "development":
		Logger.SetLevel(logrus.DebugLevel)
	case "production":
		Logger.SetLevel(logrus.InfoLevel)
	default:
		Logger.SetLevel(logrus.WarnLevel)
	}

	Logger.Infof("Logger initialized on %s level in %s mode", Logger.Level.String(), env)
}