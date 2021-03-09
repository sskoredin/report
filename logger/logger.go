package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	logfile = "./app.log"
)

type Logger struct {
	*logrus.Entry
}

func New(name string, level logrus.Level) Logger {
	log := logrus.New()
	file := checkFile()
	log.SetLevel(level)
	log.SetFormatter(&nested.Formatter{
		HideKeys: true,
	})
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Error(err)
	}
	log.SetOutput(f)

	return Logger{log.WithField("method", name)}
}

func checkFile() string {
	_, err := os.Stat(logfile)
	if os.IsNotExist(err) {
		if _, err := os.Create(logfile); err != nil {
			logrus.Error(err)
		}
		logrus.Error(err)
		return logfile
	}

	return logfile
}
