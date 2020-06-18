package logs

import (
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"yangjian.com/basicarchitecture/utils"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

// errLogger is init logger error
var errLogger error
var onceLogger sync.Once

func init() {
	onceLogger.Do(func() {
		logDir := "./storage/logs"
		if utils.Exists(logDir) {
			err := os.MkdirAll(logDir, os.ModePerm)
			if err != nil {
				errLogger = err
			}
		}
		logPath := filepath.Join(logDir, "log")
		log.SetReportCaller(true)
		log.SetLevel(log.InfoLevel)
		log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
		writer, err := rotatelogs.New(
			logPath+".%Y%m%d%H%M%S",
			rotatelogs.WithLinkName(logPath),
			rotatelogs.WithRotationTime(time.Hour*24),
			rotatelogs.WithMaxAge(time.Hour*24*30),
		)

		// lfsHook := lfshook.NewHook(lfshook.WriterMap{
		// 	log.DebugLevel: writer,
		// 	log.InfoLevel:  writer,
		// 	log.WarnLevel:  writer,
		// 	log.ErrorLevel: writer,
		// 	log.FatalLevel: writer,
		// 	log.PanicLevel: writer,
		// }, &log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
		// instance.AddHook(lfsHook)

		if err != nil {
			errLogger = err
		}
		log.SetOutput(io.MultiWriter(writer, os.Stdout))
	})
}

// Error is get logger error
func Error() error {

	return errLogger
}
