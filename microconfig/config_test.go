package microconfig_test

import (
	"fmt"

	"io/ioutil"
	"log"

	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"
)

var (
	testEnvPrefix    string
	testDataFileName = []string{"basiccfg.env", "clientcfg.env", "servercfg.env", "simpleservercfg.env", "servicecfg.env"}
)

const (
	TestEnvPath  = "./testdata/envfiles/"
	TestYamlPath = "./testdata/yamlfiles"
	fieldSpliter = "."
)

func init() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)
	if err != nil {
		log.Fatal(err)
	}
}

// LoadTestEnvData загружает тестовые данные из env-файла в переменные окружения
func LoadTestEnvData(t *testing.T, fName string) {

	fPath := path.Join(TestEnvPath, fName)
	err := godotenv.Load(fPath)
	assert.NoError(t, err, fPath+" No .env file found")
}

// LoadTestData загружает тестовые данные из файла в срез байт
func LoadTestData(t *testing.T, fName string) []byte {

	fPath := path.Join(TestYamlPath, fName)
	b, err := ioutil.ReadFile(fPath)
	assert.NoError(t, err, fPath+": file not found or corrupted")

	return b
}

func GetTypeName(f interface{}) string {
	Type := reflect.TypeOf(f)
	return Type.Name()
}

func FieldOrTypeName(fieldName string, f interface{}) string {
	if fieldName != "" {
		return fieldName
	}
	return reflect.TypeOf(f).Name()
}

// CreateFildPathhiLevel функция формирующая текстовый путь до поля структуры в нотации TypeName.FieldName
func CreateFildPathhiLevel(hiLevelTypeName, hiLevelFieldPath string, val interface{}) (string, string) {

	currentTypeName := reflect.TypeOf(val).Name()

	if hiLevelFieldPath != "" {
		return currentTypeName, strings.Join([]string{hiLevelFieldPath, currentTypeName}, fieldSpliter)
	} else if hiLevelTypeName != "" {
		return currentTypeName, strings.Join([]string{hiLevelTypeName, currentTypeName}, fieldSpliter)
	}
	return currentTypeName, currentTypeName
}

func TestMain(m *testing.M) {

	ServiceName := os.Getenv("SERVICE_NAME")
	testEnvPrefix = ServiceName
	os.Exit(m.Run())
}

// TestJoinStr тест вспомогательной функции microconfig.JoinStr
func TestJoinStr(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	var (
		testData = []struct{ A, B, Result, Msg string }{
			{A: "", B: "", Result: "", Msg: "результатом должена быть пустая строка"},
			{A: "A", B: "", Result: "A", Msg: "результат должен быть равен 'A'"},
			{A: "", B: "B", Result: "B", Msg: "результат должен быть равен 'B'"},
			{A: "A", B: "B", Result: "A_B", Msg: "результат должен быть равен 'A_B'"},
		}
	)
	for _, td := range testData {
		assert.Equal(t, microconfig.JoinStr(td.A, td.B), td.Result, td.Msg)
	}
}

// FieldTestAssert объединяет набор утверждений для тестирования полей структур данных
func FieldTestAssert(t *testing.T, fieldName string, testVal, val interface{}) {

	msgFieldName := "Поле " + fieldName

	switch val.(type) {
	case int:
		assert.Equal(t, val, testVal, fmt.Sprintf("%s содержит %s, но должно быть эквивалентно %s", msgFieldName, val, testVal))
	case bool:
		assert.Equal(t, val, testVal, fmt.Sprintf("%s содержит %s, но должно быть эквивалентно %s", msgFieldName, val, testVal))
	case []string:
		if assert.NotEmpty(t, val, msgFieldName+":должно быть не пустое") {

			testSl := testVal.([]string)
			Sl := val.([]string)
			lt := len(testSl)
			l := len(Sl)

			assert.Equal(t, l, lt, fmt.Sprintf("%s содержит %d элементов, а должно быть %d", msgFieldName, l, lt))
			assert.Equal(t, val, testVal, fmt.Sprintf("%s содержит %s, но должны быть эквивалентно %s", msgFieldName, val, testVal))
		}

	default:
		if assert.NotEmpty(t, val, msgFieldName+":должно быть не пустое") {
			assert.Equal(t, val, testVal, msgFieldName+":элементы должны быть эквивалентны")
		}
	}
}

func TestLoad(t *testing.T) {

	os.Setenv("STAGE", "test")
	LoadTestEnvData(t, "LoadTest.env")

	testCfg := microconfig.ServiceAPICfg{}
	testCfg.SetValuesFromEnv("LOAD_TEST")

	t.Run("OnlyEnvLoad", func(t *testing.T) {
		// Проверка случая загрузки конфигурации только из переменных окружения
		os.Setenv("LOAD_TEST_CONFIG_PATH", "./testdata/config/empty")
		Cfg := microconfig.ServiceAPICfg{}
		microconfig.Load(&Cfg, "LOAD_TEST", false)

		ServiceAPICfgAssert(t, testCfg, Cfg, "", "")
	})

	t.Run("OnlyYamlLoad", func(t *testing.T) {
		// Проверка случая загрузки конфигурации только из цельных yaml-файлов (переменных окржения нет)

		os.Setenv("TEST_CONFIG_PATH", "./testdata/config/normal")
		Cfg := microconfig.ServiceAPICfg{}
		microconfig.Load(&Cfg, "TEST", false)

		ServiceAPICfgAssert(t, testCfg, Cfg, "", "")
	})

	t.Run("DecomposeLoad", func(t *testing.T) {
		// Проверка случая загрузки конфигурации из декомпозированных yaml-файлов (переменных окржения нет)
		os.Setenv("TEST_CONFIG_PATH", "./testdata/config/decompose")
		Cfg := microconfig.ServiceAPICfg{}
		microconfig.Load(&Cfg, "TEST", false)

		ServiceAPICfgAssert(t, testCfg, Cfg, "", "")

	})
	t.Run("CombineLoad", func(t *testing.T) {
		// Проверка случая комбинированной загрузки конфигурации из yaml-файлов и перменных окружения

		os.Setenv("TEST_CONFIG_PATH", "./testdata/config/onlydefaults")
		Cfg := microconfig.ServiceAPICfg{}
		microconfig.Load(&Cfg, "LOAD_TEST", false)

		ServiceAPICfgAssert(t, Cfg, Cfg, "", "")
	})

}
