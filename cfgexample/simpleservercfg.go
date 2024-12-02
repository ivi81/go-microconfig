// simpleserver.go - содержит структуры данных описывающие конфигурации элементарных серверов
package cfgexample

// WssServerCfg описыввает параметры конфигурации WSS-сервера
type ServerWssCfg struct {
	BasicServerCfg `yaml:",inline"`
}
