package util

import (
	"fmt"
	"log"
	"os"
)

type YLogger struct {
	name        string
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func NewYLogger(name string) *YLogger {
	yLogger := new(YLogger)
	yLogger.name = name
	yLogger.infoLogger = initInnerLogger("INFO")
	yLogger.warnLogger = initInnerLogger("WARN")
	yLogger.errorLogger = initInnerLogger("ERR")
	return yLogger
}
func (it *YLogger) Info(v interface{}) {
	_ = it.infoLogger.Output(2, fmt.Sprintln(v))
}
func (it *YLogger) Warn(v interface{}) {
	_ = it.warnLogger.Output(2, fmt.Sprintln(v))
}
func (it *YLogger) Err(v interface{}) {
	_ = it.errorLogger.Output(2, fmt.Sprintln(v))
}
func (it *YLogger) ErrOrNil(err error) error {
	if err != nil {
		_ = it.errorLogger.Output(2, fmt.Sprintln(err))
		return err
	}
	return nil
}

func initInnerLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix+"\t", log.LstdFlags|log.Lshortfile)
}
