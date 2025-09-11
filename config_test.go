// config.go - содержит тесты функции CFGLoad и пример использования модуля
package microconfig_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/v2"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/v2/env"
	cfgexample "gitlab.cloud.gcm/i.ippolitov/go-microconfig/v2/example"
)

const (
	SERVICENAME = "SERVICE_TEST"
	ENABLEDEBUG = true
)

func TestMain(m *testing.M) {

	if err := godotenv.Load("testdata/.test.env"); err != nil {
		log.Println(" No .env file found")
	}

	os.Exit(m.Run())
}

func TestLoad(t *testing.T) {

	t.Run("TEST0 : Load from not exists folder", func(t *testing.T) {
		testCfg := struct{}{}
		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_NOT_EXIST_CONFIG", "NAME"))
		err := microconfig.CfgLoad(&testCfg, EnvPrefixAsServiceName, ENABLEDEBUG)
		if assert.Error(t, err) {
			assert.Errorf(t, err, "no such file or directory")
		}
	})
	t.Run("TEST1 : Load from empty folders", func(t *testing.T) {
		testCfg := struct{}{}
		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_EMPTY_CONFIG", "NAME"))
		err := microconfig.CfgLoad(&testCfg, EnvPrefixAsServiceName, ENABLEDEBUG)
		if assert.Error(t, err) {
			assert.Errorf(t, err, "no default config")
		}
	})

	t.Run("TEST2 : Load from onlydefaults folders", func(t *testing.T) {
		testCfg := struct {
			Client  cfgexample.BasicClientCfg  `yaml:"client" env:"CLIENT"`
			Server  cfgexample.BasicServerCfg  `yaml:"server" env:"SERVER"`
			Storage cfgexample.BasicStorageCfg `yaml:"storage" env:"STORAGE"`
		}{}

		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_ONLY_DEFAULTS_CONFIG", "NAME"))
		err := microconfig.CfgLoad(&testCfg, EnvPrefixAsServiceName, ENABLEDEBUG)
		if assert.NoError(t, err) {

			assert.Equal(t, "client_host1", testCfg.Client.Host)
			assert.Equal(t, 3, testCfg.Client.Port)

			assert.Equal(t, "server_host", testCfg.Server.Host)
			assert.Equal(t, 2, testCfg.Server.Port)

			assert.Equal(t, "local:db", testCfg.Storage.Db)
			assert.Equal(t, "db_pwd", testCfg.Storage.Pwd)
			assert.Equal(t, "db_user", testCfg.Storage.User)
		}
	})

	t.Run("TEST3 : Load from normal folders", func(t *testing.T) {
		testCfg := struct {
			Client  cfgexample.BasicClientCfg  `yaml:"client" env:"CLIENT"`
			Server  cfgexample.BasicServerCfg  `yaml:"server" env:"SERVER"`
			Storage cfgexample.BasicStorageCfg `yaml:"storage" env:"STORAGE"`
		}{}
		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_NORMAL_CONFIG", "NAME"))
		err := microconfig.CfgLoad(&testCfg, EnvPrefixAsServiceName, ENABLEDEBUG)

		if assert.NoError(t, err) {
			assert.Equal(t, "new_client_host1", testCfg.Client.Host)
			assert.Equal(t, 4, testCfg.Client.Port)

			assert.Equal(t, "new_server_host", testCfg.Server.Host)
			assert.Equal(t, 5, testCfg.Server.Port)

			assert.Equal(t, "new_local:db", testCfg.Storage.Db)
			assert.Equal(t, "new_db_pwd", testCfg.Storage.Pwd)
			assert.Equal(t, "new_db_user", testCfg.Storage.User)
		}
	})

	t.Run("TEST4 : Load from decompose folders", func(t *testing.T) {

		testCfg := struct {
			Client  cfgexample.BasicClientCfg  `yaml:"client" env:"CLIENT"`
			Server  cfgexample.BasicServerCfg  `yaml:"server" env:"SERVER"`
			Storage cfgexample.BasicStorageCfg `yaml:"storage" env:"STORAGE"`
		}{}
		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_DECOMPOSE_CONFIG", "NAME"))
		err := microconfig.CfgLoad(&testCfg, EnvPrefixAsServiceName, ENABLEDEBUG)

		if assert.NoError(t, err) {
			assert.Equal(t, "new_client_host1", testCfg.Client.Host)
			assert.Equal(t, 4, testCfg.Client.Port)

			assert.Equal(t, "new_server_host", testCfg.Server.Host)
			assert.Equal(t, 5, testCfg.Server.Port)

			assert.Equal(t, "new_local:db", testCfg.Storage.Db)
			assert.Equal(t, "new_db_pwd", testCfg.Storage.Pwd)
			assert.Equal(t, "new_db_user", testCfg.Storage.User)
		}
	})
}

func ExampleLoad() {

	const (
		//Эта константа используется как префикс перменной окружения хранящей название сервиса
		SERVICEUNIQNAME = "MYSERVICE"
	)

	//////////////////////////////////////////////////////////////////////////////
	//Ниже задаем переменные окружения необходимые для функционирования библиотеки
	//////////////////////////////////////////////////////////////////////////////
	//Задаем название среды развертывания
	os.Setenv("STAGE", "test")

	//Создаем переменную окружения хранящую название сервиса
	EnvKeyServiceName := env.JoinStr(SERVICEUNIQNAME, "NAME")
	os.Setenv(EnvKeyServiceName, "TEST_MY_SERVICE")

	//Загружаем название сервиса из перменной окружения
	ServiceName := os.Getenv(EnvKeyServiceName)
	//Создаем перменную окружения хранящую путь к файлам конфигурации
	EnvKeyConfigPath := env.JoinStr(ServiceName, "CONFIG_PATH")
	os.Setenv(EnvKeyConfigPath, "./testdata/config/exampleconfig")

	//////////////////////////////////////////////////////////////////////////////
	//Ниже создаем переменные окружения которые кофигурируют сервис
	//и должны передаваться из вне
	//////////////////////////////////////////////////////////////////////////////

	//Загружаем критичные данные конфигурации сервиса через переменные окружения
	//например предавая их как параметры docker-контенера внутри CI/CD сессии
	EnvKeySecretOption := env.JoinStr(ServiceName, "SECRET_OPTION")
	os.Setenv(EnvKeySecretOption, "1.2")

	EnvKeyApiKey := env.JoinStr(ServiceName, "CLIENT_API_KEY")
	os.Setenv(EnvKeyApiKey, "SECRET_API_KEY_VALUE")

	EnvKeyTimeout := env.JoinStr(ServiceName, "CLIENT_TIMEOUT")
	os.Setenv(EnvKeyTimeout, "10s")

	//Описывем конфигурацию сервиса
	type ServiceOption struct {
		Option1      string   `yaml:"option1" env:"OPTION1"`
		Option2      bool     `yaml:"option2" env:"OPTION2"`
		Option3      []string `yaml:"option3" env:"OPTION3"`
		SecretOption float64  `yaml:"secretOption" env:"SECRET_OPTION"`
	}

	type Client struct {
		ServerUrl  string        `yaml:"serverUrl" env:"SERVER_URL"`
		ServerPort int           `yaml:"serverPort" env:"SERVER_PORT"`
		Timeout    time.Duration `yaml:"timeout" env:"TIMEOUT"`
		ApiKey     string        `yaml:"apiKey" env:"API_KEY"`
	}

	type Server struct {
		Url  string `yaml:"url" env:"URL"`
		Port string `yaml:"port" env:"PORT"`
	}

	Cfg := struct {
		ServiceOption `yaml:",inline" env:"-"`
		Client        Client `yaml:"client" env:"CLIENT"`
		Server        Server `yaml:"server" env:"SERVER"`
	}{}

	//Загружаем конфигурацию
	if err := microconfig.CfgLoad(&Cfg, ServiceName, false); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Cfg)
	// Output: {{default_option true [test_op1 test_op2 test_op3] 0} {https://server.url 0 10s SECRET_API_KEY_VALUE} {0.0.0.0 8181}}
}
