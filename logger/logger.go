package logger

import (
	"fmt"
	"github.com/sskoredin/telegram_client"
	"log"
)

type Logger struct {
	*telegram_client.Client
	name string
}

func New(name string) *Logger {
	c, err := telegram_client.NewClient()
	if err != nil {
		log.Println(err)
	}

	return &Logger{
		Client: c,
		name:   name,
	}
}
func (l Logger) Info(msg ...interface{}) error {
	log.Print()
	return l.Send(l.convert("Info", msg))
}

func (l Logger) Infof(format string, v ...interface{}) error {
	return l.Send(l.convert("Info", fmt.Sprintf(format, v...)))
}

func (l Logger) Warn(msg ...interface{}) error {
	return l.Send(l.convert("Warn", msg))
}

func (l Logger) Warnf(format string, v ...interface{}) error {
	return l.Send(l.convert("Debug", fmt.Sprintf(format, v...)))
}

func (l Logger) Debug(msg ...interface{}) error {
	return l.Send(l.convert("Debug", msg))
}

func (l Logger) Debugf(format string, v ...interface{}) error {
	return l.Send(l.convert("Debug", fmt.Sprintf(format, v...)))
}

func (l Logger) Fatal(msg ...interface{}) error {
	return l.Send(l.convert("Fatal", msg))
}

func (l Logger) Fatalf(format string, v ...interface{}) error {
	return l.Send(l.convert("Fatal", fmt.Sprintf(format, v...)))
}

func (l Logger) Error(msg ...interface{}) error {
	return l.Send(l.convert("Error", msg))
}

func (l Logger) convert(level string, msg ...interface{}) string {
	return fmt.Sprintf("iiko_olap_report\n[%s][%s]\n%v", l.name, level, msg)
}
