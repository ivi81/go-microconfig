package env

import (
	"log"
	"os"
)

// InitConfigPath - устанавливает путь в файловой системе из которого грузятся конфигурационные файлы
func InitConfigPath(envPrefix string) string {

	configPath, ok := os.LookupEnv(JoinStr(envPrefix, "CONFIG_PATH"))
	if !ok {
		configPath = "./config"
	}
	return configPath
}

// LookupEnv получает содержимое переменной окружения ввиде строки.
// Имя переменной предается в параметре key.
func LookupEnv(key string) (string, bool) {

	valStr, ok := os.LookupEnv(key)
	if ok && valStr == "" {
		log.Println("переменная " + key + "содержит пустое значение")
		ok = false
	}
	return valStr, ok
}

// JoinStr - вспомогательная функция объединения двух строк через заданную в пакете
// константу разделитель strENVNameSplitter
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
