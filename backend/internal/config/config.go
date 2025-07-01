package config

import (
	"fmt"
	"path/filepath"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port           string `mapstructure:"PORT"`
	InfuraURL      string `mapstructure:"INFURA_URL"`
	MongoDBURI     string `mapstructure:"MONGODB_URI"`
	RedisAddr      string `mapstructure:"REDIS_ADDR"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	EncryptionKey  string `mapstructure:"ENCRYPTION_KEY"`
	ContractAddress string `mapstructure:"CONTRACT_ADDRESS"`
}

func LoadConfig() (*Config, error) {
	// Lấy đường dẫn gốc của dự án (dựa trên thư mục chứa main.go)
	exePath, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("không thể lấy đường dẫn hiện tại: %v", err)
	}
	configDir := filepath.Join(exePath, "..", "..", "internal", "config") // Quay lại gốc dự án rồi vào internal/config

	// Chuyển thành đường dẫn tuyệt đối
	absConfigDir, err := filepath.Abs(configDir)
	if err != nil {
		return nil, fmt.Errorf("không thể xác định đường dẫn config: %v", err)
	}

	// Thiết lập đường dẫn để tìm config.yaml
	viper.AddConfigPath(absConfigDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // Cho phép ghi đè bằng biến môi trường

	// Đọc file config
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("không thể tải cấu hình: %v. Đường dẫn đã kiểm tra: %s", err, viper.ConfigFileUsed())
	}

	// Unmarshal vào struct Config
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("không thể parse cấu hình: %v", err)
	}

	return cfg, nil
}