//utils.go - содержит вспомогательные функции для пактеты env

package env

import (
	"log"
	"os"
)

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
