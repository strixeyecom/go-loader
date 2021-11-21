package config

/*
	Created by aomerk at 2021-11-20 for project strixeye
*/

/*
	contains the configuration for the application
*/

type App struct {
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
	// EndpointWordlistPath is the path to the wordlist for the endpoint.
	// Is overridden by EndpointWordlist if exists.
	EndpointWordlistPath string `json:"endpoint_wordlist_path" mapstructure:"ENDPOINT_WORDLIST_PATH"`
	// EndpointWordlist is the wordlist for the endpoint.
	EndpointWordlist []string `json:"endpoint_wordlist" mapstructure:"ENDPOINT_WORDLIST"`
	VisitorCount     int      `json:"visitor_count" mapstructure:"VISITOR_COUNT"`
}
