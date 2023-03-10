package nfon

const (
	API_URL = "https://portal-api.nfon.net:8090"
)

type Config struct {
	key    string
	secret string
	sysid  string
}

func NewConfig(sysid string, key string, secret string) *Config {
	return &Config{
		key:    key,
		secret: secret,
		sysid:  sysid,
	}
}
