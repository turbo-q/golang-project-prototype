package logger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger_ *zap.SugaredLogger

const delimiter = "                    "

func init() {
	// 日志级别
	logLevel := "DEBUG"

	atomicLevel := zap.NewAtomicLevel()
	switch logLevel {
	case "DEBUG":
		atomicLevel.SetLevel(zapcore.DebugLevel)
	case "INFO":
		atomicLevel.SetLevel(zapcore.InfoLevel)
	case "WARN":
		atomicLevel.SetLevel(zapcore.WarnLevel)
	case "ERROR":
		atomicLevel.SetLevel(zapcore.ErrorLevel)
	case "DPANIC":
		atomicLevel.SetLevel(zapcore.DPanicLevel)
	case "PANIC":
		atomicLevel.SetLevel(zapcore.PanicLevel)
	case "FATAL":
		atomicLevel.SetLevel(zapcore.FatalLevel)
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		atomicLevel,
	)
	logger := zap.New(zapCore, zap.AddCaller())
	defer logger.Sync()

	logger_ = logger.Sugar()
}

func Info(args ...interface{}) {
	logger_.Info(args...)
	fmt.Println(delimiter)
}

func Infof(format string, args ...interface{}) {
	logger_.Infof(format, args...)
	fmt.Println(delimiter)
}

func Debug(args ...interface{}) {
	logger_.Debug(args...)
	fmt.Println(delimiter)
}

func Debugf(format string, args ...interface{}) {
	logger_.Debugf(format, args...)
	fmt.Println(delimiter)
}

func Error(desc string, err error) {
	logger_.Errorw(desc, "error", err)
	fmt.Println(delimiter)
}

func Errorf(format string, err error, args ...interface{}) {
	desc := fmt.Sprintf(format, args...)
	logger_.Errorw(desc, "error", err)
	fmt.Println(delimiter)
}

// http request
func Request(desc, url, method string, values url.Values) {
	if method == http.MethodGet {
		url += "?" + values.Encode()
	}
	logger_.Infow(desc, "url", url, "method", method, "values", values)
	fmt.Println(delimiter)
}

// http response
func Response(desc string, response interface{}) {
	logger_.Infow(desc, "response", response)
	fmt.Println(delimiter)
}

// debug file，如果日志太多，使用文件单独查看比较方便
func DebugFile(path string, data interface{}) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		Error("创建文件失败", err)
		return
	}
	jb, _ := json.Marshal(&data)
	f.Write(jb)
	f.Close()
}
