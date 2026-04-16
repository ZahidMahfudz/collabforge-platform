package config

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	env := GetEnv("APP_ENV")

	// output ke terminal
	Log.SetOutput(os.Stdout)

	// format khusus terminal (simple & readable)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,  // warna di terminal
	})

	// level log
	switch strings.ToLower(env) {
	case "development":
		Log.SetLevel(logrus.DebugLevel)
	case "production":
		Log.SetLevel(logrus.InfoLevel)
	default:
		Log.SetLevel(logrus.WarnLevel)
	}
}