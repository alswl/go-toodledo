package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var formatter = &logrus.TextFormatter{
	CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
		repoPath := os.Getenv("GOPATH") + "/src/github.com/alswl/go-toodledo"
		filename := strings.Replace(frame.File, repoPath, "", -1)
		return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
	},
	TimestampFormat: time.RFC3339,
	FullTimestamp:   true,
}

var factory *loggerFactory
var once sync.Once
var silentLogger = logrus.New()
var stdoutLogger = logrus.StandardLogger()

type loggerFactory struct {
	logRoot   string
	loggerMap *sync.Map
	isStd     bool
	isSilence bool
}

func NewLogger() *logrus.Logger {
	return nil
}

func NewFileLogger(path string) *logrus.Logger {
	dir, file := filepath.Split(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log := logrus.New()
			log.WithField("dir", dir).WithField("file", file).WithError(err).Error("create log dir")
			return log
		}
	}

	log := logrus.New()
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.WithField("path", path).WithError(err).Error("open log file")
	}
	log.Out = f
	logrus.SetFormatter(formatter)
	return log
}

func GetLogger(name string) *logrus.Logger {
	if factory == nil {
		logrus.Error("logger factory is not initialized")
		return nil
	}
	if factory.isSilence {
		return silentLogger
	}
	if factory.isStd {
		return stdoutLogger
	}
	path := path.Join(factory.logRoot, name+".log")
	dir, file := filepath.Split(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log := logrus.New()
			log.WithField("dir", dir).WithField("file", file).WithError(err).Error("create log dir")
			return log
		}
	}
	return NewFileLogger(path)
}

func ProvideLogger() *logrus.Logger {
	return GetLogger("default")
}

func ProvideLoggerItf() logrus.FieldLogger {
	return GetLogger("default")
}

func InitFactory(logRoot string, isStd bool, isSilence bool) error {
	var err error
	once.Do(func() {
		err = os.MkdirAll(logRoot, os.ModePerm)
		if err != nil {
			return
		}

		factory = &loggerFactory{
			logRoot:   logRoot,
			loggerMap: &sync.Map{},
			isStd:     isStd,
			isSilence: isSilence,
		}
	})
	return err
}

func init() {
	silentLogger.Out = io.Discard
	stdoutLogger.SetFormatter(formatter)
}
