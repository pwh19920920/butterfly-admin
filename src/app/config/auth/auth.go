package auth

import (
	"butterfly-admin/src/app/config"
	sysConf "github.com/pwh19920920/butterfly/config"
	"github.com/spf13/viper"
)

type auth struct {
	Auth config.AuthConfig `yaml:"auth"`
}

var conf *auth

func GetAuthConf() *config.AuthConfig {
	if conf == nil {
		// 默认配置
		viper.SetDefault("auth.expireTime", 30*60)
		viper.SetDefault("auth.headerName", "Authorization")
		viper.SetDefault("auth.headerType", "Bearer")

		// 加载配置
		sysConf.LoadConf(&conf, sysConf.GetOptions().ConfigFilePath)
	}
	return &conf.Auth
}
