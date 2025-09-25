// config.go - содержит функцию загрузки конфигурации из файлов и переменных окружения
// Package microconfig реализует функционал загрузки кофигурационных данных в структуры.
// Пакет обеспечивает загрузку данных как из указанного конфигурационного фала так и из переменных окружения среды, при этом
// значения полученные из фала перекрываются значениями из переменнх окружения в случае если они определены.
package microconfig

import (
	"context"
	"fmt"

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
func CfgLoad(cfg any, envPrefix string, verbose bool) (err error) {
	var (
		configPath  string
		configBytes []byte
	)

	//Суть вносимых измененний относительно v2.0.4 иметь возможность загрузить переменные окружения в конфигурацию, а не быть жестко привязанным в случае неудачи
	// к загрузке конфигурации из файлов. Хочется сделать этапы чтения файлов и переменных окружения друг от друга. В v2.0.4 если не задана CONFIG_PATH и STAGE
	// либо нет папок с конфигами то чтение конфигурации прекращается/

	//Устанавливаем путь в файловой системе из которого грузятся конфигурационные файлы
	if configPath, err = initConfigPath(envPrefix); err == nil {
		//Устанавливаем название среды для которой требуется загрузить конфигурацию
		envStage := config.NewEnvStage("development")
		//Загружаем конфигурацию из файлов
		if configBytes, err = config.Read(context.TODO(), envStage, config.WithConfigPath(configPath)); err == nil {
			if err = yaml.Unmarshal(configBytes, cfg); err != nil {
				return
			}
		}
	}
	if err != nil {
		fmt.Errorf("%s", err)
		err = nil
	}
	//Перекрываем загруженную конфигурацию из переменных окружения названия которых начинаются с envPrefix
	err = env.PopulateWithEnv(envPrefix, cfg)
	return
}
