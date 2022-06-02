package logger

import "github.com/sirupsen/logrus"

// Log is a custom instance of the logrus logger
var Log = logrus.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	Log.SetFormatter(&logrus.JSONFormatter{})
}
