package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
	"wowtoken-api/internal/config"
	"wowtoken-api/internal/helpers"
	"wowtoken-api/internal/models"
	"wowtoken-api/pkg/db"
)


// GenerateEthCredentials tạo địa chỉ Ethereum và khóa riêng ngẫu nhiên
func GenerateEthCredentials() (address, privateKey string, err error) {
	privateKeyBytes := make([]byte, 32)
	if _, err := rand.Read(privateKeyBytes); err != nil {
		return "", "", err
	}
	privateKey = hex.EncodeToString(privateKeyBytes)
	address = "0x" + hex.EncodeToString(privateKeyBytes)[:40] // Giả lập địa chỉ Ethereum
	return address, privateKey, nil
}

// RegisterUser đăng ký người dùng mới
func RegisterUser(email, password, name string) (string, string, error) {
	cfg, _ := config.LoadConfig()
	// Kiểm tra đầu vào
	if email == "" || !helpers.IsValidEmail(email) {
		return "", "", errors.New("email không hợp lệ")
	}
	if len(password) < 8 {
		return "", "", errors.New("mật khẩu phải có ít nhất 8 ký tự")
	}
	if name == "" {
		return "", "", errors.New("tên không được để trống")
	}

	// Kiểm tra email đã tồn tại
	if db.UserExists(email) {
		return "", "", errors.New("email đã được đăng ký")
	}

	// Tạo thông tin ví Ethereum
	address, privateKey, err := GenerateEthCredentials()
	if err != nil {
		return "", "", fmt.Errorf("lỗi tạo thông tin ví: %v", err)
	}

	// Mã hóa khóa riêng
	encryptedKey, err := helpers.EncryptPrivateKey(privateKey, cfg.EncryptionKey)
	if err != nil {
		return "", "", fmt.Errorf("lỗi mã hóa khóa riêng: %v", err)
	}

	// Hash mật khẩu
	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return "", "", fmt.Errorf("lỗi hash mật khẩu: %v", err)
	}

	// Tạo user
	user := models.User{
		UserID:             helpers.GenerateUUID(),
		Email:              email,
		HashedPassword:     hashedPassword,
		Name:               name, 
		Address:            address,
		EncryptedPrivateKey: encryptedKey,
		RegisteredAt:       time.Now().Unix(),
	}
	if err := db.SaveUser(user); err != nil {
		return "", "", fmt.Errorf("lỗi lưu người dùng: %v", err)
	}

	return user.UserID, user.Address, nil
}

// LoginUser xử lý đăng nhập người dùng
func LoginUser(email, password string) (string, string, error) {
	// Kiểm tra đầu vào
	if email == "" || !helpers.IsValidEmail(email) {
		return "", "", errors.New("email không hợp lệ")
	}

	// Lấy thông tin người dùng
	user, err := db.GetUserByEmail(email)
	if err != nil {
		return "", "", errors.New("email hoặc mật khẩu không hợp lệ")
	}
	if !helpers.CheckPassword(password, user.HashedPassword) {
		return "", "", errors.New("email hoặc mật khẩu không hợp lệ")
	}

	// Tạo JWT
	jwt, err := helpers.GenerateJWT(user.Address)
	if err != nil {
		return "", "", fmt.Errorf("lỗi tạo JWT: %v", err)
	}

	return jwt, user.Address, nil
}