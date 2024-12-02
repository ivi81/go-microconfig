//utils.go - содержит вспомогательные функции для пактеты microconfig

package microconfig

import "gitlab.cloud.gcm/i.ippolitov/go-microconfig/env"

// initConfigPath - устанавливает путь в файловой системе из которого грузятся конфигурационные файлы
func initConfigPath(envPrefix string) (string, error) {

	cfgPath := struct {
		ConfigPath string `env:"CONFIG_PATH"`
	}{}

	if err := env.PopulateWithEnv(envPrefix, &cfgPath); err != nil {
		return "", err
	}

	if cfgPath.ConfigPath == "" {
		cfgPath.ConfigPath = "./config"
	}
	return cfgPath.ConfigPath, nil
}
