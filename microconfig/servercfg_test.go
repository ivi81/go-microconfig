package microconfig_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"
	"gopkg.in/yaml.v2"
)

//TestServersCfg набор тестов для тестирования полей структур конфигурирования серверных параметров сервисов
func TestServersCfg(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	LoadTestEnvData(t, "servercfg.env")

	t.Run("ServerAPICfg", func(t *testing.T) {
		testCfg := microconfig.ServerAPICfg{}

		b := LoadTestData(t, "ServerAPICfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml Unmarshal error")

		cfg := microconfig.ServerAPICfg{}
		cfg.SetValuesFromEnv("")

		ServerAPICfgAssert(t, testCfg, cfg, "", "")

	})

	t.Run("ServerAuthCfg", func(t *testing.T) {

		testCfg := microconfig.ServerAuthCfg{}

		b := LoadTestData(t, "ServerAuthCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, GetTypeName(testCfg)+": yaml Unmarshal error")

		cfg := microconfig.ServerAuthCfg{}
		cfg.SetValuesFromEnv("")

		ServerAuthCfgAssert(t, testCfg, cfg, "", "")
	})
}

//ServerAPICfgAssert утверждения для тестирования значений в полях структуры ServerAPICfg
func ServerAPICfgAssert(t *testing.T, testCfg, Cfg microconfig.ServerAPICfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	currentFieldPath := strings.Join([]string{fieldPath, "WssCfg"}, fieldSpliter)
	ServerWssCfgAssert(t, testCfg.WssCfg, Cfg.WssCfg, currentTypeName, currentFieldPath)

	currentFieldPath = strings.Join([]string{fieldPath, "HttpCfg"}, fieldSpliter)
	ServerHttpCfgAssert(t, testCfg.HttpCfg, Cfg.HttpCfg, currentTypeName, currentFieldPath)

	currentFieldPath = strings.Join([]string{fieldPath, "TaxiCfg"}, fieldSpliter)
	ServerTaxiCfgAssert(t, testCfg.TaxiCfg, Cfg.TaxiCfg, currentTypeName, currentFieldPath)

	currentFieldPath = strings.Join([]string{fieldPath, "GrpcCfg"}, fieldSpliter)
	ServerGrpcCfgAssert(t, testCfg.GrpcCfg, testCfg.GrpcCfg, currentTypeName, currentFieldPath)
}

//ServerAuthCfgSute утверждения для тестирования значений в полях структуры ServerAuthCfg
func ServerAuthCfgAssert(t *testing.T, testCfg, Cfg microconfig.ServerAuthCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	ServerGrpcCfgAssert(t, testCfg.ServerGrpsCfg, Cfg.ServerGrpsCfg, currentTypeName, fieldPath)

	currentFieldPath := strings.Join([]string{fieldPath, "JWTSecretKey"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.JWTSecretKey, Cfg.JWTSecretKey)

	currentFieldPath = strings.Join([]string{fieldPath, "JWTSecretKeyFile"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.JWTSecretKeyFile, Cfg.JWTSecretKeyFile)

	currentFieldPath = strings.Join([]string{fieldPath, "JWTTokenTimeDuration"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.JWTTokenTimeDuration, Cfg.JWTTokenTimeDuration)
}
