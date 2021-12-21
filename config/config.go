package config

import (
	"github.com/pwh19920920/butterfly-admin/config/auth"
	"github.com/pwh19920920/butterfly-admin/config/database"
	"github.com/pwh19920920/butterfly-admin/config/sequence"
	"github.com/pwh19920920/snowflake"
	"gorm.io/gorm"
)

type Config struct {
	DatabaseForGorm *gorm.DB        // 数据库
	Sequence        *snowflake.Node // 数据库序列化工具
	AuthConfig      *auth.Config    // 权限配置
}

func InitAll() Config {
	databaseForGorm := database.GetConn()
	sequenceInstance := sequence.GetSequence()
	authConfig := auth.GetAuthConf()
	return Config{
		databaseForGorm,
		sequenceInstance,
		authConfig,
	}
}
