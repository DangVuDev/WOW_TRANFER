package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"wowtoken-api/internal/helpers"
	"wowtoken-api/internal/services"
)

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
type ApproveRequest struct {
    Owner   string `json:"owner"`
    Spender string `json:"spender"`
    Amount  string `json:"amount"`
    JWT     string `json:"jwt"`
}
type BurnRequest struct {
    Address string `json:"address"`
    Amount  string `json:"amount"`
    JWT     string `json:"jwt"`
}

var tokenService *services.TokenService

// SetServices thiết lập các service cho controller
func SetServices(ts *services.TokenService, es *services.EventService) {
	tokenService = ts
	eventService = es
}

// RegisterTokenRoutes đăng ký các route liên quan đến token
func RegisterTokenRoutes(r *mux.Router) {
	r.HandleFunc("/balances/{address}", GetBalance).Methods("GET")
	r.HandleFunc("/transfer", TransferToken).Methods("POST")
	r.HandleFunc("/approve", ApproveToken).Methods("POST")
	r.HandleFunc("/burn", BurnToken).Methods("POST")
}

// @Summary Lấy số dư token
// @Description Lấy số dư token của một địa chỉ Ethereum
// @Tags tokens
// @Accept json
// @Produce json
// @Param address path string true "Địa chỉ Ethereum"
// @Success 200 {object} BalanceResponse
// @Failure 400 {object} ErrorResponse "Yêu cầu không hợp lệ"
// @Router /balances/{address} [get]
func GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	balance, err := tokenService.GetTokenBalance(address)
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.WriteJSON(w, http.StatusOK, BalanceResponse{
		Address: address,
		Balance: balance,
	})
}

// @Summary Chuyển token
// @Description Thực hiện chuyển token từ địa chỉ này sang địa chỉ khác
// @Tags tokens
// @Accept json
// @Produce json
// @Param request body TransferRequest true "Thông tin chuyển token"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse "Yêu cầu không hợp lệ"
// @Router /transfer [post]
func TransferToken(w http.ResponseWriter, r *http.Request) {
	var req TransferRequest
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Yêu cầu không hợp lệ")
		return
	}
	txHash, err := tokenService.TransferToken(req.From, req.To, req.Amount, req.JWT)
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.WriteJSON(w, http.StatusOK, map[string]string{
		"transactionHash": txHash,
		"status":         "success",
	})
}

// @Summary Phê duyệt token
// @Description Phê duyệt số lượng token cho địa chỉ chi tiêu
// @Tags tokens
// @Accept json
// @Produce json
// @Param request body ApproveRequest true "Thông tin phê duyệt token"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse "Yêu cầu không hợp lệ"
// @Router /approve [post]
func ApproveToken(w http.ResponseWriter, r *http.Request) {
	var req ApproveRequest
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Yêu cầu không hợp lệ")
		return
	}
	txHash, err := tokenService.ApproveToken(req.Owner, req.Spender, req.Amount, req.JWT)
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.WriteJSON(w, http.StatusOK, map[string]string{
		"transactionHash": txHash,
		"status":         "success",
	})
}

// @Summary Đốt token
// @Description Đốt một số lượng token từ địa chỉ chỉ định
// @Tags tokens
// @Accept json
// @Produce json
// @Param request body BurnRequest true "Thông tin đốt token"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse "Yêu cầu không hợp lệ"
// @Router /burn [post]
func BurnToken(w http.ResponseWriter, r *http.Request) {
	var req BurnRequest
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Yêu cầu không hợp lệ")
		return
	}
	txHash, err := tokenService.BurnToken(req.Address, req.Amount, req.JWT)
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.WriteJSON(w, http.StatusOK, map[string]string{
		"transactionHash": txHash,
		"status":         "success",
	})
}