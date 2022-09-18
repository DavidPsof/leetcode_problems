package log

import (
	"github.com/DavidPsof/leetcode_problems/backend/config"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func InitLog(conf config.LogSettings) {
	var writers []io.Writer

	if conf.LogInConsole {
		writers = append(writers, os.Stdout)
	}

	if conf.LogInFile {
		file, err := os.OpenFile(conf.LogFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			file, err = os.Create(conf.LogFileName)
			if err != nil {
				log.Fatal(err)
			}
		}

		writers = append(writers, file)
	}

	mw := io.MultiWriter(writers...)

	log.SetOutput(mw)

	log.SetReportCaller(conf.LogReportCaller)

	log.SetFormatter(&log.JSONFormatter{})

	switch conf.LogLevel {
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "PANIC":
		log.SetLevel(log.PanicLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	log.Info("Log initiated")
}
