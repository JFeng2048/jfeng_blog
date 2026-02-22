package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Config 应用配置
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Log    LogConfig    `mapstructure:"log"`
	System SystemConfig `mapstructure:"system"`
	CORS   CORSConfig   `mapstructure:"cors"`
}

// MySQLConfig MySQL数据库配置
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
	TablePrefix  string `mapstructure:"table_prefix"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Enabled      bool   `mapstructure:"enabled"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
	ConnTimeout  int    `mapstructure:"conn_timeout"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Name    string `mapstructure:"name"`
	Port    int    `mapstructure:"port"`
	Host      string `mapstructure:"host"`
	Timeout string `mapstructure:"timeout"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level          string `mapstructure:"level"`
	Type           string `mapstructure:"type"`
	LogDir         string `mapstructure:"log_dir"`
	LogFileName    string `mapstructure:"log_file_name"`
	LogFileMaxSize int    `mapstructure:"log_file_max_size"`
	LogFileMaxAge  int    `mapstructure:"log_file_max_age"`
	LogFileBackups int    `mapstructure:"log_file_max_backups"`
}

// SystemConfig 系统配置
type SystemConfig struct {
	Env    string `mapstructure:"env"`
	Version string `mapstructure:"version"`
}

type CORSConfig struct {
	AllowOrigins     []string `mapstructure:"allow_origins"`
	AllowMethods     []string `mapstructure:"allow_methods"`
	AllowHeaders     []string `mapstructure:"allow_headers"`
	ExposeHeaders    []string `mapstructure:"expose_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
}

// 全局配置实例
var cfg *Config

// LoadConfig 加载配置
func LoadConfig() error {
	zap.L().Info("加载配置文件")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(*os.PathError); ok {
			return fmt.Errorf("配置文件不存在: %v", err)
		}
		return fmt.Errorf("配置加载失败: %v", err)
	}

	zap.L().Info("配置加载完成")

	// 序列化到结构体
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("配置序列化失败: %v", err)
	}

	zap.L().Info("配置序列化完成")

	return nil
}

// GetConfig 获取全局配置
func GetConfig() *Config {
	return cfg
}

func GetServerConfig() ServerConfig {
	return cfg.Server
}

func GetLogConfig() LogConfig {
	return cfg.Log
}

func GetSystemConfig() SystemConfig {
	return cfg.System
}

func GetMySQLConfig() MySQLConfig {
	return cfg.MySQL
}

func GetRedisConfig() RedisConfig {
	return cfg.Redis
}

func GetCORSConfig() CORSConfig {
	return cfg.CORS
}
