//simpleserver.go - содержит структуры данных описывающие конфигурации элементарных серверов
package microconfig

//WssServerCfg описыввает параметры конфигурации WSS-сервера
type ServerWssCfg struct {
	BasicServerCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ServerWssCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ServerWSSPrefix)

	cfg.BasicServerCfg.SetValuesFromEnv(envPref)
}

//TaxiServerCfg описыввает параметры конфигурации TAXI-сервера
type ServerTaxiCfg struct {
	BasicServerCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ServerTaxiCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ServerTAXIPrefix)

	cfg.BasicServerCfg.SetValuesFromEnv(envPref)
}

//HTTPServerCfg описыввает параметры конфигурации HTTP-сервера
type ServerHttpCfg struct {
	BasicServerCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ServerHttpCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ServerHTTPPrefix)

	cfg.BasicServerCfg.SetValuesFromEnv(envPref)
}

//HTTPServerCfg описыввает параметры конфигурации gRPC-сервера
type ServerGrpsCfg struct {
	BasicServerCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ServerGrpsCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ServerGRPCPrefix)

	cfg.BasicServerCfg.SetValuesFromEnv(envPref)
}
