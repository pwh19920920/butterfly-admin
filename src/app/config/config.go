package config

import (
	"butterfly-admin/src/app/config/auth"
	"butterfly-admin/src/app/config/database"
	"butterfly-admin/src/app/config/sequence"
	"github.com/bwmarrin/snowflake"
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
