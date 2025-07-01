package controllers

import (
	"net/http"
	"wowtoken-api/internal/helpers"
	"wowtoken-api/internal/services"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes đăng ký các route liên quan đến người dùng
func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", LoginUser).Methods("POST")
}

// @Summary Đăng ký người dùng mới
// @Description Tạo tài khoản mới với email, mật khẩu và tên
// @Tags users
// @Accept json
// @Produce json
// @Param request body RegisterUserRequest true "Thông tin đăng ký"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse "Yêu cầu không hợp lệ"
// @Router /users/register [post]
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Yêu cầu không hợp lệ")
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

// @Summary Đăng nhập người dùng
// @Description Xác thực người dùng với email và mật khẩu để lấy JWT
// @Tags users
// @Accept json
// @Produce json
// @Param request body LoginUserRequest true "Thông tin đăng nhập"
// @Success 200 {object} map[string]string
// @Failure 401 {object} ErrorResponse "Xác thực không thành công"
// @Router /users/login [post]
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Yêu cầu không hợp lệ")
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

// Định nghĩa các struct cho Swagger (nếu cần)
type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}