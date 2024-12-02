// config.go - содержит функцию загрузки конфигурации из файлов и переменных окружения
// Package microconfig реализует функционал загрузки кофигурационных данных в структуры.
// Пакет обеспечивает загрузку данных как из указанного конфигурационного фала так и из переменных окружения среды, при этом
// значения полученные из фала перекрываются значениями из переменнх окружения в случае если они определены.
package microconfig

import (
	config "github.com/spacetab-io/configuration-go"
	"github.com/spacetab-io/configuration-go/stage"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/env"

	"gopkg.in/yaml.v2"
)

// Load - главная функция пакета, загружает в переданную структуру данные конфигурации приложения.
// Сначала загружает данные из набора yaml-файлов а затем из переменных окружения названия которых
// формируются из envPrefix и значений тега "env" полей заполняемой структуры.
//
//	параметры:
//	cfg - ссылка на заполняемую структуру
//	envPrefix - префикс в названиях перменных окружения
//	verbose - включение подробного режима работы функции
func Load(cfg any, envPrefix string, verbose bool) error {

	//Устанавливаем путь в файловой системе из которого грузятся конфигурационные файлы
	configPath, err := initConfigPath(envPrefix)
	if err != nil {
		return err
	}

	envStage := stage.NewEnvStage("development")

	//Загружаем конфигурацию из файлов
	if configBytes, err := config.Read(envStage, configPath, verbose); err != nil {
		//log.Printf("%s - service config reading error: %+v", debugging.GetInfoAboutFunc(), err)
		return err
	} else if err = yaml.Unmarshal(configBytes, cfg); err != nil {
		//log.Fatalf("%s - config unmarshal error: %v", debugging.GetInfoAboutFunc(), err)
		return err

	}

	//Перекрываем загруженную конфигурацию из переменных окружения названия которых начинаются с envPrefix
	if err := env.PopulateWithEnv(envPrefix, cfg); err != nil {
		//log.Fatalf("%s - filling in from env variables error", debugging.GetInfoAboutFunc())
		return err
	}
	return nil
}
