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
        return "", errors.New("JWT không hợp lệ")
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", errors.New("claims không hợp lệ")
    }
    return claims["address"].(string), nil
}