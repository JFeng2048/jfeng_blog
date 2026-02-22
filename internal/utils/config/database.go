package config

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	zap.L().Info("连接数据库")
	// 连接数据库
	dbConfig := GetMySQLConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("数据库连接失败", zap.Error(err))
		return nil, err
	}

	zap.L().Info("数据库连接成功")
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("获取数据库连接失败", zap.Error(err))
		return nil, err
	}

	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	zap.L().Info("设置最大空闲连接数", zap.Int("max_idle_conns", dbConfig.MaxIdleConns))
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	zap.L().Info("设置最大打开连接数", zap.Int("max_open_conns", dbConfig.MaxOpenConns))
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Second)
	zap.L().Info("设置连接最大生命周期", zap.Duration("conn_max_lifetime", time.Duration(dbConfig.ConnMaxLifetime)*time.Second))

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
