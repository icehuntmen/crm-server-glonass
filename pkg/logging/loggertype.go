package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	Internal        Category = "Internal"
	MongoDB         Category = "Mongo"
	Redis           Category = "Redis"
	API             Category = "API"
	RequestResponse Category = "RequestResponse"
	Email           Category = "Email"
	IO              Category = "SocketIO"
	Prometheus      Category = "Prometheus"
	Validation      Category = "Validation"
)

const (
	Connection      SubCategory = "Connection"
	Disconnection   SubCategory = "Disconnection"
	Create          SubCategory = "Create"
	Find            SubCategory = "Find"
	Save            SubCategory = "Save"
	Insert          SubCategory = "Insert"
	Update          SubCategory = "Update"
	Delete          SubCategory = "Delete"
	CreateIndex     SubCategory = "CreateIndex"
	ExternalService SubCategory = "ExternalService"
	Api             SubCategory = "Api"
	StartUp         SubCategory = "StartUp"
	HashPassword    SubCategory = "HashPassword"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "Logger"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	Version      ExtraKey = "Version"
	RequestBody  ExtraKey = "RequestBody"
	ResponseBody ExtraKey = "ResponseBody"
	ErrorMessage ExtraKey = "ErrorMessage"
)
