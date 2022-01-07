package recibo

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	logPath, logPathExists := os.LookupEnv("LOGFILEPATH")

	if !logPathExists {
		t.Fatalf("Variable de entorno LOGFILEPATH no tiene valor asignado")
	}

	testMsg := "Test log -> " + time.Now().String()

	GetLogger().Debug(testMsg)

	data, err := ioutil.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Error al leer archivo de log: %s", err)
	}

	contenido := string(data)

	if !strings.Contains(contenido, testMsg) {
		t.Fatalf("Logger no imprime mensajes de Log correctamente.")
	}
}
