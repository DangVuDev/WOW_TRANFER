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
            helpers.WriteError(w, http.StatusUnauthorized, "Thiếu header Authorization")
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