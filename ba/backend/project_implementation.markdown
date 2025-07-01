# WoWToken Go Project Basic Implementation

This document provides basic Go code for each module in the WoWToken Web API project, as per the folder structure and API documentation. The implementation uses `gorilla/mux` for routing, `go-ethereum` for blockchain interactions, `mongo-driver` for MongoDB, `go-redis` for caching, and `jwt-go` for authentication. Each module is designed to be modular, testable, and aligned with the requirements for a decentralized banking dApp with custodial wallets.

---

## 1. `cmd/api/main.go`
**Purpose**: Entry point for the API server, initializing dependencies and starting the HTTP server.

```go
package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "wowtoken-api/internal/config"
    "wowtoken-api/internal/controllers"
    "wowtoken-api/internal/middleware"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    router := mux.NewRouter()
    router.Use(middleware.AuthMiddleware)
    router.Use(middleware.RateLimitMiddleware)

    // Register routes
    controllers.RegisterTokenRoutes(router)
    controllers.RegisterUserRoutes(router)
    controllers.RegisterEventRoutes(router)

    log.Printf("Starting server on :%s", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
```

---

## 2. `internal/config/`
### `config.go`
**Purpose**: Loads configuration from environment variables or `config.yaml`.

```go
package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Port          string `mapstructure:"PORT"`
    InfuraURL     string `mapstructure:"INFURA_URL"`
    MongoDBURI    string `mapstructure:"MONGODB_URI"`
    RedisAddr     string `mapstructure:"REDIS_ADDR"`
    JWTSecret     string `mapstructure:"JWT_SECRET"`
    EncryptionKey string `mapstructure:"ENCRYPTION_KEY"`
}

func LoadConfig() (*Config, error) {
    viper.SetConfigFile("config.yaml")
    viper.AutomaticEnv()
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    cfg := &Config{}
    if err := viper.Unmarshal(cfg); err != nil {
        return nil, err
    }
    return cfg, nil
}
```

### `config.yaml`
**Purpose**: Stores configuration settings.

```yaml
PORT: "8080"
INFURA_URL: "https://mainnet.infura.io/v3/your-infura-key"
MONGODB_URI: "mongodb://localhost:27017/wowtoken"
REDIS_ADDR: "localhost:6379"
JWT_SECRET: "your-jwt-secret"
ENCRYPTION_KEY: "32-byte-encryption-key"
```

---

## 3. `internal/controllers/`
### `token_controller.go`
**Purpose**: Handles token-related HTTP endpoints (e.g., `/balances/{address}`, `/transfer`).

```go
package controllers

import (
    "net/http"
    "github.com/gorilla/mux"
    "wowtoken-api/internal/helpers"
    "wowtoken-api/internal/services"
)

func RegisterTokenRoutes(r *mux.Router) {
    r.HandleFunc("/balances/{address}", GetBalance).Methods("GET")
    r.HandleFunc("/transfer", Transfer).Methods("POST")
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    address := vars["address"]
    balance, err := services.GetTokenBalance(address)
    if err != nil {
        helpers.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }
    helpers.WriteJSON(w, http.StatusOK, map[string]string{
        "address": address,
        "balance": balance,
    })
}

func Transfer(w http.ResponseWriter, r *http.Request) {
    var req struct {
        From   string `json:"from"`
        To     string `json:"to"`
        Amount string `json:"amount"`
        JWT    string `json:"jwt"`
    }
    if err := helpers.ReadJSON(r, &req); err != nil {
        helpers.WriteError(w, http.StatusBadRequest, "Invalid request body")
        return
    }
    txHash, err := services.TransferToken(req.From, req.To, req.Amount, req.JWT)
    if err != nil {
        helpers.WriteError(w, http.StatusForbidden, err.Error())
        return
    }
    helpers.WriteJSON(w, http.StatusOK, map[string]string{
        "transactionHash": txHash,
        "status":         "success",
    })
}
```

### `user_controller.go`
**Purpose**: Handles user management endpoints (e.g., `/users/register`, `/users/login`).

```go
package controllers

import (
    "net/http"
    "github.com/gorilla/mux"
    "wowtoken-api/internal/helpers"
    "wowtoken-api/internal/services"
)

func RegisterUserRoutes(r *mux.Router) {
    r.HandleFunc("/users/register", RegisterUser).Methods("POST")
    r.HandleFunc("/users/login", LoginUser).Methods("POST")
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
        Name     string `json:"name"`
    }
    if err := helpers.ReadJSON(r, &req); err != nil {
        helpers.WriteError(w, http.StatusBadRequest, "Invalid request body")
        return
    }
    userID, address, err := services.RegisterUser(req.Email, req.Password, req.Name)
    if err != nil {
        helpers.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }
    helpers.WriteJSON(w, http.StatusOK, map[string]string{
        "userId":  userID,
        "address": address,
        "status":  "success",
    })
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := helpers.ReadJSON(r, &req); err != nil {
        helpers.WriteError(w, http.StatusBadRequest, "Invalid request body")
        return
    }
    jwt, address, err := services.LoginUser(req.Email, req.Password)
    if err != nil {
        helpers.WriteError(w, http.StatusUnauthorized, err.Error())
        return
    }
    helpers.WriteJSON(w, http.StatusOK, map[string]string{
        "jwt":     jwt,
        "address": address,
        "status":  "success",
    })
}
```

### `event_controller.go`
**Purpose**: Handles multi-signature event endpoints (e.g., `/events/mint`, `/events/sign`).

```go
package controllers

import (
    "net/http"
    "github.com/gorilla/mux"
    "wowtoken-api/internal/helpers"
    "wowtoken-api/internal/services"
)

func RegisterEventRoutes(r *mux.Router) {
    r.HandleFunc("/events/mint", CreateMintEvent).Methods("POST")
    r.HandleFunc("/events/{eventId}", GetEvent).Methods("GET")
}

func CreateMintEvent(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Amount string `json:"amount"`
        JWT    string `json:"jwt"`
    }
    if err := helpers.ReadJSON(r, &req); err != nil {
        helpers.WriteError(w, http.StatusBadRequest, "Invalid request body")
        return
    }
    eventID, txHash, err := services.CreateMintEvent(req.Amount, req.JWT)
    if err != nil {
        helpers.WriteError(w, http.StatusForbidden, err.Error())
        return
    }
    helpers.WriteJSON(w, http.StatusOK, map[string]string{
        "eventId":        eventID,
        "transactionHash": txHash,
        "status":         "success",
    })
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    eventID := vars["eventId"]
    event, err := services.GetEvent(eventID)
    if err != nil {
        helpers.WriteError(w, http.StatusNotFound, err.Error())
        return
    }
    helpers.WriteJSON(w, http.StatusOK, event)
}
```

---

## 4. `internal/services/`
### `token_service.go`
**Purpose**: Implements business logic for token operations (e.g., balance queries, transfers).

```go
package services

import (
    "context"
    "errors"
    "wowtoken-api/internal/helpers"
    "wowtoken-api/pkg/db"
)

func GetTokenBalance(address string) (string, error) {
    if !helpers.IsValidEthAddress(address) {
        return "", errors.New("invalid Ethereum address")
    }
    balance, err := BlockchainService.GetBalance(context.Background(), address)
    if err != nil {
        return "", err
    }
    return balance.String(), nil
}

func TransferToken(from, to, amount, jwt string) (string, error) {
    userAddress, err := helpers.ValidateJWT(jwt)
    if err != nil || userAddress != from {
        return "", errors.New("invalid or unauthorized JWT")
    }
    if !helpers.IsValidEthAddress(to) {
        return "", errors.New("invalid recipient address")
    }
    privateKey, err := db.GetEncryptedPrivateKey(userAddress)
    if err != nil {
        return "", err
    }
    txHash, err := BlockchainService.TransferToken(from, to, amount, privateKey)
    if err != nil {
        return "", err
    }
    return txHash, nil
}
```

### `user_service.go`
**Purpose**: Manages user registration, login, and custodial wallet operations.

```go
package services

import (
    "errors"
    "github.com/ethereum/go-ethereum/accounts"
    "wowtoken-api/internal/helpers"
    "wowtoken-api/internal/models"
    "wowtoken-api/pkg/db"
)

func RegisterUser(email, password, name string) (string, string, error) {
    if db.UserExists(email) {
        return "", "", errors.New("email already registered")
    }
    wallet := accounts.NewWallet() // Simplified; use ethers.Wallet.createRandom equivalent
    address := wallet.Address.Hex()
    encryptedKey, err := helpers.EncryptPrivateKey(wallet.PrivateKey, "32-byte-encryption-key")
    if err != nil {
        return "", "", err
    }
    hashedPassword, _ := helpers.HashPassword(password)
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
        return "", "", err
    }
    return user.UserID, user.Address, nil
}

func LoginUser(email, password string) (string, string, error) {
    user, err := db.GetUserByEmail(email)
    if err != nil || !helpers.CheckPassword(password, user.HashedPassword) {
        return "", "", errors.New("invalid email or password")
    }
    jwt, err := helpers.GenerateJWT(user.Address)
    if err != nil {
        return "", "", err
    }
    return jwt, user.Address, nil
}
```

### `event_service.go`
**Purpose**: Manages multi-signature event logic (e.g., mint, event details).

```go
package services

import (
    "errors"
    "wowtoken-api/internal/helpers"
    "wowtoken-api/internal/models"
)

func CreateMintEvent(amount, jwt string) (string, string, error) {
    address, err := helpers.ValidateJWT(jwt)
    if err != nil {
        return "", "", errors.New("invalid JWT")
    }
    if !BlockchainService.IsCoOwner(address) {
        return "", "", errors.New("not a co-owner")
    }
    eventID := "mint_" + fmt.Sprintf("%d", time.Now().Unix())
    txHash, err := BlockchainService.CreateMintEvent(address, amount)
    if err != nil {
        return "", "", err
    }
    return eventID, txHash, nil
}

func GetEvent(eventID string) (models.Event, error) {
    event, err := BlockchainService.GetEvent(eventID)
    if err != nil {
        return models.Event{}, errors.New("event not found")
    }
    return event, nil
}
```

### `blockchain_service.go`
**Purpose**: Handles Ethereum blockchain interactions.

```go
package services

import (
    "context"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "wowtoken-api/internal/config"
)

type BlockchainService struct {
    client *ethclient.Client
}

func NewBlockchainService() (*BlockchainService, error) {
    cfg, _ := config.LoadConfig()
    client, err := ethclient.Dial(cfg.InfuraURL)
    if err != nil {
        return nil, err
    }
    return &BlockchainService{client: client}, nil
}

func (s *BlockchainService) GetBalance(ctx context.Context, address string) (*big.Int, error) {
    // Simplified; interact with WoWToken contract
    return big.NewInt(1000000000000000000), nil
}

func (s *BlockchainService) TransferToken(from, to, amount, privateKey string) (string, error) {
    // Simplified; sign and send transaction to WoWToken contract
    return "0x789...", nil
}

func (s *BlockchainService) IsCoOwner(address string) bool {
    // Simplified; query co_token_owner mapping
    return true
}

func (s *BlockchainService) CreateMintEvent(address, amount string) (string, error) {
    // Simplified; create multi-signature mint event
    return "0x789...", nil
}

func (s *BlockchainService) GetEvent(eventID string) (models.Event, error) {
    // Simplified; query event_requireMultiSignature
    return models.Event{
        IDEvent:       eventID,
        IsCompleted:   false,
        EventName:     "Mint Token",
        Description:   "Mint token",
        SignatureCount: 2,
        Signers:       []string{"0x123...", "0x456..."},
        CreatedAt:     1698765432,
    }, nil
}
```

---

## 5. `internal/models/`
### `token.go`
**Purpose**: Defines token-related data structures.

```go
package models

type BalanceResponse struct {
    Address string `json:"address"`
    Balance string `json:"balance"`
}

type TransferRequest struct {
    From   string `json:"from"`
    To     string `json:"to"`
    Amount string `json:"amount"`
    JWT    string `json:"jwt"`
}
```

### `user.go`
**Purpose**: Defines user-related data structures.

```go
package models

type User struct {
    UserID             string `bson:"userId" json:"userId"`
    Email              string `bson:"email" json:"email"`
    HashedPassword     string `bson:"hashedPassword" json:"-"`
    Name               string `bson:"name" json:"name"`
    Address            string `bson:"address" json:"address"`
    EncryptedPrivateKey string `bson:"encryptedPrivateKey" json:"-"`
    RegisteredAt       int64  `bson:"registeredAt" json:"registeredAt"`
}
```

### `event.go`
**Purpose**: Defines multi-signature event data structures.

```go
package models

type Event struct {
    IDEvent       string   `json:"id_event"`
    IsCompleted   bool     `json:"isCompleted"`
    EventName     string   `json:"event_name"`
    Description   string   `json:"description"`
    SignatureCount int      `json:"signature_count"`
    Signers       []string `json:"signers"`
    CreatedAt     int64    `json:"createdAt"`
}
```

---

## 6. `internal/helpers/`
### `jwt_helper.go`
**Purpose**: Manages JWT generation and validation.

```go
package helpers

import (
    "errors"
    "time"
    "github.com/dgrijalva/jwt-go"
    "wowtoken-api/internal/config"
)

func GenerateJWT(address string) (string, error) {
    cfg, _ := config.LoadConfig()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "address": address,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString([]byte(cfg.JWTSecret))
}

func ValidateJWT(tokenString string) (string, error) {
    cfg, _ := config.LoadConfig()
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(cfg.JWTSecret), nil
    })
    if err != nil || !token.Valid {
        return "", errors.New("invalid JWT")
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", errors.New("invalid claims")
    }
    return claims["address"].(string), nil
}
```

### `crypto_helper.go`
**Purpose**: Handles private key encryption/decryption.

```go
package helpers

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "errors"
    "io"
)

func EncryptPrivateKey(privateKey, key string) (string, error) {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }
    ciphertext := make([]byte, aes.BlockSize+len(privateKey))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(privateKey))
    return hex.EncodeToString(ciphertext), nil
}

func DecryptPrivateKey(encryptedKey, key string) (string, error) {
    ciphertext, _ := hex.DecodeString(encryptedKey)
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }
    if len(ciphertext) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]
    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)
    return string(ciphertext), nil
}
```

### `eth_helper.go`
**Purpose**: Provides Ethereum-specific utilities.

```go
package helpers

import (
    "regexp"
)

func IsValidEthAddress(address string) bool {
    re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
    return re.MatchString(address)
}
```

### `response_helper.go`
**Purpose**: Formats API responses and errors.

```go
package helpers

import (
    "encoding/json"
    "net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
    WriteJSON(w, status, map[string]string{"error": message})
}

func ReadJSON(r *http.Request, data interface{}) error {
    return json.NewDecoder(r.Body).Decode(data)
}
```

---

## 7. `internal/middleware/`
### `auth_middleware.go`
**Purpose**: Validates JWT for protected endpoints.

```go
package middleware

import (
    "net/http"
    "strings"
    "wowtoken-api/internal/helpers"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            helpers.WriteError(w, http.StatusUnauthorized, "Missing Authorization header")
            return
        }
        token := strings.TrimPrefix(authHeader, "Bearer ")
        if _, err := helpers.ValidateJWT(token); err != nil {
            helpers.WriteError(w, http.StatusUnauthorized, err.Error())
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

### `rate_limit.go`
**Purpose**: Implements rate-limiting using Redis.

```go
package middleware

import (
    "net/http"
    "github.com/go-redis/redis/v8"
    "wowtoken-api/pkg/db"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        client := db.GetRedisClient()
        key := r.RemoteAddr
        count, _ := client.Incr(r.Context(), key).Result()
        if count > 100 { // Example: 100 requests per minute
            helpers.WriteError(w, http.StatusTooManyRequests, "Rate limit exceeded")
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

---

## 8. `pkg/db/`
### `mongodb.go`
**Purpose**: Manages MongoDB connections and queries.

```go
package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "wowtoken-api/internal/config"
    "wowtoken-api/internal/models"
)

func ConnectMongoDB() (*mongo.Client, error) {
    cfg, _ := config.LoadConfig()
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoDBURI))
    if err != nil {
        return nil, err
    }
    return client, nil
}

func SaveUser(user models.User) error {
    client, _ := ConnectMongoDB()
    collection := client.Database("wowtoken").Collection("users")
    _, err := collection.InsertOne(context.Background(), user)
    return err
}

func GetUserByEmail(email string) (models.User, error) {
    client, _ := ConnectMongoDB()
    collection := client.Database("wowtoken").Collection("users")
    var user models.User
    err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
    return user, err
}

func UserExists(email string) bool {
    _, err := GetUserByEmail(email)
    return err == nil
}

func GetEncryptedPrivateKey(address string) (string, error) {
    client, _ := ConnectMongoDB()
    collection := client.Database("wowtoken").Collection("users")
    var user models.User
    err := collection.FindOne(context.Background(), bson.M{"address": address}).Decode(&user)
    if err != nil {
        return "", err
    }
    return user.EncryptedPrivateKey, nil
}
```

### `redis.go`
**Purpose**: Manages Redis connections for caching.

```go
package db

import (
    "context"
    "github.com/go-redis/redis/v8"
    "wowtoken-api/internal/config"
)

func GetRedisClient() *redis.Client {
    cfg, _ := config.LoadConfig()
    return redis.NewClient(&redis.Options{
        Addr: cfg.RedisAddr,
    })
}
```

---

## 9. `scripts/deploy.sh`
**Purpose**: Automates building and deploying the API.

```bash
#!/bin/bash
go build -o wowtoken-api ./cmd/api
docker build -t wowtoken-api .
docker run -d -p 8080:8080 wowtoken-api
```

---

## 10. `tests/token_controller_test.go`
**Purpose**: Tests the token controller endpoints.

```go
package controllers_test

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "wowtoken-api/internal/controllers"
)

func TestGetBalance(t *testing.T) {
    req, _ := http.NewRequest("GET", "/balances/0x1234567890123456789012345678901234567890", nil)
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.GetBalance)
    handler.ServeHTTP(rr, req)
    if rr.Code != http.StatusOK {
        t.Errorf("Expected status %v, got %v", http.StatusOK, rr.Code)
    }
}
```

---

## Dependencies
Add these to `go.mod`:

```go
module wowtoken-api

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    github.com/ethereum/go-ethereum v1.14.8
    go.mongodb.org/mongo-driver v1.12.0
    github.com/go-redis/redis/v8 v8.11.5
    github.com/dgrijalva/jwt-go v3.2.0+incompatible
    github.com/spf13/viper v1.16.0
)
```

---

## Notes
- **Blockchain Interactions**: The `blockchain_service.go` is simplified and assumes a `go-ethereum` client for interacting with the `WoWToken` contract. In a full implementation, you’d need to integrate with the contract’s ABI and methods (e.g., `transferToken`, `mintToken`).
- **Custodial Wallet**: The `user_service.go` simulates wallet creation; replace `accounts.NewWallet()` with `ethers.Wallet.createRandom()` from a library like `github.com/wealdtech/go-eth2-wallet`.
- **Security**: Private keys are encrypted using AES-256; consider using AWS KMS for production. JWTs are signed with HS256 for simplicity.
- **Testing**: Only one test file is shown; add tests for other controllers, services, and helpers.
- **Caching**: Redis is used for rate-limiting; extend it for caching balance or event data.
- **Error Handling**: The `response_helper.go` ensures consistent error responses as per the API spec.

This basic implementation provides a starting point for the WoWToken API, covering all required modules. Expand each module with additional endpoints, contract interactions, and error handling as needed.