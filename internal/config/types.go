package config

type PortNumber uint

// server settings
type ServerSettings struct {
	Port        PortNumber
	TLSCertFile string
	TLSKeyFile  string
}
