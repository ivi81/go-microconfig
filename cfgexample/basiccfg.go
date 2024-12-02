// basiccfg.go - содержит структуры данных общие для конфигураций всех микросервисов структуры
package cfgexample

// BasicCfg описывает параметры конфигурации общие для всех микросервисов
type BasicCfg struct {
	Host string `yaml:"host" env:"HOST"`
	Port int    `yaml:"port" env:"PORT"`
}

// BasicClientCfg описывает общие параметры конфигурации
// клиентской части всех микросервисов.
type BasicClientCfg struct {
	BasicCfg `yaml:",inline"`
}

// BasicServerCfg описывает общие параметры конфигурации
// серверной части всех микросервисов.
type BasicServerCfg struct {
	BasicCfg `yaml:",inline" env:"_"`
}

// BasicStorageClientCfg описывает общие параметры конфигурации
// для работы с хранилищем данных
type BasicStorageCfg struct {
	User string `yaml:"user" env:"DB_USER"`
	Pwd  string `yaml:"pwd" env:"DB_PWD"`
	Db   string `yaml:"db" env:"DB_NAME"`
}
