//loggercfg.go - содержит структуры данных описывающие конфигурацю логгера
package microconfig

//LoggerCfg -описывает параметры конфигурации логировния
//Поле Mode задает режим работы логера и должно иметь значение из следующего списка:
//- std - вывод сообшений на stdout,
//- file - запись сообщений в файл,
//- service - отправка сообщений удаленной службе логирования
type LoggerCfg struct {
	Mode             []string             `yaml:"mode"`
	Path             string               `yaml:"path"`
	LogService       ClientLogsStorageCfg `yaml:"logService"`
	LogFormat        string               `yaml:"logFormat"`
	LogLevel         []string             `yaml:"logLevel"`
	DisableTimeStamp bool                 `yaml:"disableTimeStamp"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *LoggerCfg) SetValuesFromEnv(envPrefix string) {

	if mode, ok := LookupEnvAsSlice(JoinStr(envPrefix, "LOG_MODE"), strSplitter); ok {
		cfg.Mode = mode
	}
	if path, ok := LookupEnv(JoinStr(envPrefix, "LOG_PATH")); ok {
		cfg.Path = path
	}

	if format, ok := LookupEnv(JoinStr(envPrefix, "LOG_FORMAT")); ok {
		cfg.Path = format
	}

	if level, ok := LookupEnvAsSlice(JoinStr(envPrefix, "LOG_LEVEL"), strSplitter); ok {
		cfg.LogLevel = level
	}

	if disableTimeStamp, ok := LookupEnvAsBool(JoinStr(envPrefix, "LOG_LEVEL")); ok {
		cfg.DisableTimeStamp = disableTimeStamp
	}

	cfg.LogService.SetValuesFromEnv(envPrefix)
}
