package services

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"
	"wowtoken-api/internal/config"
	"wowtoken-api/internal/contracts"
	"wowtoken-api/internal/helpers"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BlockchainService xử lý các tương tác với hợp đồng WoWToken
type BlockchainService struct {
    client       *ethclient.Client
    contract     *contracts.WoWToken
    contractAddr common.Address
    chainID      *big.Int
}

func NewBlockchainService() (*BlockchainService, error) {
    cfg, err := config.LoadConfig()
    if err != nil {
        return nil, fmt.Errorf("không thể tải cấu hình: %v", err)
    }

    client, err := ethclient.Dial(cfg.InfuraURL)
    if err != nil {
        return nil, fmt.Errorf("không thể kết nối với Infura: %v", err)
    }

    contractAddr := common.HexToAddress(cfg.ContractAddress)
    contract, err := contracts.NewWoWToken(contractAddr, client)
    if err != nil {
        return nil, fmt.Errorf("không thể khởi tạo hợp đồng WoWToken: %v", err)
    }

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        return nil, fmt.Errorf("không thể lấy chainID: %v", err)
    }

    return &BlockchainService{
        client:       client,
        contract:     contract,
        contractAddr: contractAddr,
        chainID:      chainID,
    }, nil
}

func (s *BlockchainService) GetBankReserveAddress() (string, error) {
    opts := &bind.CallOpts{Context: context.Background()}
    bankReserve, err := s.contract.BankReserve(opts)
    if err != nil {
        return "", fmt.Errorf("lỗi khi lấy địa chỉ bank_reserve: %v", err)
    }
    return bankReserve.Hex(), nil
}

func (s *BlockchainService) IsCoOwner(ctx context.Context, address string) (bool, error) {
    if !helpers.IsValidEthAddress(address) {
        return false, errors.New("địa chỉ Ethereum không hợp lệ")
    }
    addr := common.HexToAddress(address)
    opts := &bind.CallOpts{Context: ctx}
    isCoOwner, err := s.contract.CoTokenOwner(opts, addr)
    if err != nil {
        return false, fmt.Errorf("lỗi khi kiểm tra đồng sở hữu: %v", err)
    }
    return isCoOwner, nil
}

func (s *BlockchainService) IsPaused(ctx context.Context) (bool, error) {
    opts := &bind.CallOpts{Context: ctx}
    paused, err := s.contract.Paused(opts)
    if err != nil {
        return false, fmt.Errorf("lỗi khi kiểm tra trạng thái paused: %v", err)
    }
    return paused, nil
}

func (s *BlockchainService) GetBalance(ctx context.Context, address string) (*big.Int, error) {
    if !helpers.IsValidEthAddress(address) {
        return nil, errors.New("địa chỉ Ethereum không hợp lệ")
    }
    addr := common.HexToAddress(address)
    opts := &bind.CallOpts{Context: ctx}
    balance, err := s.contract.Balances(opts, addr)
    if err != nil {
        return nil, fmt.Errorf("lỗi khi truy vấn số dư: %v", err)
    }
    return balance, nil
}

func (s *BlockchainService) GetCurrentTotalToken(ctx context.Context) (*big.Int, error) {
    opts := &bind.CallOpts{Context: ctx}
    total, err := s.contract.CurrentTotalToken(opts)
    if err != nil {
        return nil, fmt.Errorf("lỗi khi truy vấn tổng token: %v", err)
    }
    return total, nil
}

func (s *BlockchainService) TransferToken(from, to, amount, privateKey string) (string, error) {
    // Kiểm tra địa chỉ
    if !helpers.IsValidEthAddress(to) {
        return "", errors.New("địa chỉ nhận không hợp lệ")
    }
    amountBigInt, ok := new(big.Int).SetString(amount, 10)
    if !ok || amountBigInt.Cmp(big.NewInt(0)) <= 0 {
        return "", errors.New("số lượng token không hợp lệ")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.IsPaused(context.Background())
    if err != nil {
        return "", err
    }
    if paused {
        return "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Kiểm tra số dư
    balance, err := s.GetBalance(context.Background(), from)
    if err != nil {
        return "", err
    }
    if balance.Cmp(amountBigInt) < 0 {
        return "", errors.New("số dư không đủ để chuyển")
    }

    privKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        return "", fmt.Errorf("lỗi khi parse khóa riêng: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privKey, s.chainID)
    if err != nil {
        return "", fmt.Errorf("lỗi khi tạo transactor: %v", err)
    }
    tx, err := s.contract.TransferToken(auth, common.HexToAddress(to), amountBigInt)
    if err != nil {
        return "", fmt.Errorf("lỗi khi chuyển token: %v", err)
    }
    return tx.Hash().Hex(), nil
}

func (s *BlockchainService) ApproveToken(owner, spender, amount, privateKey string) (string, error) {
    // Kiểm tra địa chỉ
    if !helpers.IsValidEthAddress(spender) {
        return "", errors.New("địa chỉ chi tiêu không hợp lệ")
    }
    amountBigInt, ok := new(big.Int).SetString(amount, 10)
    if !ok || amountBigInt.Cmp(big.NewInt(0)) < 0 {
        return "", errors.New("số lượng token không hợp lệ")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.IsPaused(context.Background())
    if err != nil {
        return "", err
    }
    if paused {
        return "", errors.New("hợp đồng đang bị tạm dừng")
    }

    privKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        return "", fmt.Errorf("lỗi khi parse khóa riêng: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privKey, s.chainID)
    if err != nil {
        return "", fmt.Errorf("lỗi khi tạo transactor: %v", err)
    }
    tx, err := s.contract.Approve(auth, common.HexToAddress(spender), amountBigInt)
    if err != nil {
        return "", fmt.Errorf("lỗi khi phê duyệt token: %v", err)
    }
    return tx.Hash().Hex(), nil
}

func (s *BlockchainService) BurnToken(address, amount, privateKey string) (string, error) {
    // Kiểm tra số lượng
    amountBigInt, ok := new(big.Int).SetString(amount, 10)
    if !ok || amountBigInt.Cmp(big.NewInt(0)) <= 0 {
        return "", errors.New("số lượng token không hợp lệ")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.IsPaused(context.Background())
    if err != nil {
        return "", err
    }
    if paused {
        return "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Kiểm tra số dư
    balance, err := s.GetBalance(context.Background(), address)
    if err != nil {
        return "", err
    }
    if balance.Cmp(amountBigInt) < 0 {
        return "", errors.New("số dư không đủ để đốt")
    }

    privKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        return "", fmt.Errorf("lỗi khi parse khóa riêng: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privKey, s.chainID)
    if err != nil {
        return "", fmt.Errorf("lỗi khi tạo transactor: %v", err)
    }
    tx, err := s.contract.BurnToken(auth, amountBigInt)
    if err != nil {
        return "", fmt.Errorf("lỗi khi đốt token: %v", err)
    }
    return tx.Hash().Hex(), nil
}

func (s *BlockchainService) MintToken(from, amount, privateKey string) (string, string, error) {
    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.IsCoOwner(context.Background(), from)
    if err != nil {
        return "", "", err
    }
    if !isCoOwner {
        return "", "", errors.New("không phải đồng sở hữu")
    }

    // Kiểm tra số lượng
    amountBigInt, ok := new(big.Int).SetString(amount, 10)
    if !ok || amountBigInt.Cmp(big.NewInt(0)) <= 0 {
        return "", "", errors.New("số lượng token không hợp lệ")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.IsPaused(context.Background())
    if err != nil {
        return "", "", err
    }
    if paused {
        return "", "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Kiểm tra giới hạn token
    currentTotal, err := s.GetCurrentTotalToken(context.Background())
    if err != nil {
        return "", "", err
    }
    limitToken := big.NewInt(1e18) // LIMIT_TOKEN = 10^18 từ hợp đồng
    if new(big.Int).Add(currentTotal, amountBigInt).Cmp(limitToken) > 0 {
        return "", "", errors.New("số lượng mint vượt quá giới hạn")
    }

    privKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi parse khóa riêng: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privKey, s.chainID)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi tạo transactor: %v", err)
    }
    tx, err := s.contract.MintToken(auth, amountBigInt)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi tạo sự kiện mint: %v", err)
    }
    eventID := "mint_" + fmt.Sprintf("%d", time.Now().Unix())
    return eventID, tx.Hash().Hex(), nil
}

func (s *BlockchainService) SignEvent(from, eventID string, approve bool, privateKey string) (string, error) {
    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.IsCoOwner(context.Background(), from)
    if err != nil {
        return "", err
    }
    if !isCoOwner {
        return "", errors.New("không phải đồng sở hữu")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.IsPaused(context.Background())
    if err != nil {
        return "", err
    }
    if paused {
        return "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Kiểm tra sự kiện
    opts := &bind.CallOpts{Context: context.Background()}
    event, err := s.contract.EventRequireMultiSignature(opts, eventID)
    if err != nil {
        return "", fmt.Errorf("lỗi khi truy vấn sự kiện: %v", err)
    }
    if event.IsCompleted {
        return "", errors.New("sự kiện đã hoàn thành")
    }

    // Kiểm tra đã ký chưa
    hasSigned, err := s.contract.EventSigners(opts, eventID, common.HexToAddress(from))
    if err != nil {
        return "", fmt.Errorf("lỗi khi kiểm tra trạng thái ký: %v", err)
    }
    if hasSigned {
        return "", errors.New("đã ký sự kiện này")
    }

    privKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        return "", fmt.Errorf("lỗi khi parse khóa riêng: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privKey, s.chainID)
    if err != nil {
        return "", fmt.Errorf("lỗi khi tạo transactor: %v", err)
    }
    tx, err := s.contract.SignEvent(auth, eventID, approve)
    if err != nil {
        return "", fmt.Errorf("lỗi khi ký sự kiện: %v", err)
    }
    return tx.Hash().Hex(), nil
}

func (s *BlockchainService) AddCoOwner(from, newCoOwner, privateKey string) (string, string, error) {
    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.IsCoOwner(context.Background(), from)
    if err != nil {
        return "", "", err
    }
    if !isCoOwner {
        return "", "", errors.New("không phải đồng sở hữu")
    }

    // Kiểm tra địa chỉ mới
    if !helpers.IsValidEthAddress(newCoOwner) {
        return "", "", errors.New("địa chỉ đồng sở hữu mới không hợp lệ")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.IsPaused(context.Background())
    if err != nil {
        return "", "", err
    }
    if paused {
        return "", "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Kiểm tra xem địa chỉ đã là đồng sở hữu chưa
    isAlreadyCoOwner, err := s.IsCoOwner(context.Background(), newCoOwner)
    if err != nil {
        return "", "", err
    }
    if isAlreadyCoOwner {
        return "", "", errors.New("địa chỉ đã là đồng sở hữu")
    }

    privKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi parse khóa riêng: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privKey, s.chainID)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi tạo transactor: %v", err)
    }
    tx, err := s.contract.AddCoOwner(auth, common.HexToAddress(newCoOwner))
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi tạo sự kiện thêm đồng sở hữu: %v", err)
    }
    eventID := "add_co_" + fmt.Sprintf("%d", time.Now().Unix())
    return eventID, tx.Hash().Hex(), nil
}

func (s *BlockchainService) GenerateBurnTokenEvent(from, amount, privateKey string) (string, string, error) {
    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.IsCoOwner(context.Background(), from)
    if err != nil {
        return "", "", err
    }
    if !isCoOwner {
        return "", "", errors.New("không phải đồng sở hữu")
    }

    // Kiểm tra số lượng
    amountBigInt, ok := new(big.Int).SetString(amount, 10)
    if !ok || amountBigInt.Cmp(big.NewInt(0)) <= 0 {
        return "", "", errors.New("số lượng token không hợp lệ")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.IsPaused(context.Background())
    if err != nil {
        return "", "", err
    }
    if paused {
        return "", "", errors.New("hợp đồng đang bị tạm dừng")
    }

    // Kiểm tra số dư bank_reserve
    reserveAddr, err := s.GetBankReserveAddress()
    if err != nil {
        return "", "", err
    }
    reserveBalance, err := s.GetBalance(context.Background(), reserveAddr)
    if err != nil {
        return "", "", err
    }
    if reserveBalance.Cmp(amountBigInt) < 0 {
        return "", "", errors.New("số dư dự trữ không đủ để đốt")
    }

    privKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi parse khóa riêng: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privKey, s.chainID)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi tạo transactor: %v", err)
    }
    tx, err := s.contract.GenerateBurnTokenEvent(auth, amountBigInt)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi tạo sự kiện đốt token: %v", err)
    }
    eventID := "burn_" + fmt.Sprintf("%d", time.Now().Unix())
    return eventID, tx.Hash().Hex(), nil
}

func (s *BlockchainService) GeneratePauseEvent(from, privateKey string) (string, string, error) {
    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.IsCoOwner(context.Background(), from)
    if err != nil {
        return "", "", err
    }
    if !isCoOwner {
        return "", "", errors.New("không phải đồng sở hữu")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.IsPaused(context.Background())
    if err != nil {
        return "", "", err
    }
    if paused {
        return "", "", errors.New("hợp đồng đã bị tạm dừng")
    }

    privKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi parse khóa riêng: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privKey, s.chainID)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi tạo transactor: %v", err)
    }
    tx, err := s.contract.GeneratePauseEvent(auth)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi tạo sự kiện tạm dừng: %v", err)
    }
    eventID := "pause_" + fmt.Sprintf("%d", time.Now().Unix())
    return eventID, tx.Hash().Hex(), nil
}

func (s *BlockchainService) GenerateUnpauseEvent(from, privateKey string) (string, string, error) {
    // Kiểm tra quyền đồng sở hữu
    isCoOwner, err := s.IsCoOwner(context.Background(), from)
    if err != nil {
        return "", "", err
    }
    if !isCoOwner {
        return "", "", errors.New("không phải đồng sở hữu")
    }

    // Kiểm tra trạng thái paused
    paused, err := s.IsPaused(context.Background())
    if err != nil {
        return "", "", err
    }
    if !paused {
        return "", "", errors.New("hợp đồng không bị tạm dừng")
    }

    privKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi parse khóa riêng: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privKey, s.chainID)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi tạo transactor: %v", err)
    }
    tx, err := s.contract.GenerateUnpauseEvent(auth)
    if err != nil {
        return "", "", fmt.Errorf("lỗi khi tạo sự kiện bỏ tạm dừng: %v", err)
    }
    eventID := "unpause_" + fmt.Sprintf("%d", time.Now().Unix())
    return eventID, tx.Hash().Hex(), nil
}