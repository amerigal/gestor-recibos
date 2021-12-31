package recibo

import (
	"fmt"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// Clave para el path del archivo de Log
const logPathKey = "logFilePath"

// Clave para el nivel del archivo de Log
const logLevelKey = "logLevel"

// Clave para establecer el uso de etcd
const etcdKey = "useEtcd"

// Clave para la dirección de etcd
const etcdAddressKey = "etcdAddress"

// Clave para el archivo config de etcd
const etcdConfigPathKey = "etcdConfigPathKey"

// Path del archivo de Log por defecto
const defaultLogFilePath = "/tmp/gestor-gastos.log"

// Nivel del Log por defecto
const defaultLogLevel = "Debug"

// Uso de etcd por defecto
const defaultEtcd = "false"

// Directorio del archivo de configuración
const configPath = "../.."

// Nombre del archivo de configuración
const configName = "config"

// Extensión del archivo de configuración
const configType = "yml"

// Config encapsula la configuración del sistema

type Config struct {
	// Path del archivo de Log
	LogFilePath string

	// Nivel del Log
	LogLevel string
}

// NewConfig construye un objeto Config tomando valores de múltiples fuentes
func NewConfig() Config {
	config := Config{}

	// Establecemos valores por defecto
	viper.SetDefault(logPathKey, defaultLogFilePath)
	viper.SetDefault(logLevelKey, defaultLogLevel)
	viper.SetDefault(etcdKey, defaultEtcd)

	// Tomamos valores de archivo de configuración
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No se ha cargado ningún archivo de configuración:", err)
	}

	// Tomamos valores de variables de entorno
	viper.AutomaticEnv()

	// Tomamos valores de etcd
	if viper.GetBool(etcdKey) {
		viper.AddRemoteProvider("etcd", viper.GetString(etcdAddressKey), viper.GetString(etcdConfigPathKey))
		viper.SetConfigType("json")
		err := viper.ReadRemoteConfig()
		if err != nil {
			fmt.Println("Error al tomar configuración de etcd:", err)
		}
	}

	// Cargamos valores en struct config
	viper.Unmarshal(&config)

	return config
}
