package admin

const (
	// DefaultBindAddress is the default bind address for the HTTP server.
	DefaultBindAddress = ":8083"
)

// Config represents the configuration for the admin service.
type Config struct {
	BindAddress      string `yaml:"bind-address"`
	HTTPSEnabled     bool   `yaml:"https-enabled"`
	HTTPSCertificate string `yaml:"https-certificate"`
	Version          string `yaml:"version"`
}

// NewConfig returns an instance of Config with defaults.
func NewConfig() Config {
	return Config{
		BindAddress:      DefaultBindAddress,
		HTTPSEnabled:     false,
		HTTPSCertificate: "/etc/ssl/influxdb.pem",
	}
}
