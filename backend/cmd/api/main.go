package main

import (
	"log"
	"net/http"
	"wowtoken-api/internal/config"
	"wowtoken-api/internal/controllers"
	"wowtoken-api/internal/middleware"
	"wowtoken-api/internal/services"
    _ "wowtoken-api/cmd/api/docs"
	"github.com/gorilla/handlers" // ← Thêm import này
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Tải cấu hình
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Không thể tải cấu hình: %v", err)
	}

	// Khởi tạo BlockchainService
	blockchainService, err := services.NewBlockchainService()
	if err != nil {
		log.Fatalf("Không thể khởi tạo BlockchainService: %v", err)
	}

	// Khởi tạo TokenService và EventService
	tokenService := services.NewTokenService(blockchainService)
	eventService := services.NewEventService(blockchainService)

	// Đăng ký service cho controllers
	controllers.SetServices(tokenService, eventService)

	// Tạo router chính
	router := mux.NewRouter()

	// Định nghĩa route public (không cần xác thực)
	public := router.PathPrefix("/").Subrouter()
	public.HandleFunc("/users/register", controllers.RegisterUser).Methods("POST")
	public.HandleFunc("/users/login", controllers.LoginUser).Methods("POST")

	// Định nghĩa route protected (cần xác thực)
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.Use(middleware.RateLimitMiddleware) // Giả định đã triển khai
	controllers.RegisterEventRoutes(protected)
	controllers.RegisterTokenRoutes(protected)

	// Thêm Swagger UI (có thể để public hoặc bảo vệ tùy ý)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// ← Thêm CORS middleware ở đây (áp dụng cho toàn bộ router)
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Hoặc chỉ định origins cụ thể, ví dụ: []string{"http://localhost:3000"}
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"}),
		handlers.AllowCredentials(), // Nếu cần cookies/auth
	)(router)

	// Khởi động server với CORS handler
	log.Printf("Khởi động server tại :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, corsHandler); err != nil {
		log.Fatalf("Server thất bại: %v", err)
	}
}