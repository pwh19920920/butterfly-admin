package config

type DatabaseConfig struct {
	Dsn                   string `yaml:"dsn"`
	MaxIdleConnect        int    `yaml:"maxIdleConnect"`
	MaxOpenConnect        int    `yaml:"maxOpenConnect"`
	ConnMaxLifeTimeSecond int    `yaml:"connMaxLifeTimeSecond"`
}

type AuthConfig struct {
	ExpireTime int      `yaml:"expireTime"`
	HeaderName string   `yaml:"headerName"`
	HeaderType string   `yaml:"headerType"`
	IgnorePath []string `yaml:"ignorePath"`
}
