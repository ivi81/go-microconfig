package microconfig_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"
	"gopkg.in/yaml.v2"
)

func TestServiceCfg(t *testing.T) {

	LoadTestEnvData(t, "servicecfg.env")

	t.Run("ServiceAPICfg", func(t *testing.T) {
		testCfg := microconfig.ServiceAPICfg{}

		b := LoadTestData(t, "ServiceAPICfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml Unmarshal error")

		cfg := microconfig.ServiceAPICfg{}
		cfg.SetValuesFromEnv("TEST")

		ServiceAPICfgAssert(t, testCfg, cfg, "", "")
	})

	t.Run("ServiceSTIXStorageCfg", func(t *testing.T) {
		testCfg := microconfig.ServiceSTIXStorageCfg{}

		b := LoadTestData(t, "ServiceSTIXStorageCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml Unmarshal error")

		cfg := microconfig.ServiceSTIXStorageCfg{}
		cfg.SetValuesFromEnv("TEST")

		ServiceSTIXStorageCfgAssert(t, testCfg, cfg, "", "")
	})

	t.Run("ServiceAuthCfg", func(t *testing.T) {
		testCfg := microconfig.ServiceAuthCfg{}

		b := LoadTestData(t, "ServiceAuthCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml Unmarshal error")

		cfg := microconfig.ServiceAuthCfg{}
		cfg.SetValuesFromEnv("TEST")

		ServiceAuthCfgAssert(t, testCfg, cfg, "", "")
	})
}

// ServerAuthCfgSute утверждения для тестирования значений в полях структуры ServerAuthCfg:
// testCfg - структура содержащая проверочные данные
// Cfg - структура значения полей которой проверяются
// hiLeveTypeName - название типа данных в который либо встроенно либо частью которого является testCfg
// hiLevelPath - текстовый путь в вышестоящей структуре к полю содержащему  testCfg
func ServiceAPICfgAssert(t *testing.T, testCfg, Cfg microconfig.ServiceAPICfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicServiceCfgAssert(t, testCfg.BasicServiceCfg, Cfg.BasicServiceCfg, currentTypeName, fieldPath)

	currentFieldPath := strings.Join([]string{fieldPath, "ServerAPI"}, fieldSpliter)
	ServerAPICfgAssert(t, testCfg.ServerAPI, Cfg.ServerAPI, currentTypeName, currentFieldPath)

	currentFieldPath = strings.Join([]string{fieldPath, "ClientAuth"}, fieldSpliter)
	ClienAutCfgAssert(t, testCfg.ClientAuth, Cfg.ClientAuth, currentTypeName, currentFieldPath)
}

// ServiceSTIXStorageCfgAssert утверждения для тестирования значений в полях структуры ServiceSTIXStorageCfg
func ServiceSTIXStorageCfgAssert(t *testing.T, testCfg, Cfg microconfig.ServiceSTIXStorageCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicServiceCfgAssert(t, testCfg.BasicServiceCfg, Cfg.BasicServiceCfg, currentTypeName, fieldPath)

	currentFieldPath := strings.Join([]string{fieldPath, "ClientSTIXStorage"}, fieldSpliter)
	ClienSTIXStorageCfgAssert(t, testCfg.ClientSTIXStorage, Cfg.ClientSTIXStorage, currentTypeName, currentFieldPath)
}

// ServiceAuthCfgAssert утверждения для тестирования значений в полях структуры ServiceAuthCfg
func ServiceAuthCfgAssert(t *testing.T, testCfg, Cfg microconfig.ServiceAuthCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	currentFieldPath := strings.Join([]string{fieldPath, "Logger"}, fieldSpliter)
	LoggerCfgAssert(t, testCfg.Logger, Cfg.Logger, currentTypeName, currentFieldPath)

	currentFieldPath = strings.Join([]string{fieldPath, "ServerAuth"}, fieldSpliter)
	ServerAuthCfgAssert(t, testCfg.ServerAuth, Cfg.ServerAuth, currentTypeName, currentFieldPath)

	currentFieldPath = strings.Join([]string{fieldPath, "ClientAuthStorage"}, fieldSpliter)
	ClienAuthStorageCfgAssert(t, testCfg.ClientAuthStorage, Cfg.ClientAuthStorage, currentTypeName, currentFieldPath)
}
