package recibo

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Niveles de error
var logLevels = map[string]zapcore.Level{"Debug": zapcore.DebugLevel, "Info": zapcore.InfoLevel, "Warn": zapcore.WarnLevel,
	"Error": zapcore.ErrorLevel, "Panic": zapcore.PanicLevel, "Fatal": zapcore.FatalLevel}

// Logger permite el registro de la actividad del sistema

type Logger struct {
	sugarLogger *zap.SugaredLogger
}

// NewLogger construye un objeto de tipo Logger a partir de la configuración del entorno
func NewLogger() Logger {
	logger := Logger{}
	config := NewConfig()

	// Establecemos medio de salida
	salida, err := os.OpenFile(config.LogFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Salida estándar asignada al Log debido a error al abrir archivo de configuración", err)
		salida = os.Stdout
	}
	writerSyncer := zapcore.AddSync(salida)

	// Configuramos formato de salida
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Establecemos nivel del logger
	level, existe := logLevels[config.LogLevel]
	if !existe {
		level = logLevels[defaultLogLevel]
	}

	// Construimos logger
	core := zapcore.NewCore(encoder, writerSyncer, level)
	logger.sugarLogger = zap.New(core).Sugar()
	logger.sugarLogger.Sync()

	return logger
}

// Debug registra un mensaje de log en nivel Debug
func (l *Logger) Debug(formato string, args ...interface{}) {
	l.sugarLogger.Debugf(formato, args...)
}

// Info registra un mensaje de log en nivel Info
func (l *Logger) Info(formato string, args ...interface{}) {
	l.sugarLogger.Infof(formato, args...)
}

// Warn registra un mensaje de log en nivel Warn
func (l *Logger) Warn(formato string, args ...interface{}) {
	l.sugarLogger.Warnf(formato, args...)
}

// Error registra un mensaje de log en nivel Error
func (l *Logger) Error(formato string, args ...interface{}) {
	l.sugarLogger.Errorf(formato, args...)
}

// Panic registra un mensaje de log en nivel Panic
func (l *Logger) Panic(formato string, args ...interface{}) {
	l.sugarLogger.Panicf(formato, args...)
}

// Fatal registra un mensaje de log en nivel Fatal
func (l *Logger) Fatal(formato string, args ...interface{}) {
	l.sugarLogger.Fatalf(formato, args...)
}
