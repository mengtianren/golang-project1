package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 定义配置结构体
type Config struct {
	App struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
		DBName    string `yaml:"dbname"`
		Charset   string `yaml:"charset"`
		ParseTime string `yaml:"parseTime"`
		Loc       string `yaml:"loc"`
	} `yaml:"database"`
}

var AppConfig Config

// LoadConfig 加载配置文件
func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("无法读取配置文件: %w", err)
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		return fmt.Errorf("无法解析配置文件: %w", err)
	}

	return nil
}
