package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config 应用配置
var Config = struct {
	DB struct {
		Driver   string
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
}{}

// InitConfig 初始化配置
func InitConfig() {
	// 从环境变量获取配置，如果没有则使用默认值
	Config.DB.Driver = getEnv("DB_DRIVER", "mysql")
	Config.DB.Host = getEnv("DB_HOST", "127.0.0.1")
	Config.DB.Port = getEnv("DB_PORT", "3306")
	Config.DB.User = getEnv("DB_USER", "root")
	Config.DB.Password = getEnv("DB_PASSWORD", "")
	Config.DB.DBName = getEnv("DB_NAME", "blog")
}

// InitDB 初始化数据库连接
func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.DB.User,
		Config.DB.Password,
		Config.DB.Host,
		Config.DB.Port,
		Config.DB.DBName,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
