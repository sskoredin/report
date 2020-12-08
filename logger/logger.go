package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	logfile = "app.log"
)

type Logger struct {
	*logrus.Entry
}

func New(name string, level logrus.Level) Logger {
	log := logrus.New()
	file := checkFile(name)
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

func (l Logger) ErrorWithCode(msg string, code int) {
	l.Errorf("status:[%d] error:'%v'", code, msg)
}

func checkFile(name string) string {
	p, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		return name
	}
	file := path.Join(p, logfile)
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		if _, err := os.Create(file); err != nil {
			logrus.Error(err)
		}
		logrus.Error(err)
		return logfile
	}

	return file
}
