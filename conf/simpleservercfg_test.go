package conf_test

import (
	"go-microconfig/conf"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

//TestSimpleServersCfg набор тестов для тестирования полей структур конфигурирования серверных параметров единичных служб
func TestSimpleServersCfg(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	LoadTestEnvData(t, "simpleservercfg.env")

	t.Run("ServerWssCfg", func(t *testing.T) {
		testCfg := conf.ServerWssCfg{}
		typeName := GetTypeName(testCfg)

		b := LoadTestData(t, "ServerWssCfg.yaml")
		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, typeName+": yaml Unmarshal error")

		cfg := conf.ServerWssCfg{}
		cfg.SetValuesFromEnv("")
		ServerWssCfgAssert(t, testCfg, cfg, "", "")

	})

	t.Run("ServerTaxiCfg", func(t *testing.T) {
		testCfg := conf.ServerTaxiCfg{}

		b := LoadTestData(t, "ServerTaxiCfg.yaml")
		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml Unmarshal error")

		cfg := conf.ServerTaxiCfg{}
		cfg.SetValuesFromEnv("")
		ServerTaxiCfgAssert(t, testCfg, cfg, "", "")
	})

	t.Run("ServerHttpCfg", func(t *testing.T) {

		testCfg := conf.ServerHttpCfg{}
		typeName := GetTypeName(testCfg)

		b := LoadTestData(t, "ServerHttpCfg.yaml")
		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, typeName+": yaml Unmarshal error")

		cfg := conf.ServerHttpCfg{}
		cfg.SetValuesFromEnv("")
		ServerHttpCfgAssert(t, testCfg, cfg, "", "")
	})

	t.Run("ServerGrpsCfg", func(t *testing.T) {

		testCfg := conf.ServerGrpsCfg{}
		typeName := GetTypeName(testCfg)

		b := LoadTestData(t, "ServerGrpsCfg.yaml")
		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, typeName+": yaml Unmarshal error")

		cfg := conf.ServerGrpsCfg{}
		cfg.SetValuesFromEnv("")
		ServerGrpcCfgAssert(t, testCfg, cfg, "", "")
	})
}

//ServerWssCfgAssert утверждения для тестирования значений в базовых полях структуры ServerWssCfg
func ServerWssCfgAssert(t *testing.T, testCfg, Cfg conf.ServerWssCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}

//ServerTaxiCfgAssert утверждения для тестирования значений в базовых полях структуры ServeTaxiCfg
func ServerTaxiCfgAssert(t *testing.T, testCfg, Cfg conf.ServerTaxiCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}

//ServerHttpCfgAssert утверждения для тестирования значений в базовых полях структуры ServeHttpCfg
func ServerHttpCfgAssert(t *testing.T, testCfg, Cfg conf.ServerHttpCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}

//ServerGrpcCfgAssert утверждения для тестирования значений в базовых полях структуры ServeGrpcCfg
func ServerGrpcCfgAssert(t *testing.T, testCfg, Cfg conf.ServerGrpsCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}
