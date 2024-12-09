package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

var (
	formatSimple  = "%s>%s: %v"
	formatComplex = "%s>%s: "
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	context string
}

func New(context string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, context, log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		debug:   log.New(writer, "[DEBUG]: ", logger.Flags()),
		info:    log.New(writer, "[INFO]: ", logger.Flags()),
		warning: log.New(writer, "[WARNING]: ", logger.Flags()),
		err:     log.New(writer, "[ERRO]: ", logger.Flags()),
		context: context,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Printf(formatSimple, l.context, getCaller(), fmt.Sprint(v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Printf(formatSimple, l.context, getCaller(), fmt.Sprint(v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Printf(formatSimple, l.context, getCaller(), fmt.Sprint(v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Printf(formatSimple, l.context, getCaller(), fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(formatComplex+format, append([]interface{}{l.context, getCaller()}, v...)...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(formatComplex+format, append([]interface{}{l.context, getCaller()}, v...)...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(formatComplex+format, append([]interface{}{l.context, getCaller()}, v...)...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(formatComplex+format, append([]interface{}{l.context, getCaller()}, v...)...)
}

func getCaller() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	fullName := fn.Name()
	// Extrai apenas o nome do método (remove o pacote)
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
}
