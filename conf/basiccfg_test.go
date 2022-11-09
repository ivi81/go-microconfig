package conf_test

import (
	"go-microconfig/conf"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestBasicCfg(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	LoadTestEnvData(t, "basiccfg.env")

	//тестирование полей базовой структуры BasicCfg
	t.Run("BasicCfg", func(t *testing.T) {

		testCfg := conf.BasicCfg{}
		typeName := GetTypeName(testCfg)

		b := LoadTestData(t, "BasicCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, typeName+": yaml Unmarshal error")

		cfg := conf.BasicCfg{}
		cfg.SetValuesFromEnv("")

		BasicCfgAssert(t, testCfg, cfg, "", "")
	})

	//тестирование полей базовой структуры BasicClientCfg
	t.Run("BasicClientCfg", func(t *testing.T) {

		testCfg := conf.BasicClientCfg{}
		typeName := GetTypeName(testCfg)

		b := LoadTestData(t, "BasicClientCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, typeName+": yaml Unmarshal error")

		cfg := conf.BasicClientCfg{}
		cfg.SetValuesFromEnv("BASIC_CLIENT")

		BasicClientCfgAssert(t, testCfg, cfg, "", "")
	})

	//тестирование полей базовой структуры BasicServerCfg
	t.Run("BasicServerCfg", func(t *testing.T) {

		testCfg := conf.BasicServerCfg{}
		typeName := GetTypeName(testCfg)

		b := LoadTestData(t, "BasicServerCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, typeName+": yaml.Unmarshal error")

		cfg := conf.BasicServerCfg{}
		cfg.SetValuesFromEnv("BASIC_SERVER")

		BasicServerCfgAssert(t, testCfg, cfg, "", "")
	})

	//тестирование полей базовой структуры BasicStorageCfg
	t.Run("BasicStorageCfg", func(t *testing.T) {

		testCfg := conf.BasicStorageCfg{}
		typeName := GetTypeName(testCfg)

		b := LoadTestData(t, "BasicStorageCfg.yaml")

		err := yaml.Unmarshal(b, &testCfg)
		assert.NoError(t, err, typeName+": yaml.Unmarshal error")

		cfg := conf.BasicStorageCfg{}
		cfg.SetValuesFromEnv("")
		BasicStorageCfgAssert(t, testCfg, cfg, "", "")
	})
}

//BasicCfgAssert утверждения для тестирования значений в полях структуры BasicCfg
func BasicCfgAssert(t *testing.T, testCfg, Cfg conf.BasicCfg, hiLeveTypeName, hiLevelPath string) {

	_, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	currentFieldPath := strings.Join([]string{fieldPath, "Host"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.Host, Cfg.Host)

	currentFieldPath = strings.Join([]string{fieldPath, "Port"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.Port, Cfg.Port)
}

//BasicClientCfgAssert утверждения для тестирования значений в полях структуры BasicClientCfg
func BasicClientCfgAssert(t *testing.T, testCfg, Cfg conf.BasicClientCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}

//BasicClientCfgAssert утверждения для тестирования значений в полях структуры BasicClientCfg
func BasicServerCfgAssert(t *testing.T, testCfg, Cfg conf.BasicServerCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	BasicCfgAssert(t, testCfg.BasicCfg, Cfg.BasicCfg, currentTypeName, fieldPath)
}

//BasicStorageCfgAssert утверждения для тестирования значений в полях структуры BasicStorageCfg содержащих информацию о соединении с хранилищем
func BasicStorageCfgAssert(t *testing.T, testCfg, Cfg conf.BasicStorageCfg, hiLeveTypeName, hiLevelPath string) {

	_, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	currentFieldPath := strings.Join([]string{fieldPath, "User"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.User, Cfg.User)

	currentFieldPath = strings.Join([]string{fieldPath, "Pwd"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.Pwd, Cfg.Pwd)

	currentFieldPath = strings.Join([]string{fieldPath, "Db"}, fieldSpliter)
	FieldTestAssert(t, currentFieldPath, testCfg.Db, Cfg.Db)
}
