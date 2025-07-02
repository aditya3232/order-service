package config

var Config AppConfig

type AppConfig struct {
	Port                  int             `json:"port"`
	Appname               string          `json:"appName"`
	AppEnv                string          `json:"appEnv"`
	SignatureKey          string          `json:"signatureKey"`
	Database              Database        `json:"database"`
	RateLimiterMaxRequest float64         `json:"rateLimiterMaxRequest"`
	RateLimiterTimeSecond int             `json:"rateLimiterTimeSecond"`
	InternalService       InternalService `json:"internalService"`
	Minio                 Minio           `json:"minio"`
	Kafka                 Kafka           `json:"kafka"`
}

type Database struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnections    int    `json:"maxOpenConnections"`
	MaxLifeTimeConnection int    `json:"maxLifeTimeConnection"`
	MaxIdleConnections    int    `json:"maxIdleConnections"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

type InternalService struct {
	User    User    `json:"user"`
	Field   Field   `json:"field"`
	Payment Payment `json:"payment"`
}

type User struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}

type Field struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}

type Payment struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}

type Minio struct {
	Address    string `json:"address"`
	AccessKey  string `json:"accessKey"`
	Secret     string `json:"secret"`
	UseSsl     bool   `json:"useSsl"`
	BucketName string `json:"bucketName"`
}

type Kafka struct {
	Brokers               []string `json:"brokers"`
	TimeoutInMS           int      `json:"timeoutInMS"`
	MaxRetry              int      `json:"maxRetry"`
	Topics                []string `json:"topics"`
	GroupID               string   `json:"groupID"`
	MaxWaitTimeInMs       int      `json:"maxWaitTimeInMs"`
	MaxProcessingTimeInMs int      `json:"maxProcessingTimeInMs"`
	BackOffTimeInMs       int      `json:"backOffTimeInMs"`
}
