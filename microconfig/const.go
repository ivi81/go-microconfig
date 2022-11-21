//const.go - содержит константы хранящие префиксы названий в переменных окружения и различные служебные символы
package microconfig

//Константы хранящие различные служебные символы
const (
	strENVNameSplitter = "_"
	strSplitter        = ","
	hostPortSplitter   = ":"

	StrENVNameSplitter = strENVNameSplitter
	StrSplitter        = strSplitter
	HostPortSplitter   = hostPortSplitter
)

//Константы хранящие префиксы названий в переменных окружения
const (
	ClientAuthStoragePrefix  = "CLIENT_AUTH_STORAGE"
	ClientNatsPrefix         = "CLIENT_NATS"
	ClientSTIXStoragePrefix  = "CLIENT_STIX_STORAGE"
	ClientLogsStoragePrefix  = "CLIENT_LOG_STORAGE"
	ClientHttpPrefix         = "CLIENT_HTTP"
	ClientAuthPrefix         = "CLIENT_AUTH"
	ServerAuthPrefix         = "SERVER_AUTH"
	ServerAPIPrefix          = "SERVER_API"
	ServerWSSPrefix          = "SERVER_WSS"
	ServerHTTPPrefix         = "SERVER_HTTP"
	ServerGRPCPrefix         = "SERVER_GRPC"
	ServerTAXIPrefix         = "SERVER_TAXI"
	ServiceAPIPrefix         = "SERVICE_API"
	ServiceAUTHrefix         = "SERVICE_AUTH"
	ServiceSTIXStoragePrefix = "SERVICE_STIX_STORAGE"
)
