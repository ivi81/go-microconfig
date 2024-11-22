package microconfig

//BasicServiceCfg описывает базовую (общую для всех) конфигурацию сервиса
/*type BasicServiceCfg struct {
	DataExchangeClient ClientNatsCfg `yaml:"clientNats"`
	Logger             LoggerCfg     `yam:"logger"`
}
*/

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
/*func (cfg *BasicServiceCfg) SetValuesFromEnv(envPrefix string) {

	cfg.DataExchangeClient.SetValuesFromEnv(envPrefix)
	cfg.Logger.SetValuesFromEnv(envPrefix)
}
*/