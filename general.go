package nfon

type config struct {
	key    string // api key
	secret string // apisecret
	uri    string // api uri
	debug  bool   // debug the requests
}

type Client struct {
	config       config
	requestCount uint
}

// param api key, api secret, api uri, debug the requests
func NewClient(key string, secret string, uri string, debug bool) *Client {
	return &Client{
		config: config{
			key:    key,
			secret: secret,
			uri:    uri,
			debug:  debug,
		},
		requestCount: 0,
	}
}

func (c *Client) GetRequestCount() uint {
	return c.requestCount
}
