package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type CertificateConfig struct {
	CertPath   string `yaml:"cert_path"`
	KeyPath    string `yaml:"key_path"`
	CommonName string `yaml:"common_name"`
}

type Config struct {
	RootCAPath         string                        `yaml:"root_ca_path"`
	IntermediateCAPath string                        `yaml:"intermediate_ca_path"`
	Certificates       map[string]CertificateConfig `yaml:"certificates"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
