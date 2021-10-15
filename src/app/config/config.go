package config

import (
	"github.com/bwmarrin/snowflake"
	"github.com/pwh19920920/butterfly-admin/src/app/config/auth"
	"github.com/pwh19920920/butterfly-admin/src/app/config/database"
	"github.com/pwh19920920/butterfly-admin/src/app/config/sequence"
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
