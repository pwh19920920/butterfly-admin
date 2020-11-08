package database

import (
	conf "butterfly-admin/src/app/config"
	"github.com/pwh19920920/butterfly/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type dbConf struct {
	Db conf.DatabaseConfig `yaml:"db"`
}

// 初始化db
func GetConn() *gorm.DB {
	// 默认值
	viper.SetDefault("db.maxIdleConnect", 10)
	viper.SetDefault("db.maxOpenConnect", 100)
	viper.SetDefault("db.ConnMaxLifeTimeSecond", 3600)

	// 加载配置
	databaseConf := new(dbConf)
	config.LoadConf(&databaseConf, config.GetOptions().ConfigFilePath)

	// 创建连接
	db, err := gorm.Open(mysql.Open(databaseConf.Db.Dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil || db == nil {
		logrus.Panic("db connect open failure")
		return nil
	}

	// 关闭sql log
	db.Logger = logger.Default.LogMode(logger.Silent)

	// 打开连接
	sqlDB, err := db.DB()
	if err != nil || sqlDB == nil {
		logrus.Panic("db open failure")
		return nil
	}

	// 连接池设置
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(databaseConf.Db.MaxIdleConnect)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(databaseConf.Db.MaxOpenConnect)

	// SetConnMaxLifeTime 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(time.Duration(databaseConf.Db.ConnMaxLifeTimeSecond) * time.Second)

	// 初始化全局
	return db
}
