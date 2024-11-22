//servercfg.go - содержит структуры данных описывающие конфигурации серверных служб микросервисов
package microconfig

import "time"

//ServerAPICfg описывает параметры конфигурации API-сервера
type ServerAPICfg struct {
	WssCfg  ServerWssCfg  `yaml:"wssServer"`
	TaxiCfg ServerTaxiCfg `yaml:"taxiServer"`
	HttpCfg ServerHttpCfg `yaml:"httpServer"`
	GrpcCfg ServerGrpsCfg `yaml:"grpcServer"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ServerAPICfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ServerAPIPrefix)

	cfg.WssCfg.SetValuesFromEnv(envPref)
	cfg.TaxiCfg.SetValuesFromEnv(envPref)
	cfg.HttpCfg.SetValuesFromEnv(envPref)
	cfg.GrpcCfg.SetValuesFromEnv(envPref)
}

//ServerAuthCfg описыввает параметры конфигурации сервера авторизации
type ServerAuthCfg struct {
	ServerGrpsCfg        `yaml:",inline"`
	JWTSecretKey         string        `yaml:"jwtSecretKey"`
	JWTSecretKeyFile     string        `yaml:"jwtSecretKeyFile"`
	JWTTokenTimeDuration time.Duration `yaml:"jwtTokenTimeDuration"`
}

//SetValuesFromEnv загружает в параметры значения перменных окружения среды
func (cfg *ServerAuthCfg) SetValuesFromEnv(envPrefix string) {

	envPref := JoinStr(envPrefix, ServerAuthPrefix)

	cfg.ServerGrpsCfg.SetValuesFromEnv(envPref)

	if secretKey, ok := LookupEnv(JoinStr(envPref, "JWT_SECRET_KEY")); ok {
		cfg.JWTSecretKey = secretKey
	}
	if secretKeyFile, ok := LookupEnv(JoinStr(envPref, "JWT_SECRET_KEY_FILE")); ok {
		cfg.JWTSecretKeyFile = secretKeyFile
	}
	if tokenDuration, ok := LookupEnvAsTime(JoinStr(envPref, "JWT_DURATION_TIME")); ok {
		cfg.JWTTokenTimeDuration = tokenDuration
	}
}
