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

	"github.com/alswl/go-toodledo/pkg/version"

	"github.com/sirupsen/logrus"
)

var (
	// formatter is the default log formatter.
	formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			repoPath := os.Getenv("GOPATH") + "/src/" + version.Package
			filename := strings.ReplaceAll(frame.File, repoPath, "")
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
	}
	factory *loggerFactory
	// once is used to make sure the factory is initialized only once.
	once sync.Once
	// silentLogger is a logger that does not output anything.
	silentLogger = logrus.New()
	// stdoutLogger is a logger that outputs to stdout
	// should NOT use in TUI mode.
	stdoutLogger = logrus.StandardLogger()
)

// loggerFactory is the logger factory, it holds the log root path and logger map.
type loggerFactory struct {
	logRoot   string
	loggerMap *sync.Map
	isStd     bool
	isSilence bool
}

func NewFileLogger(path string) (logrus.FieldLogger, error) {
	dir, _ := filepath.Split(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		ierr := os.MkdirAll(path, os.ModePerm)
		if ierr != nil {
			return nil, ierr
		}
	}

	log := logrus.New()
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	log.Out = f
	logrus.SetFormatter(formatter)
	return log, nil
}

// GetLoggerOrCreate returns a logger with the given name.
// If the logger is not initialized, it will try to initialize it.
// If the logger initialization fails, it will return nil.
func GetLoggerOrCreate(name string) logrus.FieldLogger {
	if factory == nil {
		logrus.Warn("logger factory is not initialized")
		return nil
	}
	if factory.isSilence {
		return silentLogger
	}
	if factory.isStd {
		return stdoutLogger
	}

	newPath := path.Join(factory.logRoot, name+".log")
	dir, file := filepath.Split(newPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		ierr := os.MkdirAll(newPath, os.ModePerm)
		if ierr != nil {
			logrus.WithField("dir", dir).WithField("file", file).WithError(err).Error("create log dir")
			return nil
		}
	}
	logger, err := NewFileLogger(newPath)
	if err != nil {
		logrus.WithField("dir", dir).WithField("file", file).WithError(err).Error("create log file")
		return nil
	}
	return logger
}

// GetLoggerOrDefault returns a logger with the given name.
// If the logger is not initialized, it will return a default logger.
func GetLoggerOrDefault(name string) logrus.FieldLogger {
	logger := GetLoggerOrCreate(name)
	if logger == nil {
		return stdoutLogger
	}
	return logger
}

func ProvideLogger() logrus.FieldLogger {
	return GetLoggerOrDefault("default")
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
