package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var formatter = &logrus.TextFormatter{
	CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
		repoPath := os.Getenv("GOPATH") + "/src/github.com/alswl/go-toodledo"
		filename := strings.ReplaceAll(frame.File, repoPath, "")
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

func NewLogger() logrus.FieldLogger {
	return nil
}

func NewFileLogger(path string) logrus.FieldLogger {
	dir, file := filepath.Split(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		ierr := os.MkdirAll(path, os.ModePerm)
		if ierr != nil {
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

func GetLogger(name string) logrus.FieldLogger {
	if factory == nil {
		logrus.Warn("logger factory is not initialized")
		return logrus.StandardLogger()
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
		ierr := os.MkdirAll(path, os.ModePerm)
		if ierr != nil {
			log := logrus.New()
			log.WithField("dir", dir).WithField("file", file).WithError(err).Error("create log dir")
			return log
		}
	}
	return NewFileLogger(path)
}

func ProvideLogger() logrus.FieldLogger {
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
