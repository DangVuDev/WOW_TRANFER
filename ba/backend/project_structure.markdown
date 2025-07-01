# WoWToken Go Project Structure

This document outlines the folder structure and key components for implementing the WoWToken Web API in Go. The structure separates concerns into services, controllers, models, helpers, and configuration, following Go best practices for a RESTful API integrated with an Ethereum blockchain smart contract.

## Folder Structure

```
wowtoken-api/
├── cmd/
│   └── api/
│       └── main.go                 # Entry point for the API server
├── internal/
│   ├── config/
│   │   ├── config.go              # Configuration loading and environment variables
│   │   └── config.yaml            # Configuration file (e.g., API keys, blockchain provider)
│   ├── controllers/
│   │   ├── token_controller.go    # Handlers for token-related endpoints (e.g., /balances, /transfer)
│   │   ├── user_controller.go     # Handlers for user management endpoints (e.g., /users/register, /users/login)
│   │   └── event_controller.go    # Handlers for multi-signature event endpoints (e.g., /events/mint, /events/sign)
│   ├── services/
│   │   ├── token_service.go       # Business logic for token operations (interacts with smart contract)
│   │   ├── user_service.go        # Business logic for user management (custodial wallet, authentication)
│   │   ├── event_service.go       # Business logic for multi-signature events
│   │   └── blockchain_service.go   # Blockchain interaction logic (e.g., ethers.go, Infura/Alchemy)
│   ├── models/
│   │   ├── token.go               # Data structures for token-related data (e.g., balance, transfer)
│   │   ├── user.go                # Data structures for user data (e.g., email, address, encrypted key)
│   │   └── event.go               # Data structures for multi-signature events
│   ├── helpers/
│   │   ├── jwt_helper.go          # JWT generation and validation
│   │   ├── crypto_helper.go       # Encryption/decryption for private keys (e.g., AES-256)
│   │   ├── eth_helper.go          # Ethereum-specific utilities (e.g., address validation)
│   │   └── response_helper.go     # Standard API response formatting
│   └── middleware/
│       ├── auth_middleware.go     # Middleware for JWT authentication
│       └── rate_limit.go          # Rate-limiting middleware
├── pkg/
│   └── db/
│       ├── mongodb.go             # MongoDB connection and queries
│       └── redis.go               # Redis connection for caching
├── scripts/
│   └── deploy.sh                  # Deployment script for API
├── tests/
│   ├── controllers/               # Unit tests for controllers
│   ├── services/                  # Unit tests for services
│   └── helpers/                   # Unit tests for helpers
├── go.mod                         # Go module dependencies
├── go.sum                         # Dependency checksums
└── README.md                      # Project overview and setup instructions
```

## Component Details

### 1. `cmd/api/main.go`
- **Purpose**: Entry point for the API server.
- **Responsibilities**:
  - Initialize configuration, database, and blockchain provider.
  - Set up HTTP router (using `gorilla/mux` or `chi`).
  - Register routes with controllers and middleware.
  - Start the HTTP server.
- **Example**:
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

      // Register controllers
      controllers.RegisterTokenRoutes(router)
      controllers.RegisterUserRoutes(router)
      controllers.RegisterEventRoutes(router)

      log.Printf("Starting server on port %s", cfg.Port)
      http.ListenAndServe(":"+cfg.Port, router)
  }
  ```

### 2. `internal/config/`
- **Purpose**: Manage configuration loading from environment variables or `config.yaml`.
- **Files**:
  - `config.go`: Defines configuration struct and loading logic.
  - `config.yaml`: Stores API keys, blockchain provider URLs, database credentials, etc.
- **Example (`config.go`)**:
  ```go
  package config

  import (
      "github.com/spf13/viper"
  )

  type Config struct {
      Port           string `mapstructure:"PORT"`
      InfuraURL      string `mapstructure:"INFURA_URL"`
      MongoDBURI     string `mapstructure:"MONGODB_URI"`
      RedisAddr      string `mapstructure:"REDIS_ADDR"`
      JWTSecret      string `mapstructure:"JWT_SECRET"`
      EncryptionKey  string `mapstructure:"ENCRYPTION_KEY"` 
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
- **Example (`config.yaml`)**:
  ```yaml
  PORT: "8080"
  INFURA_URL: "https://mainnet.infura.io/v3/your-infura-key"
  MONGODB_URI: "mongodb://localhost:27017/wowtoken"
  REDIS_ADDR: "localhost:6379"
  JWT_SECRET: "your-jwt-secret"
  ENCRYPTION_KEY: "32-byte-encryption-key"
  ```

### 3. `internal/controllers/`
- **Purpose**: Handle HTTP requests and responses, invoking services for business logic.
- **Files**:
  - `token_controller.go`: Endpoints like `/balances/{address}`, `/transfer`, `/approve`.
  - `user_controller.go`: Endpoints like `/users/register`, `/users/login`, `/users/{address}`.
  - `event_controller.go`: Endpoints like `/events/mint`, `/events/sign`, `/events/execute/{eventType}`.
- **Example (`token_controller.go`)**:
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
      address := mux.Vars(r)["address"]
      balance, err := services.GetTokenBalance(address)
      if err != nil {
          helpers.WriteError(w, http.StatusBadRequest, err.Error())
          return
      }
      helpers.WriteJSON(w, http.StatusOK, map[string]string{"address": address, "balance": balance})
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
      helpers.WriteJSON(w, http.StatusOK, map[string]string{"transactionHash": txHash, "status": "success"})
  }
  ```

### 4. `internal/services/`
- **Purpose**: Contain business logic, interacting with the blockchain, database, and helpers.
- **Files**:
  - `token_service.go`: Logic for token operations (e.g., balance queries, transfers).
  - `user_service.go`: Logic for user registration, login, and custodial wallet management.
  - `event_service.go`: Logic for multi-signature events (e.g., mint, burn, pause).
  - `blockchain_service.go`: Blockchain interactions using `ethers-go` or equivalent.
- **Example (`token_service.go`)**:
  ```go
  package services

  import (
      "context"
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

### 5. `internal/models/`
- **Purpose**: Define data structures for API requests, responses, and database entities.
- **Files**:
  - `token.go`: Structs for balance, transfer, approval.
  - `user.go`: Structs for user data (email, address, encrypted private key).
  - `event.go`: Structs for multi-signature events.
- **Example (`user.go`)**:
  ```go
  package models

  type User struct {
      UserID             string `bson:"userId"`
      Email              string `bson:"email"`
      HashedPassword     string `bson:"hashedPassword"`
      Name               string `bson:"name"`
      Address            string `bson:"address"`
      EncryptedPrivateKey string `bson:"encryptedPrivateKey"`
      RegisteredAt       int64  `bson:"registeredAt"`
  }
  ```

### 6. `internal/helpers/`
- **Purpose**: Utility functions for common tasks.
- **Files**:
  - `jwt_helper.go`: JWT generation and validation.
  - `crypto_helper.go`: Private key encryption/decryption (e.g., AES-256).
  - `eth_helper.go`: Ethereum address validation and utilities.
  - `response_helper.go`: Standard JSON response and error handling.
- **Example (`jwt_helper.go`)**:
  ```go
  package helpers

  import (
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

### 7. `internal/middleware/`
- **Purpose**: Handle cross-cutting concerns like authentication and rate-limiting.
- **Files**:
  - `auth_middleware.go`: Validates JWT for protected endpoints.
  - `rate_limit.go`: Implements rate-limiting using Redis.
- **Example (`auth_middleware.go`)**:
  ```go
  package middleware

  import (
      "net/http"
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

### 8. `pkg/db/`
- **Purpose**: Database connection and query logic.
- **Files**:
  - `mongodb.go`: MongoDB connection and CRUD operations.
  - `redis.go`: Redis connection for caching.
- **Example (`mongodb.go`)**:
  ```go
  package db

  import (
      "context"
      "go.mongodb.org/mongo-driver/mongo"
      "wowtoken-api/internal/config"
  )

  func ConnectMongoDB() (*mongo.Client, error) {
      cfg, _ := config.LoadConfig()
      client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoDBURI))
      if err != nil {
          return nil, err
      }
      return client, nil
  }

  func GetEncryptedPrivateKey(address string) (string, error) {
      // Query MongoDB for user by address
      // Return encrypted private key
  }
  ```

### 9. `scripts/`
- **Purpose**: Deployment and utility scripts.
- **Files**:
  - `deploy.sh`: Script to build and deploy the API to a server or container.
- **Example (`deploy.sh`)**:
  ```bash
  #!/bin/bash
  go build -o wowtoken-api ./cmd/api
  docker build -t wowtoken-api .
  docker run -p 8080:8080 wowtoken-api
  ```

### 10. `tests/`
- **Purpose**: Unit and integration tests for controllers, services, and helpers.
- **Structure**:
  - Mirrors the `internal/` directory with test files (e.g., `token_controller_test.go`).
- **Example (`token_controller_test.go`)**:
  ```go
  package controllers_test

  import (
      "net/http"
      "net/http/httptest"
      "testing"
      "wowtoken-api/internal/controllers"
  )

  func TestGetBalance(t *testing.T) {
      req, _ := http.NewRequest("GET", "/balances/0x123...", nil)
      rr := httptest.NewRecorder()
      handler := http.HandlerFunc(controllers.GetBalance)
      handler.ServeHTTP(rr, req)
      if status := rr.Code; status != http.StatusOK {
          t.Errorf("Expected status %v, got %v", http.StatusOK, status)
      }
  }
  ```

## Dependencies
- **HTTP Router**: `github.com/gorilla/mux` or `github.com/go-chi/chi`
- **Blockchain**: `github.com/ethereum/go-ethereum` (ethers-go equivalent)
- **Database**: `go.mongodb.org/mongo-driver/mongo`, `github.com/redis/go-redis`
- **JWT**: `github.com/dgrijalva/jwt-go`
- **Config**: `github.com/spf13/viper`
- **Encryption**: Standard Go crypto packages (e.g., `crypto/aes`)
- **Logging**: `github.com/sirupsen/logrus`

## Notes
- **Modularity**: The `internal/` directory ensures private packages, preventing external imports.
- **Security**: Private keys are encrypted using AES-256 or AWS KMS, stored in MongoDB.
- **Caching**: Redis is used for read-heavy endpoints like `/balances/{address}`.
- **Testing**: Unit tests cover controllers, services, and helpers; integration tests verify blockchain and database interactions.
- **Blockchain**: The `blockchain_service.go` uses `go-ethereum` to interact with the `WoWToken` smart contract via Infura/Alchemy.
- **Scalability**: Rate-limiting and caching are implemented to handle high traffic.
- **Custodial Wallet**: User private keys are managed securely by the backend, with `user_service.go` handling wallet creation and transaction signing.

This structure provides a robust foundation for the WoWToken API, balancing simplicity for non-technical users with secure blockchain interactions.