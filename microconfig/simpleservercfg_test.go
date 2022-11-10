package microconfig_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"
	"gopkg.in/yaml.v2"
)

//TestSimpleServersCfg набор тестов для тестирования полей структур конфигурирования серверных параметров единичных служб
func TestSimpleServersCfg(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	LoadTestEnvData(t, "simpleservercfg.env")

	t.Run("ServerWssCfg", func(t *testing.T) {
		testCfg := microconfig.ServerWssCfg{}
		typeName := GetTypeName(testCfg)

		b := LoadTestData(t, "ServerWssCfg.yaml")
		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, typeName+": yaml Unmarshal error")

		cfg := microconfig.ServerWssCfg{}
		cfg.SetValuesFromEnv("")
		ServerWssCfgAssert(t, testCfg, cfg, "", "")

	})

	t.Run("ServerTaxiCfg", func(t *testing.T) {
		testCfg := microconfig.ServerTaxiCfg{}

		b := LoadTestData(t, "ServerTaxiCfg.yaml")
		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml Unmarshal error")

		cfg := microconfig.ServerTaxiCfg{}
		cfg.SetValuesFromEnv("")
		ServerTaxiCfgAssert(t, testCfg, cfg, "", "")
	})

	t.Run("ServerHttpCfg", func(t *testing.T) {

		testCfg := microconfig.ServerHttpCfg{}
		typeName := GetTypeName(testCfg)

		b := LoadTestData(t, "ServerHttpCfg.yaml")
		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, typeName+": yaml Unmarshal error")

		cfg := microconfig.ServerHttpCfg{}
		cfg.SetValuesFromEnv("")
		ServerHttpCfgAssert(t, testCfg, cfg, "", "")
	})

	t.Run("ServerGrpsCfg", func(t *testing.T) {

		testCfg := microconfig.ServerGrpsCfg{}
		typeName := GetTypeName(testCfg)

		b := LoadTestData(t, "ServerGrpsCfg.yaml")
		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, typeName+": yaml Unmarshal error")

		cfg := microconfig.ServerGrpsCfg{}
		cfg.SetValuesFromEnv("")
		ServerGrpcCfgAssert(t, testCfg, cfg, "", "")
	})
}

//ServerWssCfgAssert утверждения для тестирования значений в базовых полях структуры ServerWssCfg
func ServerWssCfgAssert(t *testing.T, testCfg, Cfg microconfig.ServerWssCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}

//ServerTaxiCfgAssert утверждения для тестирования значений в базовых полях структуры ServeTaxiCfg
func ServerTaxiCfgAssert(t *testing.T, testCfg, Cfg microconfig.ServerTaxiCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}

//ServerHttpCfgAssert утверждения для тестирования значений в базовых полях структуры ServeHttpCfg
func ServerHttpCfgAssert(t *testing.T, testCfg, Cfg microconfig.ServerHttpCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}

//ServerGrpcCfgAssert утверждения для тестирования значений в базовых полях структуры ServeGrpcCfg
func ServerGrpcCfgAssert(t *testing.T, testCfg, Cfg microconfig.ServerGrpsCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}
