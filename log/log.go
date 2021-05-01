package log

import (
	"github.com/pengcheng789/dumpus/util"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var logFilePath string

var Logger = log.New()

func init() {
	Logger.Out = os.Stdout
	Logger.Level = log.DebugLevel
	Logger.Hooks.Add(&FileHock{})

	runtimeDir, err := util.RuntimeDir()
	if err != nil {
		Logger.Fatal(errors.Wrap(err, "Get runtime path failure."))
	}

	logFileParent := "log"
	if err := os.MkdirAll(filepath.Join(runtimeDir, logFileParent), os.ModePerm); err != nil {
		Logger.Fatal(errors.Wrap(err, "Make log directory failure."))
	}

	logFilename := "dumpus.log"
	logFilePath = filepath.Join(runtimeDir, logFileParent, logFilename)
}

type FileHock struct {}

func (h FileHock) Levels() []log.Level {
	return []log.Level {
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
		log.WarnLevel,
		log.InfoLevel,
	}
}

func (h FileHock) Fire(entry *log.Entry) error {
	msg, err := entry.String()
	if err != nil {
		return errors.Wrap(err, "Get error message failure.")
	}

	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrap(err, "Open the log file failure.")
	}

	_, err = logFile.WriteString(msg)
	if err != nil {
		return errors.Wrap(err, "Write log failure.")
	}

	return nil
}
