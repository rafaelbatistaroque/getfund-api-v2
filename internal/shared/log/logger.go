package shared_logger

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

// Logger provides a structured logger with different levels (Debug, Info, Warning, Error).
// It adds context and caller information to the log messages.
type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	context string
}

// New creates a new Logger instance with a given context.
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

// Debug logs a message at the DEBUG level.
func (l *Logger) Debug(v ...any) {
	l.debug.Printf(formatSimple, l.context, getCaller(), fmt.Sprint(v...))
}

// Info logs a message at the INFO level.
func (l *Logger) Info(v ...any) {
	l.info.Printf(formatSimple, l.context, getCaller(), fmt.Sprint(v...))
}

// Warn logs a message at the WARNING level.
func (l *Logger) Warn(v ...any) {
	l.warning.Printf(formatSimple, l.context, getCaller(), fmt.Sprint(v...))
}

// Error logs a message at the ERROR level.
func (l *Logger) Error(v ...any) {
	l.err.Printf(formatSimple, l.context, getCaller(), fmt.Sprint(v...))
}

// Debugf logs a formatted message at the DEBUG level.
func (l *Logger) Debugf(format string, v ...any) {
	l.debug.Printf(formatComplex+format, append([]any{l.context, getCaller()}, v...)...)
}

// Infof logs a formatted message at the INFO level.
func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(formatComplex+format, append([]any{l.context, getCaller()}, v...)...)
}

// Warnf logs a formatted message at the WARNING level.
func (l *Logger) Warnf(format string, v ...any) {
	l.warning.Printf(formatComplex+format, append([]any{l.context, getCaller()}, v...)...)
}

// Errorf logs a formatted message at the ERROR level.
func (l *Logger) Errorf(format string, v ...any) {
	l.err.Printf(formatComplex+format, append([]any{l.context, getCaller()}, v...)...)
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
