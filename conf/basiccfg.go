//basiccfg.go - содержит структуры данных общие для конфигураций всех микросервисов структуры
package conf

//BasicCfg описывает параметры конфигурации общие для всех микросервисов
type BasicCfg struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения
func (cfg *BasicCfg) SetValuesFromEnv(envPrefix string) {

	tt := JoinStr(envPrefix, "HOST")
	if host, ok := LookupEnv(tt); ok {
		cfg.Host = host
	}
	if port, ok := LookupEnvAsInt(JoinStr(envPrefix, "PORT")); ok {
		cfg.Port = port
	}
}

//HostsAsString формирует строку вида "Host[0]:Port,Host[1]:Port,..." из содержимого полей Hosts и Port
/*func (cfg *BasicCfg) HostsAsString() string {

	return strings.Join(cfg.HostList(), strSplitter)
}*/

//HostList формирует список строк вида ["Host[0]:Port","Host[1]:Port"]
/*func (cfg *BasicCfg) HostList() []string {

	var hosts []string
	portAsStr := strconv.Itoa(cfg.Port)

	for _, v := range cfg.Host {
		hosts = append(hosts, strings.Join([]string{v, portAsStr}, hostPortSplitter))
	}

	return hosts
}*/

//BasicClientCfg описывает общие параметры конфигурации
//клиентской части всех микросервисов.
type BasicClientCfg struct {
	BasicCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения
func (cfg *BasicClientCfg) SetValuesFromEnv(envPrefix string) {

	cfg.BasicCfg.SetValuesFromEnv(envPrefix)
}

//BasicServerCfg описывает общие параметры конфигурации
//серверной части всех микросервисов.
type BasicServerCfg struct {
	BasicCfg `yaml:",inline"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения
func (cfg *BasicServerCfg) SetValuesFromEnv(envPrefix string) {

	cfg.BasicCfg.SetValuesFromEnv(envPrefix)
}

//BasicStorageClientCfg описывает общие параметры конфигурации
//для работы с хранилищем данных
type BasicStorageCfg struct {
	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`
	Db   string `yaml:"db"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения
func (cfg *BasicStorageCfg) SetValuesFromEnv(envPrefix string) {

	if user, ok := LookupEnv(JoinStr(envPrefix, "DB_USER")); ok {
		cfg.User = user
	}
	if pwd, ok := LookupEnv(JoinStr(envPrefix, "DB_PWD")); ok {
		cfg.Pwd = pwd
	}
	if dbName, ok := LookupEnv(JoinStr(envPrefix, "DB_NAME")); ok {
		cfg.Db = dbName
	}
}
