package recibo

import (
	"os"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	testLogFile := "/tmp/test.log"

	logPathEnvValue, logPathEnvExists := os.LookupEnv("LOGFILEPATH")
	logLevelEnvValue, logLevelEnvExists := os.LookupEnv("LOGLEVEL")
	os.Setenv("LOGFILEPATH", testLogFile)
	os.Setenv("LOGLEVEL", "Debug")

	logger := NewLogger()

	logger.Debug("Prueba DEBUG")
	logger.Info("Prueba INFO")
	logger.Warn("Prueba WARN")
	logger.Error("Prueba ERROR")

	if logPathEnvExists {
		os.Setenv("LOGFILEPATH", logPathEnvValue)
	} else {
		os.Unsetenv("LOGFILEPATH")
	}
	if logLevelEnvExists {
		os.Setenv("LOGLEVEL", logLevelEnvValue)
	} else {
		os.Unsetenv("LOGLEVEL")
	}

	file, err := os.Open(testLogFile)
	if err != nil {
		t.Fatalf("Error al abrir fichero de Log: %s", err)
	}
	data := make([]byte, 1000)
	count, err := file.Read(data)
	os.Remove(testLogFile)
	if err != nil {
		t.Fatalf("Error al leer fichero de Log: %s", err)
	}

	contenido := string(data[:count])

	if !strings.Contains(contenido, "Prueba DEBUG") ||
		!strings.Contains(contenido, "Prueba INFO") ||
		!strings.Contains(contenido, "Prueba WARN") ||
		!strings.Contains(contenido, "Prueba ERROR") {

		t.Fatalf("Logger no imprime mensajes de Log correctamente.")
	}
}
