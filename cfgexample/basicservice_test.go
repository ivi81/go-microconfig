package cfgexample_test

//TestBasicServiceCfg
/*func TestBasicServiceCfg(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	testCfg := microconfig.BasicServiceCfg{}
	typeName := GetTypeName(testCfg)

	b := LoadTestData(t, "BasicServiceCfg.yaml")

	err := yaml.Unmarshal(b, &testCfg)
	assert.NoError(t, err, typeName+": yaml Unmarshal error")

	LoadTestEnvData(t, "basicservicecfg.env")

	cfg := microconfig.BasicServiceCfg{}
	cfg.SetValuesFromEnv("BASIC_SERVICE")

	BasicServiceCfgAssert(t, testCfg, cfg, "", "")
}
*/

//BasicServiceCfgassert
/*func BasicServiceCfgAssert(t *testing.T, testCfg, Cfg microconfig.BasicServiceCfg, hiLeveTypeName, hiLevelPath string) {

	currentTypeName, fieldPath := CreateFildPathhiLevel(hiLeveTypeName, hiLevelPath, testCfg)

	currentFieldPath := strings.Join([]string{fieldPath, "DataExchangeClient"}, fieldSpliter)
	ClienNatsCfgAssert(t, testCfg.DataExchangeClient, Cfg.DataExchangeClient, currentTypeName, currentFieldPath)

	currentFieldPath = strings.Join([]string{fieldPath, "Logger"}, fieldSpliter)
	LoggerCfgAssert(t, testCfg.Logger, Cfg.Logger, currentTypeName, currentFieldPath)
}
*/
