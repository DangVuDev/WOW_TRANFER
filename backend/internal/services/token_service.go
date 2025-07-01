package services

import (
    "context"
    "errors"
    "math/big"
    "time"
    "wowtoken-api/internal/helpers"
    "wowtoken-api/pkg/db"
)

// TokenService xử lý các thao tác liên quan đến token
type TokenService struct {
    blockchainService *BlockchainService
}

// NewTokenService tạo một instance mới của TokenService
func NewTokenService(blockchainService *BlockchainService) *TokenService {
    return &TokenService{blockchainService: blockchainService}
}

// GetTokenBalance truy vấn số dư token của một địa chỉ Ethereum
func (s *TokenService) GetTokenBalance(address string) (string, error) {
    // Kiểm tra tính hợp lệ của địa chỉ Ethereum
    if !helpers.IsValidEthAddress(address) {
        return "", errors.New("địa chỉ Ethereum không hợp lệ")
    }

    // Kiểm tra cache Redis
    client := db.GetRedisClient()
    cacheKey := "balance:" + address
    cachedBalance, _ := client.Get(context.Background(), cacheKey).Result()
    if cachedBalance != "" {
        return cachedBalance, nil
    }

    // Gọi BlockchainService để lấy số dư
    balance, err := s.blockchainService.GetBalance(context.Background(), address)
    if err != nil {
        return "", errors.New("lỗi khi truy vấn số dư: " + err.Error())
    }

    // Lưu số dư vào Redis
    client.Set(context.Background(), cacheKey, balance.String(), time.Hour)
    return balance.String(), nil
}

// TransferToken thực hiện chuyển token từ một địa chỉ sang địa chỉ khác
func (s *TokenService) TransferToken(from, to, amount, jwt string) (string, error) {
    // Xác thực JWT
    userAddress, err := helpers.ValidateJWT(jwt)
    if err != nil || userAddress != from {
        return "", errors.New("jWT không hợp lệ hoặc không được ủy quyền")
    }

    // Kiểm tra tính hợp lệ của địa chỉ nhận
    if !helpers.IsValidEthAddress(to) {
        return "", errors.New("dịa chỉ nhận không hợp lệ")
    }

    // Kiểm tra số lượng token
    amountBigInt, ok := new(big.Int).SetString(amount, 10)
    if !ok || amountBigInt.Cmp(big.NewInt(0)) <= 0 {
        return "", errors.New("số lượng token không hợp lệ")
    }

    // Kiểm tra số dư
    balance, err := s.blockchainService.GetBalance(context.Background(), from)
    if err != nil {
        return "", errors.New("lỗi khi kiểm tra số dư: " + err.Error())
    }
    if balance.Cmp(amountBigInt) < 0 {
        return "", errors.New("số dư không đủ để chuyển")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.blockchainService.IsPaused(context.Background())
    if err != nil {
        return "", errors.New("lỗi khi kiểm tra trạng thái hợp đồng: " + err.Error())
    }
    if paused {
        return "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Lấy khóa riêng
    encryptedPrivateKey, err := db.GetEncryptedPrivateKey(from)
    if err != nil {
        return "", errors.New("lỗi khi lấy khóa riêng: " + err.Error())
    }
    privateKey, err := helpers.DecryptPrivateKey(encryptedPrivateKey, "32-byte-encryption-key")
    if err != nil {
        return "", errors.New("lỗi khi giải mã khóa riêng: " + err.Error())
    }

    // Gọi BlockchainService để chuyển token
    txHash, err := s.blockchainService.TransferToken(from, to, amount, privateKey)
    if err != nil {
        return "", errors.New("lỗi khi chuyển token: " + err.Error())
    }

    return txHash, nil
}

// ApproveToken cho phép một địa chỉ khác chi tiêu token
func (s *TokenService) ApproveToken(owner, spender, amount, jwt string) (string, error) {
    // Xác thực JWT
    userAddress, err := helpers.ValidateJWT(jwt)
    if err != nil || userAddress != owner {
        return "", errors.New("jWT không hợp lệ hoặc không được ủy quyền")
    }

    // Kiểm tra tính hợp lệ của địa chỉ spender
    if !helpers.IsValidEthAddress(spender) {
        return "", errors.New("dịa chỉ chi tiêu không hợp lệ")
    }

    // Kiểm tra số lượng
    amountBigInt, ok := new(big.Int).SetString(amount, 10)
    if !ok || amountBigInt.Cmp(big.NewInt(0)) < 0 {
        return "", errors.New("số lượng token không hợp lệ")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.blockchainService.IsPaused(context.Background())
    if err != nil {
        return "", errors.New("lỗi khi kiểm tra trạng thái hợp đồng: " + err.Error())
    }
    if paused {
        return "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Lấy khóa riêng
    encryptedPrivateKey, err := db.GetEncryptedPrivateKey(owner)
    if err != nil {
        return "", errors.New("lỗi khi lấy khóa riêng: " + err.Error())
    }
    privateKey, err := helpers.DecryptPrivateKey(encryptedPrivateKey, "32-byte-encryption-key")
    if err != nil {
        return "", errors.New("lỗi khi giải mã khóa riêng: " + err.Error())
    }

    // Gọi BlockchainService để phê duyệt
    txHash, err := s.blockchainService.ApproveToken(owner, spender, amount, privateKey)
    if err != nil {
        return "", errors.New("lỗi khi phê duyệt token: " + err.Error())
    }

    return txHash, nil
}

// BurnToken đốt token từ số dư của người dùng
func (s *TokenService) BurnToken(address, amount, jwt string) (string, error) {
    // Xác thực JWT
    userAddress, err := helpers.ValidateJWT(jwt)
    if err != nil || userAddress != address {
        return "", errors.New("JWT không hợp lệ hoặc không được ủy quyền")
    }

    // Kiểm tra số lượng
    amountBigInt, ok := new(big.Int).SetString(amount, 10)
    if !ok || amountBigInt.Cmp(big.NewInt(0)) <= 0 {
        return "", errors.New("số lượng token không hợp lệ")
    }

    // Kiểm tra số dư
    balance, err := s.blockchainService.GetBalance(context.Background(), address)
    if err != nil {
        return "", errors.New("Lỗi khi kiểm tra số dư: " + err.Error())
    }
    if balance.Cmp(amountBigInt) < 0 {
        return "", errors.New("số dư không đủ để đốt")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.blockchainService.IsPaused(context.Background())
    if err != nil {
        return "", errors.New("Lỗi khi kiểm tra trạng thái hợp đồng: " + err.Error())
    }
    if paused {
        return "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Lấy khóa riêng
    encryptedPrivateKey, err := db.GetEncryptedPrivateKey(address)
    if err != nil {
        return "", errors.New("Lỗi khi lấy khóa riêng: " + err.Error())
    }
    privateKey, err := helpers.DecryptPrivateKey(encryptedPrivateKey, "32-byte-encryption-key")
    if err != nil {
        return "", errors.New("Lỗi khi giải mã khóa riêng: " + err.Error())
    }

    // Gọi BlockchainService để đốt token
    txHash, err := s.blockchainService.BurnToken(address, amount, privateKey)
    if err != nil {
        return "", errors.New("Lỗi khi đốt token: " + err.Error())
    }

    // Xóa cache Redis
    client := db.GetRedisClient()
    cacheKey := "balance:" + address
    client.Del(context.Background(), cacheKey)
    return txHash, nil
}