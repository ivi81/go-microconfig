// clientcfg.go - содержит структуры данных описывающие конфигурацию клиентов взаимодействия с микросервисами
package cfgexample

import "time"

// ClientAuthStorageCfg - тип данных хранящий информацию о конфигурации
// клиентских настроек сервиса авторизации для подключения к хранилищу информации
type ClientAuthStorageCfg struct {
	BasicClientCfg  `yaml:",inline" env:"_"`
	BasicStorageCfg `yaml:",inline" env:"_"`
}

// ClientSTIXStorageCfg хранит конфигурацию подключения к хранилищу STIX-объектов
type ClientSTIXStorageCfg struct {
	BasicClientCfg  `yaml:",inline" env:"_"`
	BasicStorageCfg `yaml:",inline" env:"_"`
}

// ClientNatsCfg - тип данных хранящий конфигурацию подключения
// к брокеру сообщений NATS
type ClientNatsCfg struct {
	BasicClientCfg   `yaml:",inline" env:"_"`
	TurnOffEcho      bool          `yaml:"turnOffEcho" env:"TURN_OF_ECHO"`
	PingInterval     time.Duration `yaml:"pingInterval" env:"PING_INTERVAL"`
	AuthWithToken    bool          `yaml:"authWithToken" env:"AUTH_WITH_TOKEN"`
	AuthWithUser     bool          `yaml:"authWithUser" env:"AUTH_WITH_USER"`
	AuthWithCredFile bool          `yaml:"authWithCredFile" env:"AUTH_WITH_CRED_FILE"`
	TLSOn            bool          `yaml:"tlsOn" env:"TLS_ON"`
}

// ClientAuthCfg описывает конфигурацию параметров подключения
// к сервису авторизации
type ClientAuthCfg struct {
	BasicClientCfg `yaml:",inline" env:"_"`
}

// ClientLogsStorage описывает параметры подключения к хранилищу логов
type ClientLogsStorageCfg struct {
	BasicClientCfg `yaml:",inline" env:"_"`
}

// ClientHttpCfg описывает конфигурацию параметров подключения
// к HTTP сервреру
type ClientHttpCfg struct {
	BasicClientCfg `yaml:",inline" env:"_"`
	User           string        `yaml:"user" env:"USER"`
	Pwd            string        `yaml:"pwd" env:"PWD"`
	ApiKey         string        `yaml:"apiKey" env:"API_KEY"`
	BearerToken    string        `yaml:"bearerToken" env:"BEARER_TOKEN"`
	Timeout        time.Duration `yaml:"timeout" env:"TIMEOUT"`
	UserAgent      string        `yaml:"userAgent" env:"USER_AGENT"`
	SleepTime      time.Duration `yaml:"sleepTime" env:"SLEEP_TIME"`
	ResponseLimit  int           `yaml:"responseLimit" env:"RESPONSE_LIMIT"`
}
