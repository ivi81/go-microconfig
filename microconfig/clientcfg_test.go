package microconfig_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"
	"gopkg.in/yaml.v2"
)

//TestClientsCfg набор тестов для тестирования полей структур конфигурирования клиентскийх параметров
func TestClientsCfg(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	LoadTestEnvData(t, "clientcfg.env")

	t.Run("ClientAuthStorageCfg", func(t *testing.T) {

		testCfg := microconfig.ClientAuthStorageCfg{}
		b := LoadTestData(t, "ClientAuthStorageCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml Unmarshal error")

		cfg := microconfig.ClientAuthStorageCfg{}
		cfg.SetValuesFromEnv("")

		ClienAuthStorageCfgAssert(t, testCfg, cfg, "", "")
	})

	t.Run("ClientSTIXStorageCfg", func(t *testing.T) {

		testCfg := microconfig.ClientSTIXStorageCfg{}
		b := LoadTestData(t, "ClientSTIXStorageCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml Unmarshal error")

		cfg := microconfig.ClientSTIXStorageCfg{}
		cfg.SetValuesFromEnv("")

		ClienSTIXStorageCfgAssert(t, testCfg, cfg, "", "")
	})

	t.Run("ClientNatsCfg", func(t *testing.T) {

		testCfg := microconfig.ClientNatsCfg{}
		b := LoadTestData(t, "ClientNatsCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml Unmarshal error")

		cfg := microconfig.ClientNatsCfg{}
		cfg.SetValuesFromEnv("")

		ClienNatsCfgAssert(t, testCfg, cfg, "", "")
	})

	t.Run("ClientAuthCfg", func(t *testing.T) {

		testCfg := microconfig.ClientAuthCfg{}
		b := LoadTestData(t, "ClientAuthCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml.Unmarshal error")

		cfg := microconfig.ClientAuthCfg{}
		cfg.SetValuesFromEnv("")

		ClienAutCfgAssert(t, testCfg, cfg, "", "")
	})

	t.Run("ClientLogsStorageCfg", func(t *testing.T) {

		testCfg := microconfig.ClientLogsStorageCfg{}
		b := LoadTestData(t, "ClientLogsStorageCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml.Unmarshal error")

		cfg := microconfig.ClientLogsStorageCfg{}
		cfg.SetValuesFromEnv("")

		ClienLogsStorageCfgAssert(t, testCfg, cfg, "", "")
	})

}

//ClienAuthStorageCfgAssert утверждения для тестирования значений в полях структуры ClienAuthStorageCfg
func ClienAuthStorageCfgAssert(t *testing.T, testCfg, Cfg microconfig.ClientAuthStorageCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	//Утверждения для основных полей
	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
	BasicStorageCfgAssert(t, testCfg.BasicStorageCfg, Cfg.BasicStorageCfg, currentTypeName, fieldPath)
}

//ClienSTIXStorageCfgAssert утверждения для тестирования значений в полях структуры ClienAuthStorageCfg
func ClienSTIXStorageCfgAssert(t *testing.T, testCfg, Cfg microconfig.ClientSTIXStorageCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	//Утверждения для основных полей
	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
	BasicStorageCfgAssert(t, testCfg.BasicStorageCfg, Cfg.BasicStorageCfg, currentTypeName, fieldPath)
}

//ClienAuthCfgAssert утверждения для тестирования значений в полях структуры ClienAuthCfg
func ClienAutCfgAssert(t *testing.T, testCfg, Cfg microconfig.ClientAuthCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	//Утверждения для основных полей
	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}

// ClienNatsCfgAssert утверждения для тестирования значений в полях структуры ClienNatsCfg
func ClienNatsCfgAssert(t *testing.T, testCfg, Cfg microconfig.ClientNatsCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	//Утверждения для основных полей
	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)

	//Утверждения для полей специфичных для ClienNatsCfg
	currentFieldPath := strings.Join([]string{fieldPath, "TurnOffEcho"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.TurnOffEcho, Cfg.TurnOffEcho)

	currentFieldPath = strings.Join([]string{fieldPath, "PingInterval"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.PingInterval, Cfg.PingInterval)

	currentFieldPath = strings.Join([]string{fieldPath, "AuthWithToken"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.AuthWithToken, Cfg.AuthWithToken)

	currentFieldPath = strings.Join([]string{fieldPath, "AuthWithUser"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.AuthWithUser, Cfg.AuthWithUser)

	currentFieldPath = strings.Join([]string{fieldPath, "AuthWithCredFile"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.AuthWithCredFile, Cfg.AuthWithCredFile)

	currentFieldPath = strings.Join([]string{fieldPath, "TLSOn"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.TLSOn, Cfg.TLSOn)
}

//ClienLogsStorageCfgAssert утверждения для тестирования значений в полях структуры ClienLogsStorageCfg
func ClienLogsStorageCfgAssert(t *testing.T, testCfg, Cfg microconfig.ClientLogsStorageCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	//Утверждения для основных полей
	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}
