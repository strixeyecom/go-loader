package app

import (
	`context`
	`fmt`
	`log`
	`math`
	`math/rand`
	`net/http`
	`net/url`
	`strings`
	`sync`
	`time`
	
	`github.com/pkg/errors`
	`github.com/sirupsen/logrus`
	`github.com/strixeyecom/go-loader/internal/ip`
	`github.com/valyala/fasthttp`
	`gonum.org/v1/gonum/stat/distuv`
)

/*
	Created by aomerk at 2021-11-20 for project strixeye
*/

/*
	visitor is a struct that represents a web application visitor. it can make requests,
wait for a random time etc.
*/

// global constants for file
const (
	DefaultUserAgent        = "go-loader"
	DefaultIPSourceHeader   = `X-Forwarded-For`
	DefaultPortSourceHeader = `X-Forwarded-Port`
	DefaultSessionLength    = 35
)

var (
	DefaultEndpointList = []string{"/"}
)

// Visitor keeps track of the configuration for a visitor.
type Visitor struct {
	sync.Mutex
	// RequestWaitDistribution decides the distribution of the time it takes for a visitor to make a request
	requestWaitDistribution distuv.LogNormal
	// headers is a map of custom headers to send to the server
	headers map[string]string
	// 	Client is the HTTP client to use for the requests
	client *fasthttp.Client
	// IP is the IP address of the visitor
	ip ip.IPv4
	// requestCount is the number of requests made so far
	requestCount int
	// IPSourceHeader is the header to use for the IP address. Defaults to `X-Forwarded-For`
	IPSourceHeader string `json:"ip_source_header" mapstructure:"IP_SOURCE_HEADER"`
	// PortSourceHeader is the header to use for the port. Defaults to `X-Forwarded-Port`
	PortSourceHeader string `json:"port_source_header" mapstructure:"PORT_SOURCE_HEADER"`
	// SessionLength is the number of requests to make before ending the session
	SessionLength int `json:"session_length" mapstructure:"SESSION_LENGTH"`
	// TargetScheme is the scheme for URL to send the request to
	TargetScheme string `json:"target_scheme" mapstructure:"TARGET_SCHEME"`
	// TargetHost is the host for URL to send the request to
	TargetHost string `json:"target_host" mapstructure:"TARGET_HOST"`
	// Endpoints is the wordlist for the endpoint.
	Endpoints []string `json:"endpoint_wordlist" mapstructure:"ENDPOINTS"`
}

// SetHeader  a header for the visitor
func (v *Visitor) SetHeader(key, value string) {
	v.headers[key] = value
}

// AddHeader a value to given header for the visitor
func (v *Visitor) AddHeader(key, value string) {
	v.headers[key] = strings.Join([]string{v.headers[key], value}, ",")
}

// Ip getter for mock ipv4 address of the visitor
func (v *Visitor) Ip() ip.IPv4 {
	return v.ip
}

// SetIp setter for mock ipv4 address of the visitor. see ip/ipv4.go
func (v *Visitor) SetIp(ip ip.IPv4) {
	v.ip = ip
}

func NewVisitor() *Visitor {
	headers := make(map[string]string)
	headers["User-Agent"] = DefaultUserAgent
	randomIP := ip.NewRandom()
	headers[DefaultIPSourceHeader] = randomIP.String()
	
	v := Visitor{
		headers:                 headers,
		requestWaitDistribution: distuv.LogNormal{Mu: 3, Sigma: 1},
		client: &fasthttp.Client{
			MaxConnsPerHost: 1e4,
		},
		ip:               randomIP,
		SessionLength:    DefaultSessionLength,
		IPSourceHeader:   DefaultIPSourceHeader,
		PortSourceHeader: DefaultPortSourceHeader,
		Endpoints:        DefaultEndpointList,
	}
	
	return &v
}

// SetRequestWaitDistribution is how long it takes for a visitor to make a request, or the distribution of the time
// the visitor sleeps between requests. Must return a positive value.
// Values less than or equal to 0 will be set to 1.
func (v *Visitor) SetRequestWaitDistribution(requestWaitDistribution distuv.LogNormal) {
	v.requestWaitDistribution = requestWaitDistribution
}

// Run starts making requests to the target host.
func (v *Visitor) Run(ctx context.Context) error {
	fmt.Printf("visitor %s starts with session length %d \n", v.ip.String(), v.SessionLength)
	
	var (
		ticker *time.Ticker
	)
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()
	ticker = time.NewTicker(v.getNextWaitTime())
	for {
		select {
		case <-ticker.C:
			if v.SessionLength > 0 && v.SessionLength <= v.requestCount {
				fmt.Printf("visitor %s session ends. \n", v.ip.String())
				ticker.Stop()
				return nil
			}
			v.requestCount++
			
			endpoint := v.getNextEndpoint()
			nextWaitTime := v.getNextWaitTime()
			v.headers[DefaultPortSourceHeader] = ip.RandomPort()
			err := v.makeRequest(context.Background(), endpoint)
			if err != nil {
				return errors.Wrap(err, `failed to make request`)
			}
			logrus.Infof("%s %s %s\n", v.ip.String(), endpoint, nextWaitTime)
			
			ticker.Stop()
			// new random ticker
			ticker = time.NewTicker(nextWaitTime)
		case <-ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

func (v *Visitor) makeRequest(ctx context.Context, endpoint string) error {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()
	
	u := url.URL{
		Scheme: v.TargetScheme,
		Host:   v.TargetHost,
		Path:   endpoint,
	}
	
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodGet)
	req.SetRequestURI(u.String())
	
	for k, v := range v.headers {
		switch strings.ToLower(k) {
		// case "host":
		// 	req.SetHost(v)
		// TODO: check if there are more cases to add
		default:
			req.Header.Set(k, v)
		}
	}
	
	// execute request
	err := v.client.Do(req, nil)
	if err != nil {
		return errors.WithMessage(err, "visitor can't send request")
	}
	
	return nil
}

// getNextWaitTime returns the next wait time for the next request.
func (v *Visitor) getNextWaitTime() time.Duration {
	d := time.Duration(math.Abs(v.requestWaitDistribution.Rand()))
	if d <= 0 {
		d = 1
	}
	return d * time.Second
}

// getNextEndpoint returns the next endpoint to make a request to.
func (v *Visitor) getNextEndpoint() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(v.Endpoints) - 1)
	return v.Endpoints[r]
}
