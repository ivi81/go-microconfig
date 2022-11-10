package microconfig_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"
	"gopkg.in/yaml.v2"
)

//TestLoggerCfg тест для тестирования полей структуры конфигурирования параметров логирования
func TestLoggerCfg(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	testCfg := microconfig.LoggerCfg{}

	b := LoadTestData(t, "LoggerCfg.yaml")

	err := yaml.Unmarshal(b, &testCfg)
	assert.NoError(t, err, GetTypeName(testCfg)+": yaml Unmarshal error")

	LoadTestEnvData(t, "loggercfg.env")
	cfg := microconfig.LoggerCfg{}
	cfg.SetValuesFromEnv("")

	LoggerCfgAssert(t, testCfg, cfg, "", "")
}

//LoggerCfgSute утверждения для тестирования значений в специфичных для структуры LoggerCfg полях
func LoggerCfgAssert(t *testing.T, testCfg, Cfg microconfig.LoggerCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	testLogServ := testCfg.LogService
	LogServCfg := Cfg.LogService

	currentFieldPath := strings.Join([]string{fieldPath, "LogService"}, fieldSpliter)
	BasicClientCfgAssert(t, testLogServ.BasicClientCfg, LogServCfg.BasicClientCfg, currentTypeName, currentFieldPath)

	currentFieldPath = strings.Join([]string{fieldPath, "Mode"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.Mode, Cfg.Mode)

	currentFieldPath = strings.Join([]string{fieldPath, "Path"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.Path, Cfg.Path)
}
