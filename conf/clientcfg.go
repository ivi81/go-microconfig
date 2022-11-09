//clientcfg.go - содержит структуры данных описывающие конфигурацию клиентов взаимодействия с микросервисами
package conf

import "time"

//ClientAuthStorageCfg - тип данных хранящий информацию о конфигурации
// клиентских настроек сервиса авторизации для подключения к хранилищу информации
type ClientAuthStorageCfg struct {
	BasicClientCfg  `yaml:",inline"`
	BasicStorageCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ClientAuthStorageCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ClientAuthStoragePrefix)

	cfg.BasicClientCfg.SetValuesFromEnv(envPref)
	cfg.BasicStorageCfg.SetValuesFromEnv(envPref)
}

//ClientSTIXStorageCfg хранит конфигурацию подключения к хранилищу STIX-объектов
type ClientSTIXStorageCfg struct {
	BasicClientCfg  `yaml:",inline"`
	BasicStorageCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ClientSTIXStorageCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ClientSTIXStoragePrefix)

	cfg.BasicClientCfg.SetValuesFromEnv(envPref)
	cfg.BasicStorageCfg.SetValuesFromEnv(envPref)
}

//ClientNatsCfg - тип данных хранящий конфигурацию подключения
//к брокеру сообщений NATS
type ClientNatsCfg struct {
	BasicClientCfg   `yaml:",inline"`
	TurnOffEcho      bool          `yaml:"turnOffEcho"`
	PingInterval     time.Duration `yaml:"pingInterval"`
	AuthWithToken    bool          `yaml:"authWithToken"`
	AuthWithUser     bool          `yaml:"authWithUser"`
	AuthWithCredFile bool          `yaml:"authWithCredFile"`
	TLSOn            bool          `yaml:"tlsOn"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ClientNatsCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ClientNatsPrefix)

	cfg.BasicClientCfg.SetValuesFromEnv(envPref)

	if turnOfEcho, ok := LookupEnvAsBool(JoinStr(envPref, "TURN_OF_ECHO")); ok {
		cfg.TurnOffEcho = turnOfEcho
	}
	if pingInterval, ok := LookupEnvAsTime(JoinStr(envPref, "PING_INTERVAL")); ok {
		cfg.PingInterval = pingInterval
	}
	if authWithToken, ok := LookupEnvAsBool(JoinStr(envPref, "AUTH_WITH_TOKEN")); ok {
		cfg.AuthWithToken = authWithToken
	}
	if authWithUser, ok := LookupEnvAsBool(JoinStr(envPref, "AUTH_WITH_USER")); ok {
		cfg.AuthWithUser = authWithUser
	}
	if authWithCredFile, ok := LookupEnvAsBool(JoinStr(envPref, "AUTH_WITH_CRED_FILE")); ok {
		cfg.AuthWithCredFile = authWithCredFile
	}
	if tlsOn, ok := LookupEnvAsBool(JoinStr(envPref, "TLS_ON")); ok {
		cfg.TLSOn = tlsOn
	}
}

//ClientAuthCfg описывает конфигурацию параметров подключения
//к сервису авторизации
type ClientAuthCfg struct {
	BasicClientCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ClientAuthCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ClientAuthPrefix)

	cfg.BasicClientCfg.SetValuesFromEnv(envPref)
}

//ClientLogsStorage описывает параметры подключения к хранилищу логов
type ClientLogsStorageCfg struct {
	BasicClientCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ClientLogsStorageCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ClientLogsStoragePrefix)
	cfg.BasicClientCfg.SetValuesFromEnv(envPref)
}

//ClientHttpCfg описывает конфигурацию параметров подключения
//к HTTP сервреру
type ClientHttpCfg struct {
	BasicClientCfg `yaml:",inline"`
	User           string        `yaml:"user"`
	Pwd            string        `yaml:"pwd"`
	ApiKey         string        `yaml:"apiKey"`
	BearerToken    string        `yaml:"bearerToken"`
	Timeout        time.Duration `yaml:"timeout"`
	UserAgent      string        `yaml:"userAgent"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ClientHttpCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ClientHttpPrefix)
	cfg.BasicClientCfg.SetValuesFromEnv(envPref)

	if user, ok := LookupEnv(JoinStr(envPref, "USER")); ok {
		cfg.User = user
	}

	if pwd, ok := LookupEnv(JoinStr(envPref, "PWD")); ok {
		cfg.Pwd = pwd
	}

	if apiKey, ok := LookupEnv(JoinStr(envPref, "API_KEY")); ok {
		cfg.ApiKey = apiKey
	}

	if bearierToken, ok := LookupEnv(JoinStr(envPref, "BEARER_TOKEN")); ok {
		cfg.BearerToken = bearierToken
	}

	if timeout, ok := LookupEnvAsTime(JoinStr(envPref, "TIMEOUT")); ok {
		cfg.Timeout = timeout
	}

	if userAgent, ok := LookupEnv(JoinStr(envPref, "USER_AGENT")); ok {
		cfg.UserAgent = userAgent
	}
}
