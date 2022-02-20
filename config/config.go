package config

type Database struct {
	URI      string `yaml:"-"`
	Database string `yaml:"database"`
}
