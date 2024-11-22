package env_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/env"
)

type TestCfg1 struct {
	Hosts          []string      `env:"HOSTS"`
	Port           int           `env:"PORT"`
	Url            string        `env:"URL"`
	SameFloatValue float64       `env:"FLOAT_VALUE"`
	Flag           bool          `env:"FLAG"`
	Time           time.Duration `env:"Time"`
}

type TestCfg2 struct {
	TestCfg1 `env:"_"`
}

func TestMain(m *testing.M) {

	if err := godotenv.Load("testdata/.test.env"); err != nil {
		log.Println(" No .env file found")
	}

	os.Exit(m.Run())
}

func TestPopulateWithEnv(t *testing.T) {

	t.Run("Test0: Передача nil в качестве указателя на структуру в env.PopulateWithEnv ", func(t *testing.T) {

		err := env.PopulateWithEnv("", nil)
		assert.EqualError(t, err, "reflect: call of reflect.Value.Elem on zero Value")
	})

	t.Run("Test1: Передача непосредственно экземпляра структуры в env.PopulateWithEnv ", func(t *testing.T) {
		testCfg := struct{}{}
		err := env.PopulateWithEnv("", testCfg)
		assert.EqualError(t, err, "reflect: call of reflect.Value.Elem on struct Value")
	})

	t.Run("Test2: Загрузка переменных окружения в пустую структуру ", func(t *testing.T) {
		testCfg := struct{}{}
		err := env.PopulateWithEnv("EMPTY_STRUCT", &testCfg)
		assert.NoError(t, err)
		assert.EqualValues(t, struct{}{}, testCfg)
	})

	t.Run("Test3: Загрузка переменных окружения в плоскую структуру", func(t *testing.T) {

		testCfg := TestCfg1{}
		err := env.PopulateWithEnv("TEST3", &testCfg)
		assert.NoError(t, err)

		assert.Len(t, testCfg.Hosts, 3)
		assert.Equal(t, testCfg.Port, 80)
		assert.Equal(t, "http:/test.url", testCfg.Url)
		assert.Equal(t, 1.00001, testCfg.SameFloatValue)
		assert.Equal(t, true, testCfg.Flag)
		//assert.Equal(t, time.Duration(), testCfg.Time)
	})

	t.Run("Test4.1: Загрузка переменных окружения в структуру имплементирующую поля других структур без дополнения префикса в переменной окружения", func(t *testing.T) {

		type TestWithOutPrefixCfg struct {
			TestCfg2 `env:"_"`
		}

		testWithOutPrefix := TestWithOutPrefixCfg{}

		err := env.PopulateWithEnv("TEST_CLIENT2", &testWithOutPrefix)

		assert.NoError(t, err)
		fmt.Println(testWithOutPrefix.Hosts)
		assert.NoError(t, err)
		//fmt.Println(testWithPrefixCfg.TestCfg1.Hosts)
		assert.Len(t, testWithOutPrefix.Hosts, 3)
		//TODO Добавить проверку сравнения срезов
		//
		assert.Equal(t, testWithOutPrefix.Port, 8080)
		assert.Equal(t, testWithOutPrefix.Flag, false)
		assert.Equal(t, testWithOutPrefix.SameFloatValue, 1.00002)
	})
	t.Run("Test4.2: Загрузка переменных окружения в структуру имплементирующую поля других структур", func(t *testing.T) {

		type TestWithPrefixCfg struct {
			TestCfg1 `env:"CLIENT1"`
			TestCfg2 `env:"CLIENT2"`
		}

		testWithPrefixCfg := TestWithPrefixCfg{}

		err := env.PopulateWithEnv("TEST", &testWithPrefixCfg)

		assert.NoError(t, err)
		//fmt.Println(testWithPrefixCfg.TestCfg1.Hosts)
		assert.Len(t, testWithPrefixCfg.TestCfg1.Hosts, 3)
		//TODO Добавить проверку сравнения срезов
		//
		assert.Equal(t, testWithPrefixCfg.TestCfg1.Port, 80)
		assert.Equal(t, testWithPrefixCfg.TestCfg1.Flag, true)
		assert.Equal(t, testWithPrefixCfg.TestCfg1.SameFloatValue, 1.00001)

		//fmt.Println(testWithPrefixCfg.TestCfg1.Hosts)
		assert.Len(t, testWithPrefixCfg.TestCfg2.Hosts, 3)
		//TODO Добавить проверку сравнения срезов
		//
		assert.Equal(t, testWithPrefixCfg.TestCfg2.Port, 8080)
		assert.Equal(t, testWithPrefixCfg.TestCfg2.Flag, false)
		assert.Equal(t, testWithPrefixCfg.TestCfg2.SameFloatValue, 1.00002)
	})

	t.Run("Test5.1: Загрузка переменных окружения в структуру содержащую поля имеющих типы других структур без дополнения префикса в переменной окружения", func(t *testing.T) {

		type TestWithOutPrefixCfg struct {
			Client1 TestCfg1 `env:"_"`
		}

		testWithOutPrefix := TestWithOutPrefixCfg{}

		err := env.PopulateWithEnv("TEST_CLIENT1", &testWithOutPrefix)

		assert.NoError(t, err)
		fmt.Println(testWithOutPrefix.Client1.Hosts)

		assert.Len(t, testWithOutPrefix.Client1.Hosts, 3)
	})
	t.Run("Test5: Load env to struct cfg with include type struct", func(t *testing.T) {

		type TestCfg struct {
			Client1 TestCfg1 `env:"CLIENT1"`
			Client2 TestCfg2 `env:"CLIENT2"`
		}
		testCfg := TestCfg{}

		err := env.PopulateWithEnv("TEST", &testCfg)
		assert.NoError(t, err)

	})
}
