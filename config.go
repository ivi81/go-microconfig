// config.go - содержит функцию загрузки конфигурации из файлов и переменных окружения
// Package microconfig реализует функционал загрузки кофигурационных данных в структуры.
// Пакет обеспечивает загрузку данных как из указанного конфигурационного фала так и из переменных окружения среды, при этом
// значения полученные из фала перекрываются значениями из переменнх окружения в случае если они определены.
package microconfig

import (
	"context"

	config "github.com/spacetab-io/configuration-go"
	"gopkg.in/yaml.v2"

	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/v2/env"
)

// CfgLoad - главная функция пакета, загружает в переданную структуру данные конфигурации приложения.
// Сначала загружает данные из набора yaml-файлов а затем из переменных окружения названия которых
// формируются из envPrefix и значений тега "env" полей заполняемой структуры.
//
//		параметры:
//		cfg - ссылка на заполняемую структуру
//		envPrefix - префикс в названиях перменных окружения
//		verbose - включение подробного режима работы функции, в данной версии не работает т.к. в
//	           github.com/spacetab-io/configuration-go@v1.3.0 не реализовано логирование, есть только stub логера.
func CfgLoad(cfg any, envPrefix string, verbose bool) error {

	//Устанавливаем путь в файловой системе из которого грузятся конфигурационные файлы
	configPath, err := initConfigPath(envPrefix)
	if err != nil {
		return err
	}

	envStage := config.NewEnvStage("development") //stage.NewEnvStage("development")

	//Загружаем конфигурацию из файлов
	if configBytes, err := config.Read(context.TODO(), envStage, config.WithConfigPath(configPath)); err != nil {
		return err
	} else if err = yaml.Unmarshal(configBytes, cfg); err != nil {
		return err
	}

	//Перекрываем загруженную конфигурацию из переменных окружения названия которых начинаются с envPrefix
	if err := env.PopulateWithEnv(envPrefix, cfg); err != nil {
		return err
	}
	return nil
}
