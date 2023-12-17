package config

type SiteInfo struct {
	CreatedAt string `yaml:"created_at" json:"created_at"`
	Title     string `yaml:"title" json:"title"`
	Version   string `yaml:"version" json:"version"`
	Name      string `yaml:"name" json:"name"`
	Job       string `yaml:"job" json:"job"`
	Addr      string `yaml:"addr" json:"addr"`
	Slogan    string `yaml:"slogan" json:"slogan"`
	SloganEn  string `yaml:"slogan_en" json:"slogan_en"`
}
