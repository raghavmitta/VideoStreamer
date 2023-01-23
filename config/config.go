package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

var configVar *Config

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	}
	Database struct {
		DriverName string `yaml:"driver_name"`
		HostName   string `yaml:"host_name"`
		DbName     string `yaml:"db_name"`
	}
	Pagination struct {
		PageSize int `yaml:"page_size"`
	}
	Ticker struct {
		Time time.Duration `yaml:"time"`
	}
	Api struct{}
}

func LoadConfig(configPath string) {
	// Create config structure
	configVar = &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&configVar); err != nil {
		log.Fatalln(err.Error())
		return
	}
	return
}
func GetConfig() *Config {
	return configVar

}
