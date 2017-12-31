package shortfundly

import "time"

const (
	// host is the default host of Shortfundly's API.
	host = "http://api.shortfundly.com"
	// defaultTimeout is the default timeout duration used on HTTP requests.
	defaultTimeout = 5 * time.Second
)

// Shortfundly defines the Shortfundly client
type Shortfundly struct {
	APIKey  string
	Host    string
	Timeout time.Duration
}

// New returns a new Shortfundly's API client credentials which can be used to make the requests.
func New(key string) *Shortfundly {
	return &Shortfundly{
		APIKey:  key,
		Host:    host,
		Timeout: defaultTimeout,
	}
}
