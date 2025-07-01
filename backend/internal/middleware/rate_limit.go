package middleware

import (
    "net/http"
    "wowtoken-api/internal/helpers"
    "wowtoken-api/pkg/db"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        client := db.GetRedisClient()
        key := r.RemoteAddr
        count, _ := client.Incr(r.Context(), key).Result()
        if count > 100 { // Giới hạn 100 yêu cầu/phút
            helpers.WriteError(w, http.StatusTooManyRequests, "Vượt quá giới hạn tỷ lệ")
            return
        }
        next.ServeHTTP(w, r)
    })
}