package config

type PortNumber uint

// server settings
type ServerSettings struct {
	Port        PortNumber
	TLSCertFile string
	TLSKeyFile  string
	// specify whether to serve a light HTML page alongside the signaling API
	DebugMode bool
}
