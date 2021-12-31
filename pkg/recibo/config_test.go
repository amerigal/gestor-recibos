package recibo

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	testLogFile := "./test-config.log"
	testLogLevel := "Panic"

	logPathEnvValue, logPathEnvExists := os.LookupEnv("LOGFILEPATH")
	logLevelEnvValue, logLevelEnvExists := os.LookupEnv("LOGLEVEL")
	os.Setenv("LOGFILEPATH", testLogFile)
	os.Setenv("LOGLEVEL", testLogLevel)

	config := NewConfig()

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

	if config.LogFilePath != testLogFile || config.LogLevel != testLogLevel {

		t.Fatalf("Config no toma valores del entorno correctamente.")
	}
}
