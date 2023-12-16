package config

type System struct {
	host string `yaml:"host"`
	port string `yaml:"port"`
	env  string `yaml:"env"`
}
