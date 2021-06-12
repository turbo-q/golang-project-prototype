package logger

import (
	"fmt"
	"golang-project-prototype/config"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

const delimiter = "                    "

func init() {
	log = logrus.New()
	if config.DefaultConfig.Env == "prod" {
		log.Formatter = new(logrus.JSONFormatter)
	} else {
		log.Formatter = new(logrus.TextFormatter)
	}
	log.Level = logrus.InfoLevel
	log.Out = os.Stdout
}

func map2LogFields(m map[string]interface{}) logrus.Fields {
	return logrus.Fields(m)
}

func Info(msg ...interface{}) {
	log.Info(msg...)
	fmt.Println(delimiter)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
	fmt.Println(delimiter)
}

func Infom(msg string, m map[string]interface{}) {
	log.WithFields(map2LogFields(m)).Info(msg)
	fmt.Println(delimiter)
}

func Warn(msg string) {
	log.Warn(msg)
	fmt.Println(delimiter)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
	fmt.Println(delimiter)
}

func Error(msg string, err error) {
	m := map[string]interface{}{
		"error": err,
	}
	log.WithFields(map2LogFields(m)).Error(msg)
	fmt.Println(delimiter)
}

func Errorf(format string, err error, f ...interface{}) {
	m := map[string]interface{}{"error": err}
	log.WithFields(map2LogFields(m)).Errorf(format, f...)
	fmt.Println(delimiter)
}

//	http请求
func Request(msg, method, url string, values url.Values) {
	m := map[string]interface{}{
		"method": method,
		"url":    url,
		"values": values.Encode(),
	}
	log.WithFields(map2LogFields(m)).Info(msg)
	fmt.Println(delimiter)
}

// http 响应
func Response(msg string, resp interface{}) {
	m := map[string]interface{}{
		"response": resp,
	}
	log.WithFields(map2LogFields(m)).Info(msg)
	fmt.Println(delimiter)
}
