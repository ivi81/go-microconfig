//servicecfg.go - содержит структуры данных описывающие конфигурации интерфейсов взаимодействия микросервисов
package microconfig

//ServiceAPICfg тип данных хранящий информацию о конфигурации API-сервиса
type ServiceAPICfg struct {
	BasicServiceCfg `yaml:",inline"`
	ServerAPI       ServerAPICfg  `yaml:"serverAPI"`
	ClientAuth      ClientAuthCfg `yaml:"clientAuth"`
}

func (cfg *ServiceAPICfg) SetValuesFromEnv(envPrefix string) {
	envPref := JoinStr(envPrefix, ServiceAPIPrefix)
	cfg.BasicServiceCfg.SetValuesFromEnv(envPref)
	cfg.ServerAPI.SetValuesFromEnv(envPref)
	cfg.ClientAuth.SetValuesFromEnv(envPref)
}

//ServiceSTIXStorageCfg - тип данных хранящий информацию о конфигурации
// сервиса хранилища STIX-объектов
type ServiceSTIXStorageCfg struct {
	BasicServiceCfg   `yaml:",inline"`
	ClientSTIXStorage ClientSTIXStorageCfg `yaml:"clientSTIXStorage"`
}

func (cfg *ServiceSTIXStorageCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ServiceSTIXStoragePrefix)
	cfg.BasicServiceCfg.SetValuesFromEnv(envPref)
	cfg.ClientSTIXStorage.SetValuesFromEnv(envPref)
}

//ServiceAuthCfg - тип данных хранящий информацию о конфигурации
// сервиса авторизации
type ServiceAuthCfg struct {
	Logger            LoggerCfg            `yaml:"logger"`
	ServerAuth        ServerAuthCfg        `yaml:"serverAuth"`
	ClientAuthStorage ClientAuthStorageCfg `yaml:"clientAuthStorage"`
}

func (cfg *ServiceAuthCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ServiceAUTHrefix)
	cfg.Logger.SetValuesFromEnv(envPref)
	cfg.ClientAuthStorage.SetValuesFromEnv(envPref)
	cfg.ServerAuth.SetValuesFromEnv(envPref)
}
