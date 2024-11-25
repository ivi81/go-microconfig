package env_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/env"
)

type TestCfg struct {
	Hosts          []string      `env:"HOSTS"`
	Port           int           `env:"PORT"`
	Url            string        `env:"URL"`
	SameFloatValue float64       `env:"FLOAT_VALUE"`
	Flag           bool          `env:"FLAG"`
	Time           time.Duration `env:"Time"`
}

type TestCfg1 struct {
	TestCfg `env:"_"`
}

var expecterResults [2]TestCfg

func TestMain(m *testing.M) {

	if err := godotenv.Load("testdata/.test.env"); err != nil {
		log.Println(" No .env file found")
	}

	//Инициализируем ожидаемый результат тестов
	expecterResults = [2]TestCfg{
		TestCfg{
			Hosts:          []string{"host1", "host2", "host3"},
			Port:           80,
			Url:            "http:/test1.url",
			SameFloatValue: 1.00001,
			Flag:           true,
		},
		TestCfg{
			Hosts:          []string{"host4", "host5", "host6"},
			Port:           8080,
			Url:            "http:/test2.url",
			SameFloatValue: 1.00002,
			Flag:           false,
		},
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

		cfg := TestCfg{}
		err := env.PopulateWithEnv("TEST_CLIENT1", &cfg)

		if assert.NoError(t, err) {
			AssertStruct(t, &cfg, expecterResults[0])
		}
	})

	t.Run("Test4.1: Загрузка переменных окружения в структуру имплементирующую поля других структур без дополнения префикса в переменной окружения", func(t *testing.T) {

		type TestWithOutPrefixCfg struct {
			TestCfg1 `env:"_"`
		}

		cfg := TestWithOutPrefixCfg{}
		err := env.PopulateWithEnv("TEST_CLIENT2", &cfg)

		if assert.NoError(t, err) {
			AssertStruct(t, &cfg.TestCfg, expecterResults[1])
		}

	})
	t.Run("Test4.2: Загрузка переменных окружения в структуру имплементирующую поля других структур", func(t *testing.T) {

		type TestWithPrefixCfg struct {
			TestCfg  `env:"CLIENT1"`
			TestCfg1 `env:"CLIENT2"`
		}

		cfg := TestWithPrefixCfg{}
		err := env.PopulateWithEnv("TEST", &cfg)

		if assert.NoError(t, err) {
			AssertStruct(t, &cfg.TestCfg, expecterResults[0])
			AssertStruct(t, &cfg.TestCfg1.TestCfg, expecterResults[1])
		}

	})

	t.Run("Test5.1: Загрузка переменных окружения в структуру содержащую поля имеющих типы других структур без дополнения префикса в переменной окружения", func(t *testing.T) {

		type TestWithOutPrefixCfg struct {
			Client1 TestCfg `env:"_"`
		}

		cfg := TestWithOutPrefixCfg{}

		err := env.PopulateWithEnv("TEST_CLIENT1", &cfg)

		if assert.NoError(t, err) {
			AssertStruct(t, &cfg.Client1, expecterResults[0])
		}
	})
	t.Run("Test5.2: Загрузка переменных окружения в структуру содержащую поля имеющих типы других структур c дополнением префикса в переменной окружения", func(t *testing.T) {

		type testCfg struct {
			Client1 TestCfg  `env:"CLIENT1"`
			Client2 TestCfg1 `env:"CLIENT2"`
		}
		cfg := testCfg{}

		err := env.PopulateWithEnv("TEST", &cfg)
		if assert.NoError(t, err) {
			AssertStruct(t, &cfg.Client1, expecterResults[0])
			AssertStruct(t, &cfg.Client2.TestCfg, expecterResults[1])
		}

	})
}

func AssertStruct(t *testing.T, cfg *TestCfg, expectedResult TestCfg) {

	if assert.Len(t, cfg.Hosts, len(expectedResult.Hosts)) {
		assert.EqualValues(t, cfg.Hosts, expectedResult.Hosts)
	}

	assert.Equal(t, cfg.Port, expectedResult.Port)
	assert.Equal(t, cfg.Url, expectedResult.Url)
	assert.Equal(t, cfg.SameFloatValue, expectedResult.SameFloatValue)
	assert.Equal(t, cfg.Flag, expectedResult.Flag)
	//assert.Equal(t, time.Duration(), testCfg.Time)
}
