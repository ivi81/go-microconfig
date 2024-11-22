package microconfig_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig"
	"gitlab.cloud.gcm/i.ippolitov/go-microconfig/env"
)

const (
	SERVICENAME = "SERVICE_TEST"
)

func TestMain(m *testing.M) {

	if err := godotenv.Load(".test.env"); err != nil {
		log.Println(" No .env file found")
	}

	os.Exit(m.Run())
}

func TestLoad(t *testing.T) {
	t.Run("Load from not exists folder", func(t *testing.T) {
		testCfg := struct{}{}
		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_NOT_EXIST_CONFIG", "NAME"))
		microconfig.Load(testCfg, EnvPrefixAsServiceName, true)
	})
	t.Run("Load from empty folders", func(t *testing.T) {
		testCfg := struct{}{}
		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_EMPTY_CONFIG", "NAME"))
		microconfig.Load(testCfg, EnvPrefixAsServiceName, true)
	})
	t.Run("Load from onlydefaults folders", func(t *testing.T) {
		testCfg := struct{}{}
		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_ONLY_DEFAULTS_CONFIG", "NAME"))
		microconfig.Load(testCfg, EnvPrefixAsServiceName, false)
	})
	t.Run("Load from decompose folders", func(t *testing.T) {
		testCfg := struct{}{}
		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_DECOMPOSE_CONFIG", "NAME"))
		microconfig.Load(testCfg, EnvPrefixAsServiceName, true)
	})
	t.Run("Load from normal folders", func(t *testing.T) {
		testCfg := struct{}{}
		EnvPrefixAsServiceName := os.Getenv(env.JoinStr("SERVICE_TEST_NORMAL_CONFIG", "NAME"))
		microconfig.Load(testCfg, EnvPrefixAsServiceName, true)
	})

}
