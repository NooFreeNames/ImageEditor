// configs is used to interact with the server config
package configs

import (
	"os"

	"github.com/joho/godotenv"
)

// ConfigI defines the methods to retrieve the host, port, and site directory.
type ConfigI interface {
	GetHost() string
	GetPort() string
	GetSiteDir() string
}

// Config stores the server configuration
type Config struct {
	host    string
	port    string
	siteDir string
}

// GetHost returns the server host
func (conf *Config) GetHost() string {
	return conf.host
}

// GetPort returns the server port
func (conf *Config) GetPort() string {
	return conf.port
}

// GetSiteDir returns the site directory
func (conf *Config) GetSiteDir() string {
	return conf.siteDir
}

// New creates the server configuration by reading information from the 
// configuration file. Returns an error if the file is read unsuccessfully
func New(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}
	return &Config{
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
			os.Getenv("SITE_DIR"),
		},
		nil
}
