## Библиотека на golang для работы с конфигурацией микросервисов
## Версия v2.0.0
В данной библиотеке совмещен функционал позволяющий заружать настройки приложения из набора yaml-файлов и переменных окружения среды. Значения настроек в yaml-файлах и переменных среды могут дополнять либо перекрывать друг друга.

## Установка
* Установка последней версии *go get sourcecraft.dev/ivi-ippolitov/enummethods* 
* Установка вместе с последним коммитом из main *go get go get sourcecraft.dev/ivi-ippolitov/enummethods@main*
* Обновление до последней минорной версии *go get -u go get sourcecraft.dev/ivi-ippolitov/enummethods*

## Условия и порядок работы модуля
Для того чтобы модуль читал конфигурацию из yaml-файлов предназначенную для конкретной среды выполнения (DEV, STAGE, TEST, PROD и т.д) необходимо задание следующих преременных окружения:

* **STAGE** - содержит название среды выполнение для которой необходимо загрузить конфигурацию, если *STAGE* не зада то по умолчанию ее значение принято  *"development"*. 

* **{{EnvPrefixAsServiceName}}_CONFIG_PATH** - содержит общий путь к каталогу в котором расположены каталоги хранящие конфигурации для различных сред выполнения. *{{EnvPrefixAsServiceName}}* это произвольное значение (например название сервиса или его идентификатор) которое подставляется в качестве префикса задающий уникальность названия данной переменной, 
	* Если *{{EnvPrefixAsServiceName}}* не задан то библиотека обращается к переменно окружения *CONFIG_PATH*.
	* Если *CONFIG_PATH* вообще не задан то библиотека пытается прочитать конфигурацию из каталога "./config".

Далее происходит чтение конфигурации из переменных окружения.

## Пример испольования

```
...

//Устанавливаем путь в файловой системе из которого грузятся конфигурационные файлы
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
...

```


## Требование к структурам данных описывающим конфигурацию сервиса

Поля структур должны быть следующих типов:
* int
* float
* tyme.Duration
* string
* []string

Названия переменных окружения и полей в yaml-файлах задаются при описание структур конфигурации сервиса в коде приложения заначениями следующих тегов рефлексии:
* **yaml:"{{parametrName}}"** - задает название поля в yaml-файле. {{parametrName}} - название поля 
* **env:{{ENV_NAME}}** - задает название переменной. {{ENV_NAME}} - название переменной окружения

При встаививании структур, для обработки тегов в полях встраиваемых структур без модификации значений этих тегов в структуре встраивания указываем следующие значения:
* yaml:",inline"
* inline", env:"-"

*Примечание:* см. пример ниже
## Пример описания структур конфигурации
```
...

type ServiceOption struct {
		Option1      string   `yaml:"option1" env:"OPTION1"`
		Option2      bool     `yaml:"option2" env:"OPTION2"`
		Option3      []string `yaml:"option3" env:"OPTION3"`
		SecretOption float64  `yaml:"secretOption" env:"SECRET_OPTION"`
	}

type Server struct {
		Url  string `yaml:"url" env:"URL"`
		Port string `yaml:"port" env:"PORT"`
	}

Cfg := struct {
		ServiceOption `yaml:",inline" env:"-"`
		Server        Server `yaml:"server" env:"SERVER"`
	}{}
...

```
## Особенности описания yaml-файлов c конфигурацией
Файл конфигурации является yaml-файлом кторый должен содержать в начале название среды выполнения которое задается в переменной *STAGE*, оно же должно совпадать с названием каталога в котором располагаются данные файлы.

Например для среды выполнения ***"defaults"*** должны располагаться в каталоге с названием *"defaults"* и иметь в своем составе директиву *"defaults"*:
```
---
defaults:
  client:
    host: client_host1
    port: 3
  server:
      host: server_host
      port: 2
  storage:
    user: db_user
    pwd: db_pwd
    db: local:db
```

или например для среды выполнения ***"production"*** располагаться в каталоге с названием *"production"* и иметь в своем составе директиву *"production"*:

```
---
production:
  client:
    host: prod_client_host1
  server:
      host: prod_server_host
```

## Структура каталогов для хранения конфигурации сервиса для разных сред выполнения

Думаю все понятно из примера ниже. В данном случае переменная окружения *"CONFIG_CONF"*="./testdata/config/exampleconfig/"
```
./testdata/config/exampleconfig/
├── default
│   └── defaultcfg.yaml
├── development
│   ├── devClienCfg.yaml
│   └── devServerCfg.yaml
├── production
│   └── prodCfg.yaml
└── test
    └── testCfg.yaml
    ....
    и т.д.
```

*Примечание:* для корректного чтения файлов конфигурации в *"CONFIG_CONF"* необходимо наличие каталога *"default"*.

## Структура проекта

```
├── config.go
├── config_test.go
├── doc.go
├── env
│   ├── const.go
│   ├── doc.go
│   ├── env.load.go
│   ├── env.load_test.go
│   ├── error.go
│   ├── testdata
│   └── utils.go
├── example
│   ├── basiccfg.go
│   ├── clientcfg.go
│   ├── servercfg.go
│   ├── servicecfg.go
│   └── simpleservercfg.go
├── go.mod
├── go.sum
├── README.md
├── README.md.backup
├── testdata
│   └── config
│       ├── decompose
│       │   ├── defaults
│       │   │   └── defaultServiceCfg.yaml
│       │   └── test
│       │       ├── client.yaml
│       │       ├── db.yaml
│       │       └── server.yaml
│       ├── empty
│       ├── exampleconfig
│       │   ├── defaults
│       │   │   └── defaultcfg.yaml
│       │   ├── development
│       │   │   ├── devClienCfg.yaml
│       │   │   └── devServerCfg.yaml
│       │   ├── production
│       │   │   ├── optionCfg.yaml
│       │   │   └── prodCfg.yaml
│       │   └── test
│       │       ├── optionCfg.yaml
│       │       └── prodCfg.yaml
│       ├── normal
│       │   ├── defaults
│       │   │   └── defaultServiceCfg.yaml
│       │   └── test
│       │       └── testServiceCfg.yaml
│       ├── onlydefaults
│       │   ├── defaults
│       │   │   └── defaultServiceCfg.yaml
│       │   └── test
│       └── onlytest
│           └── test
│               └── testServiceCfg.yaml
└── utils.go
```
* **env** - логика загрузки переменных окружения в структуры данных
* **env/testdata** - тестовые данные для тестирования функционала пакета *./env*
* **example** - примеры структур конфигурации 
* **testdata/config** - тестовые данные для тестирования функционала *microconfig*
* **testdata/.test.env** - файл содержит пременные окружения необходимые для конфигурации тестов
* **testdata/config/decompose** - тестовые данные для конфигурация во многих файлах
* **testdata/config/exampleconfig** - конфигурация для использования в примере кода
* **testdata/config/empty** - тестовые данные для случая с пустой папкой
* **testdata/config/normal** - тестовые данные для случая с конфигурацией в одном файле
* **testdata/config/onlydefaults** - тестовые данные для случая с конфигурацией только default
* **testdata/config/onlytest** - тестовые данные для случая с отсутствующим каталогом default




