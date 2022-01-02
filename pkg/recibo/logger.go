package recibo

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Única intancia de Logger del sistema
var myLogger *Logger

// Niveles de error
var logLevels = map[string]zapcore.Level{"Debug": zapcore.DebugLevel, "Info": zapcore.InfoLevel, "Warn": zapcore.WarnLevel,
	"Error": zapcore.ErrorLevel, "Panic": zapcore.PanicLevel, "Fatal": zapcore.FatalLevel}

// Logger permite el registro de la actividad del sistema

type Logger struct {
	sugarLogger *zap.SugaredLogger
}

// GetLogger devuelve la única instancia de tipo Logger del sistema. Si aún no ha sido inicializada,
// se contruye a partir de la configuración del entorno
func GetLogger() *Logger {

	if myLogger != nil {
		return myLogger
	}

	logger := Logger{}
	config := NewConfig()

	// Establecemos medio de salida
	salida, err := os.OpenFile(config.LogFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Se ha asignado al Log la salida estándar debido a un error abriendo archivo de Log", err)
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

	myLogger = &logger

	return myLogger
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
