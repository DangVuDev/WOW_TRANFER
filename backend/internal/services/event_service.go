package services

import (
    "context"
    "errors"
    "fmt"
    "math/big"
    "time"
    "wowtoken-api/internal/helpers"
    "wowtoken-api/internal/models"
    "wowtoken-api/pkg/db"
)

// EventService xử lý các sự kiện đa chữ ký
type EventService struct {
    blockchainService *BlockchainService
}

// NewEventService tạo một instance mới của EventService
func NewEventService(blockchainService *BlockchainService) *EventService {
    return &EventService{blockchainService: blockchainService}
}

// CreateMintEvent tạo sự kiện mint token mới
func (s *EventService) CreateMintEvent(amount, jwt string) (string, string, error) {
    // Xác thực JWT
    address, err := helpers.ValidateJWT(jwt)
    if err != nil {
        return "", "", errors.New("jWT không hợp lệ")
    }

    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.blockchainService.IsCoOwner(context.Background(), address)
    if err != nil {
        return "", "", errors.New("Lỗi khi kiểm tra quyền đồng sở hữu: " + err.Error())
    }
    if !isCoOwner {
        return "", "", errors.New("không phải đồng sở hữu")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.blockchainService.IsPaused(context.Background())
    if err != nil {
        return "", "", errors.New("Lỗi khi kiểm tra trạng thái hợp đồng: " + err.Error())
    }
    if paused {
        return "", "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Kiểm tra số lượng mint
    amountBigInt, ok := new(big.Int).SetString(amount, 10)
    if !ok || amountBigInt.Cmp(big.NewInt(0)) <= 0 {
        return "", "", errors.New("số lượng token không hợp lệ")
    }

    // Lấy khóa riêng
    encryptedPrivateKey, err := db.GetEncryptedPrivateKey(address)
    if err != nil {
        return "", "", errors.New("Lỗi khi lấy khóa riêng: " + err.Error())
    }
    privateKey, err := helpers.DecryptPrivateKey(encryptedPrivateKey, "32-byte-encryption-key")
    if err != nil {
        return "", "", errors.New("Lỗi khi giải mã khóa riêng: " + err.Error())
    }

    // Gọi BlockchainService để tạo sự kiện mint
    eventID, txHash, err := s.blockchainService.MintToken(address, amount, privateKey)
    if err != nil {
        return "", "", errors.New("Lỗi khi tạo sự kiện mint: " + err.Error())
    }

    // Lưu sự kiện vào MongoDB
    event := models.Event{
        IDEvent:        eventID,
        IsCompleted:    false,
        EventName:      "Mint Token",
        Description:    fmt.Sprintf("Mint %s tokens", amount),
        SignatureCount: 1,
        Signers:        []string{address},
        CreatedAt:      time.Now().Unix(),
    }
    if err := db.SaveEvent(event); err != nil {
        return "", "", errors.New("Lỗi khi lưu sự kiện: " + err.Error())
    }

    // Lưu vào cache Redis
    client := db.GetRedisClient()
    cacheKey := "event:" + eventID
    client.Set(context.Background(), cacheKey, eventID, time.Hour*24)

    return eventID, txHash, nil
}

// SignEvent ký vào một sự kiện đa chữ ký
func (s *EventService) SignEvent(eventID, jwt string, approve bool) (string, error) {
    // Xác thực JWT
    address, err := helpers.ValidateJWT(jwt)
    if err != nil {
        return "", errors.New("JWT không hợp lệ")
    }

    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.blockchainService.IsCoOwner(context.Background(), address)
    if err != nil {
        return "", errors.New("Lỗi khi kiểm tra quyền đồng sở hữu: " + err.Error())
    }
    if !isCoOwner {
        return "", errors.New("không phải đồng sở hữu")
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

    // Gọi BlockchainService để ký sự kiện
    txHash, err := s.blockchainService.SignEvent(address, eventID, approve, privateKey)
    if err != nil {
        return "", errors.New("Lỗi khi ký sự kiện: " + err.Error())
    }

    // Cập nhật sự kiện trong MongoDB
    event, err := db.GetEventByID(eventID)
    if err != nil {
        return "", errors.New("sự kiện không tồn tại")
    }
    if approve {
        event.Signers = append(event.Signers, address)
        event.SignatureCount++
    }
    if err := db.UpdateEvent(event); err != nil {
        return "", errors.New("Lỗi khi cập nhật sự kiện: " + err.Error())
    }

    return txHash, nil
}

// CreateAddCoOwnerEvent tạo sự kiện thêm đồng sở hữu mới
func (s *EventService) CreateAddCoOwnerEvent(newCoOwner, jwt string) (string, string, error) {
    // Xác thực JWT
    address, err := helpers.ValidateJWT(jwt)
    if err != nil {
        return "", "", errors.New("JWT không hợp lệ")
    }

    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.blockchainService.IsCoOwner(context.Background(), address)
    if err != nil {
        return "", "", errors.New("Lỗi khi kiểm tra quyền đồng sở hữu: " + err.Error())
    }
    if !isCoOwner {
        return "", "", errors.New("không phải đồng sở hữu")
    }

    // Kiểm tra địa chỉ mới
    if !helpers.IsValidEthAddress(newCoOwner) {
        return "", "", errors.New("dịa chỉ đồng sở hữu mới không hợp lệ")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.blockchainService.IsPaused(context.Background())
    if err != nil {
        return "", "", errors.New("Lỗi khi kiểm tra trạng thái hợp đồng: " + err.Error())
    }
    if paused {
        return "", "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Lấy khóa riêng
    encryptedPrivateKey, err := db.GetEncryptedPrivateKey(address)
    if err != nil {
        return "", "", errors.New("Lỗi khi lấy khóa riêng: " + err.Error())
    }
    privateKey, err := helpers.DecryptPrivateKey(encryptedPrivateKey, "32-byte-encryption-key")
    if err != nil {
        return "", "", errors.New("Lỗi khi giải mã khóa riêng: " + err.Error())
    }

    // Gọi BlockchainService để tạo sự kiện thêm đồng sở hữu
    eventID, txHash, err := s.blockchainService.AddCoOwner(address, newCoOwner, privateKey)
    if err != nil {
        return "", "", errors.New("Lỗi khi tạo sự kiện thêm đồng sở hữu: " + err.Error())
    }

    // Lưu sự kiện vào MongoDB
    event := models.Event{
        IDEvent:        eventID,
        IsCompleted:    false,
        EventName:      "Add Co-Owner",
        Description:    fmt.Sprintf("Thêm đồng sở hữu: %s", newCoOwner),
        SignatureCount: 1,
        Signers:        []string{address},
        CreatedAt:      time.Now().Unix(),
    }
    if err := db.SaveEvent(event); err != nil {
        return "", "", errors.New("Lỗi khi lưu sự kiện: " + err.Error())
    }

    return eventID, txHash, nil
}

// CreateBurnTokenEvent tạo sự kiện đốt token từ bank_reserve
func (s *EventService) CreateBurnTokenEvent(amount, jwt string) (string, string, error) {
    // Xác thực JWT
    address, err := helpers.ValidateJWT(jwt)
    if err != nil {
        return "", "", errors.New("JWT không hợp lệ")
    }

    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.blockchainService.IsCoOwner(context.Background(), address)
    if err != nil {
        return "", "", errors.New("Lỗi khi kiểm tra quyền đồng sở hữu: " + err.Error())
    }
    if !isCoOwner {
        return "", "", errors.New("không phải đồng sở hữu")
    }

    // Kiểm tra số lượng
    amountBigInt, ok := new(big.Int).SetString(amount, 10)
    if !ok || amountBigInt.Cmp(big.NewInt(0)) <= 0 {
        return "", "", errors.New("số lượng token không hợp lệ")
    }

    // Kiểm tra số dư bank_reserve
    reserveAddr, err := s.blockchainService.GetBankReserveAddress()
    if err != nil {
        return "", "", errors.New("Lỗi khi lấy địa chỉ dự trữ: " + err.Error())
    }
    reserveBalance, err := s.blockchainService.GetBalance(context.Background(), reserveAddr)
    if err != nil {
        return "", "", errors.New("Lỗi khi kiểm tra số dư dự trữ: " + err.Error())
    }
    if reserveBalance.Cmp(amountBigInt) < 0 {
        return "", "", errors.New("số dư dự trữ không đủ để đốt")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.blockchainService.IsPaused(context.Background())
    if err != nil {
        return "", "", errors.New("Lỗi khi kiểm tra trạng thái hợp đồng: " + err.Error())
    }
    if paused {
        return "", "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Lấy khóa riêng
    encryptedPrivateKey, err := db.GetEncryptedPrivateKey(address)
    if err != nil {
        return "", "", errors.New("Lỗi khi lấy khóa riêng: " + err.Error())
    }
    privateKey, err := helpers.DecryptPrivateKey(encryptedPrivateKey, "32-byte-encryption-key")
    if err != nil {
        return "", "", errors.New("Lỗi khi giải mã khóa riêng: " + err.Error())
    }

    // Gọi BlockchainService để tạo sự kiện đốt token
    eventID, txHash, err := s.blockchainService.GenerateBurnTokenEvent(address, amount, privateKey)
    if err != nil {
        return "", "", errors.New("Lỗi khi tạo sự kiện đốt token: " + err.Error())
    }

    // Lưu sự kiện vào MongoDB
    event := models.Event{
        IDEvent:        eventID,
        IsCompleted:    false,
        EventName:      "Burn Token",
        Description:    fmt.Sprintf("Đốt %s tokens từ dự trữ", amount),
        SignatureCount: 1,
        Signers:        []string{address},
        CreatedAt:      time.Now().Unix(),
    }
    if err := db.SaveEvent(event); err != nil {
        return "", "", errors.New("Lỗi khi lưu sự kiện: " + err.Error())
    }

    return eventID, txHash, nil
}

// CreatePauseEvent tạo sự kiện tạm dừng hợp đồng
func (s *EventService) CreatePauseEvent(jwt string) (string, string, error) {
    // Xác thực JWT
    address, err := helpers.ValidateJWT(jwt)
    if err != nil {
        return "", "", errors.New("JWT không hợp lệ")
    }

    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.blockchainService.IsCoOwner(context.Background(), address)
    if err != nil {
        return "", "", errors.New("Lỗi khi kiểm tra quyền đồng sở hữu: " + err.Error())
    }
    if !isCoOwner {
        return "", "", errors.New("không phải đồng sở hữu")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.blockchainService.IsPaused(context.Background())
    if err != nil {
        return "", "", errors.New("Lỗi khi kiểm tra trạng thái hợp đồng: " + err.Error())
    }
    if paused {
        return "", "", errors.New("hợp đồng đã bị tạm dừng")
    }

    // Lấy khóa riêng
    encryptedPrivateKey, err := db.GetEncryptedPrivateKey(address)
    if err != nil {
        return "", "", errors.New("Lỗi khi lấy khóa riêng: " + err.Error())
    }
    privateKey, err := helpers.DecryptPrivateKey(encryptedPrivateKey, "32-byte-encryption-key")
    if err != nil {
        return "", "", errors.New("Lỗi khi giải mã khóa riêng: " + err.Error())
    }

    // Gọi BlockchainService để tạo sự kiện tạm dừng
    eventID, txHash, err := s.blockchainService.GeneratePauseEvent(address, privateKey)
    if err != nil {
        return "", "", errors.New("Lỗi khi tạo sự kiện tạm dừng: " + err.Error())
    }

    // Lưu sự kiện vào MongoDB
    event := models.Event{
        IDEvent:        eventID,
        IsCompleted:    false,
        EventName:      "Pause Contract",
        Description:    "Tạm dừng hợp đồng",
        SignatureCount: 1,
        Signers:        []string{address},
        CreatedAt:      time.Now().Unix(),
    }
    if err := db.SaveEvent(event); err != nil {
        return "", "", errors.New("Lỗi khi lưu sự kiện: " + err.Error())
    }

    return eventID, txHash, nil
}

// CreateUnpauseEvent tạo sự kiện bỏ tạm dừng hợp đồng
func (s *EventService) CreateUnpauseEvent(jwt string) (string, string, error) {
    // Xác thực JWT
    address, err := helpers.ValidateJWT(jwt)
    if err != nil {
        return "", "", errors.New("JWT không hợp lệ")
    }

    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.blockchainService.IsCoOwner(context.Background(), address)
    if err != nil {
        return "", "", errors.New("Lỗi khi kiểm tra quyền đồng sở hữu: " + err.Error())
    }
    if !isCoOwner {
        return "", "", errors.New("không phải đồng sở hữu")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.blockchainService.IsPaused(context.Background())
    if err != nil {
        return "", "", errors.New("Lỗi khi kiểm tra trạng thái hợp đồng: " + err.Error())
    }
    if !paused {
        return "", "", errors.New("hợp đồng không bị tạm dừng")
    }

    // Lấy khóa riêng
    encryptedPrivateKey, err := db.GetEncryptedPrivateKey(address)
    if err != nil {
        return "", "", errors.New("Lỗi khi lấy khóa riêng: " + err.Error())
    }
    privateKey, err := helpers.DecryptPrivateKey(encryptedPrivateKey, "32-byte-encryption-key")
    if err != nil {
        return "", "", errors.New("Lỗi khi giải mã khóa riêng: " + err.Error())
    }

    // Gọi BlockchainService để tạo sự kiện bỏ tạm dừng
    eventID, txHash, err := s.blockchainService.GenerateUnpauseEvent(address, privateKey)
    if err != nil {
        return "", "", errors.New("Lỗi khi tạo sự kiện bỏ tạm dừng: " + err.Error())
    }

    // Lưu sự kiện vào MongoDB
    event := models.Event{
        IDEvent:        eventID,
        IsCompleted:    false,
        EventName:      "Unpause Contract",
        Description:    "Bỏ tạm dừng hợp đồng",
        SignatureCount: 1,
        Signers:        []string{address},
        CreatedAt:      time.Now().Unix(),
    }
    if err := db.SaveEvent(event); err != nil {
        return "", "", errors.New("Lỗi khi lưu sự kiện: " + err.Error())
    }

    return eventID, txHash, nil
}