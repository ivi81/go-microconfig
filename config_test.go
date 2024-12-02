package microconfig_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/cfgexample"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/env"
)

const (
	SERVICENAME = "SERVICE_TEST"
	ENABLEDEBUG = false
)

func TestMain(m *testing.M) {

	if err := godotenv.Load(".test.env"); err != nil {
		log.Println(" No .env file found")
	}

	os.Exit(m.Run())
}

func TestLoad(t *testing.T) {

	t.Run("TEST0 : Load from not exists folder", func(t *testing.T) {
		testCfg := struct{}{}
		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_NOT_EXIST_CONFIG", "NAME"))
		err := microconfig.Load(&testCfg, EnvPrefixAsServiceName, ENABLEDEBUG)
		if assert.Error(t, err) {
			assert.ErrorContains(t, err, "no such file or directory")
		}
	})
	t.Run("TEST1 : Load from empty folders", func(t *testing.T) {
		testCfg := struct{}{}
		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_EMPTY_CONFIG", "NAME"))
		err := microconfig.Load(&testCfg, EnvPrefixAsServiceName, ENABLEDEBUG)
		if assert.Error(t, err) {
			assert.ErrorContains(t, err, "no default config")
		}
	})

	t.Run("TEST2 : Load from onlydefaults folders", func(t *testing.T) {
		testCfg := struct {
			Client  cfgexample.BasicClientCfg  `yaml:"client" env:"CLIENT"`
			Server  cfgexample.BasicServerCfg  `yaml:"server" env:"SERVER"`
			Storage cfgexample.BasicStorageCfg `yaml:"storage" env:"STORAGE"`
		}{}

		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_ONLY_DEFAULTS_CONFIG", "NAME"))
		err := microconfig.Load(&testCfg, EnvPrefixAsServiceName, ENABLEDEBUG)
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
		err := microconfig.Load(&testCfg, EnvPrefixAsServiceName, ENABLEDEBUG)

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
		err := microconfig.Load(&testCfg, EnvPrefixAsServiceName, ENABLEDEBUG)

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
