//config.go - содержит функции загрузки конфигурации из файла и переменных окружения
//Package microconfig реализует функционал загрузки кофигурационных данных в микросервисы.
//Cодержит типы данных описывающие варианты конфигурации сервисов, методы и вспомогательные функции работы с переменными окружения среды.
//Пакет обеспечивает загрузку данных как из указанного конфигурационного фала так и из переменных окружения среды, при этом
//значения полученные из фала перекрываются значениями из переменнх окружения в случае если они определены.
package microconfig

import (
	"log"

	"os"
	"strconv"
	"strings"
	"time"

	config "github.com/spacetab-io/configuration-go"
	"github.com/spacetab-io/configuration-go/stage"
	"gitlab.cloud.gcm/i.ippolitov/debugging"
	"gopkg.in/yaml.v2"
)

//Configer - интерфейсный тип для конфигураций
type Configer interface {
	SetValuesFromEnv(envPrefix string)
}

//Load основная функия пакета, загружает конфигурацию сервиса из yaml-файла.
//Так же после загрузки из файла перезаписывает значения параметров конфигурации
//в случае если для них определена альтернатива в переменных окружения
func Load(cfg Configer, envPrefix string, verbose bool) {

	configPath, ok := os.LookupEnv(JoinStr(envPrefix, "CONFIG_PATH"))
	if !ok {
		configPath = "./config"
	}

	envStage := stage.NewEnvStage("development")

	if configBytes, err := config.Read(envStage, configPath, verbose); err == nil {

		if err = yaml.Unmarshal(configBytes, cfg); err != nil {
			log.Fatalf("%s - config unmarshal error: %v", debugging.GetInfoAboutFunc(), err)
		}

	} else {
		log.Printf("%s - service config reading error: %+v", debugging.GetInfoAboutFunc(), err)
	}

	cfg.SetValuesFromEnv(envPrefix)
}

//LookupEnvAsSlice получает содержимое переменной окружения
//ввиде среза строк. Имя переменной предается в параметре key.
func LookupEnvAsSlice(key string, sep string) ([]string, bool) {

	var slice []string

	valStr, ok := LookupEnv(key)
	if ok {
		slice = strings.Split(valStr, sep)
	}
	return slice, ok
}

//LookupEnvAsInt получает содержимое переменной окружения
// ввиде целого числа. Имя переменной предается в параметре key.
func LookupEnvAsInt(key string) (int, bool) {

	var (
		value int
		err   error
	)

	valStr, ok := LookupEnv(key)
	if ok {
		if value, err = strconv.Atoi(valStr); err != nil {
			log.Println("переменная " + key + " не содержит числовое значение: " + valStr)
			ok = false
		}
	}
	return value, ok
}

//LookupEnv получает содержимое переменной окружения
// ввиде строки. Имя переменной предается в параметре key.
func LookupEnv(key string) (string, bool) {

	valStr, ok := os.LookupEnv(key)
	if ok && valStr == "" {
		log.Println("переменная " + key + "содержит пустое значение")
		ok = false
	}
	return valStr, ok
}

//LookupEnvAsBool получает содержимое переменной окружения
//ввиде булева значения. Имя переменной предается в параметре key.
func LookupEnvAsBool(key string) (bool, bool) {

	var (
		value bool
		err   error
	)

	valStr, ok := LookupEnv(key)
	if ok {
		if value, err = strconv.ParseBool(valStr); err != nil {
			log.Println("переменная " + key + " не содержит логический тип данных: " + valStr)
			ok = false
		}
	}

	return value, ok
}

//LookupEnvAsTime получает содержимое переменной окружения ввиде
//значения time.Duration. Имя переменной предается в параметре key.
func LookupEnvAsTime(key string) (time.Duration, bool) {
	var (
		value time.Duration
		err   error
	)
	valStr, ok := LookupEnv(key)
	if ok {
		if value, err = time.ParseDuration(valStr); err != nil {
			log.Println("переменная " + key + "для предстваления ввиде интервала времени должен содержать строку в формате 'NhNmNs': " + valStr)
			ok = false
		}
	}
	return value, ok
}

//JoinStr - вспомогательная функция объединения двух строк через заданную в пакете
//константу разделитель strENVNameSplitter
func JoinStr(x string, y string) string {

	lx := len(x)
	ly := len(y)

	if lx+ly == 0 {
		return ""
	}

	if lx > 0 {
		if ly > 0 {
			return x + strENVNameSplitter + y
		}
		return x
	}

	return y
}
