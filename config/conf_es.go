package config

type Es struct {
	Ip     string `json:"ip" yaml:"ip"`
	Port   int    `json:"port" yaml:"port"`
	LogErr string `json:"log_err" yaml:"log_err"`
}
