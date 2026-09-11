package util

import (
	model "GitHub/management-/Model"
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger

func init() {
	Logger = *logrus.New()
	Logger.Out = os.Stdout
	Logger.SetFormatter(&logrus.TextFormatter{
		//DisableTimestamp:  false,
		TimestampFormat: "2006-01-02 15:04:05",
		//DisableColors:    false,
		QuoteEmptyFields:       true,
		PadLevelText:           true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,

		//customize delimiter for logrus output

		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "servity",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "func",
		},
	})
}

func SetLogger() {

	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, "Set the log level for the application(info, debug, warn, error)")

	flag.Parse()

	switch *logLevel {
	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case model.LogLevelInfo:
		Logger.SetLevel(logrus.InfoLevel)
	case model.LogLevelWarning:
		Logger.SetLevel(logrus.WarnLevel)
	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	}
}

func Log(logLevel, packageLevel, functionName string, message, parameters interface{}) {
	switch logLevel {
	case model.LogLevelDebug:
		if parameters != nil {
			Logger.Debugf("packegeLevel: %s, functionName: %s, message: %s, parameters: %v", packageLevel, functionName, message, parameters)

		} else {
			Logger.Debugf("packegeLevel: %s, functionName: %s, message: %s", packageLevel, functionName, message)
		}

	// case model.LogLevelInfo:
	// 	if parameters != nil {
	// 		Logger.Infof("packegeLevel: %s, functionName: %s, message: %s, parameters: %v", packageLevel, functionName, message, parameters)
	// 	} else {
	// 		Logger.Infof("packegeLevel: %s, functionName: %s, message: %s", packageLevel, functionName, message)
	// 	}

	case model.LogLevelError:
		if parameters != nil {
			Logger.Errorf("packegeLevel: %s, functionName: %s, message: %v, parameters: %v", packageLevel, functionName, message, parameters)
		} else {
			Logger.Errorf("packegeLevel: %s, functionName: %s, message: %v", packageLevel, functionName, message)
		}

	case model.LogLevelWarning:
		if parameters != nil {
			Logger.Warnf("packegeLevel: %s, functionName: %s, message: %v, parameters: %v", packageLevel, functionName, message, parameters)
		} else {
			Logger.Warnf("packegeLevel: %s, functionName: %s, message: %v", packageLevel, functionName, message)
		}
	default:
		if parameters != nil {
			Logger.Infof("packegeLevel: %s, functionName: %s, message: %v, parameters: %v", packageLevel, functionName, message, parameters)
		} else {
			Logger.Infof("packegeLevel: %s, functionName: %s, message: %v", packageLevel, functionName, message)
		}
	}
}
