package controllers

import (
	"net/http"
	"wowtoken-api/internal/helpers"
	"wowtoken-api/internal/services"

	"github.com/gorilla/mux"
)
type ErrorResponse struct {
      Status  int    `json:"status"`
      Message string `json:"message"`
  }

var eventService *services.EventService

// RegisterEventRoutes đăng ký các route liên quan đến sự kiện
func RegisterEventRoutes(r *mux.Router) {
	r.HandleFunc("/events/mint", CreateMintEvent).Methods("POST")
	r.HandleFunc("/events/sign", SignEvent).Methods("POST")
	r.HandleFunc("/events/add-co-owner", CreateAddCoOwnerEvent).Methods("POST")
	r.HandleFunc("/events/burn", CreateBurnTokenEvent).Methods("POST")
	r.HandleFunc("/events/pause", CreatePauseEvent).Methods("POST")
	r.HandleFunc("/events/unpause", CreateUnpauseEvent).Methods("POST")
}

// @Summary Tạo sự kiện mint token
// @Description Tạo một sự kiện mint token với số lượng chỉ định
// @Tags events
// @Accept json
// @Produce json
// @Param request body CreateMintEventRequest true "Thông tin yêu cầu mint"
// @Success 200 {object} CreateMintEventResponse
// @Failure 400 {object} ErrorResponse "Yêu cầu không hợp lệ"
// @Failure 403 {object} ErrorResponse "Lỗi xử lý sự kiện"
// @Router /events/mint [post]
func CreateMintEvent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Amount string `json:"amount"`
		JWT    string `json:"jwt"`
	}
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Yêu cầu không hợp lệ")
		return
	}
	eventID, txHash, err := eventService.CreateMintEvent(req.Amount, req.JWT)
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

// @Summary Ký sự kiện
// @Description Ký hoặc từ chối một sự kiện hiện có
// @Tags events
// @Accept json
// @Produce json
// @Param request body SignEventRequest true "Thông tin yêu cầu ký"
// @Success 200 {object} SignEventResponse
// @Failure 400 {object} ErrorResponse "Yêu cầu không hợp lệ"
// @Failure 403 {object} ErrorResponse "Lỗi xử lý sự kiện"
// @Router /events/sign [post]
func SignEvent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		EventID string `json:"eventId"`
		Approve bool   `json:"approve"`
		JWT     string `json:"jwt"`
	}
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Yêu cầu không hợp lệ")
		return
	}
	txHash, err := eventService.SignEvent(req.EventID, req.JWT, req.Approve)
	if err != nil {
		helpers.WriteError(w, http.StatusForbidden, err.Error())
		return
	}
	helpers.WriteJSON(w, http.StatusOK, map[string]string{
		"transactionHash": txHash,
		"status":         "success",
	})
}

// @Summary Tạo sự kiện thêm đồng sở hữu
// @Description Tạo sự kiện để thêm một đồng sở hữu mới
// @Tags events
// @Accept json
// @Produce json
// @Param request body CreateAddCoOwnerEventRequest true "Thông tin yêu cầu thêm đồng sở hữu"
// @Success 200 {object} CreateAddCoOwnerEventResponse
// @Failure 400 {object} ErrorResponse "Yêu cầu không hợp lệ"
// @Failure 403 {object} ErrorResponse "Lỗi xử lý sự kiện"
// @Router /events/add-co-owner [post]
func CreateAddCoOwnerEvent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		NewCoOwner string `json:"newCoOwner"`
		JWT        string `json:"jwt"`
	}
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Yêu cầu không hợp lệ")
		return
	}
	eventID, txHash, err := eventService.CreateAddCoOwnerEvent(req.NewCoOwner, req.JWT)
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

// @Summary Tạo sự kiện đốt token
// @Description Tạo sự kiện để đốt token với số lượng chỉ định
// @Tags events
// @Accept json
// @Produce json
// @Param request body CreateBurnTokenEventRequest true "Thông tin yêu cầu đốt token"
// @Success 200 {object} CreateBurnTokenEventResponse
// @Failure 400 {object} ErrorResponse "Yêu cầu không hợp lệ"
// @Failure 403 {object} ErrorResponse "Lỗi xử lý sự kiện"
// @Router /events/burn [post]
func CreateBurnTokenEvent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Amount string `json:"amount"`
		JWT    string `json:"jwt"`
	}
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Yêu cầu không hợp lệ")
		return
	}
	eventID, txHash, err := eventService.CreateBurnTokenEvent(req.Amount, req.JWT)
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

// @Summary Tạo sự kiện tạm dừng
// @Description Tạo sự kiện để tạm dừng hoạt động hợp đồng
// @Tags events
// @Accept json
// @Produce json
// @Param request body CreatePauseEventRequest true "Thông tin yêu cầu tạm dừng"
// @Success 200 {object} CreatePauseEventResponse
// @Failure 400 {object} ErrorResponse "Yêu cầu không hợp lệ"
// @Failure 403 {object} ErrorResponse "Lỗi xử lý sự kiện"
// @Router /events/pause [post]
func CreatePauseEvent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		JWT string `json:"jwt"`
	}
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Yêu cầu không hợp lệ")
		return
	}
	eventID, txHash, err := eventService.CreatePauseEvent(req.JWT)
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

// @Summary Tạo sự kiện bỏ tạm dừng
// @Description Tạo sự kiện để bỏ tạm dừng hoạt động hợp đồng
// @Tags events
// @Accept json
// @Produce json
// @Param request body CreateUnpauseEventRequest true "Thông tin yêu cầu bỏ tạm dừng"
// @Success 200 {object} CreateUnpauseEventResponse
// @Failure 400 {object} ErrorResponse "Yêu cầu không hợp lệ"
// @Failure 403 {object} ErrorResponse "Lỗi xử lý sự kiện"
// @Router /events/unpause [post]
func CreateUnpauseEvent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		JWT string `json:"jwt"`
	}
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Yêu cầu không hợp lệ")
		return
	}
	eventID, txHash, err := eventService.CreateUnpauseEvent(req.JWT)
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

// Định nghĩa các struct cho Swagger (nếu cần)
type CreateMintEventRequest struct {
	Amount string `json:"amount"`
	JWT    string `json:"jwt"`
}

type CreateMintEventResponse struct {
	EventId        string `json:"eventId"`
	TransactionHash string `json:"transactionHash"`
	Status         string `json:"status"`
}

type SignEventRequest struct {
	EventID string `json:"eventId"`
	Approve bool   `json:"approve"`
	JWT     string `json:"jwt"`
}

type SignEventResponse struct {
	TransactionHash string `json:"transactionHash"`
	Status         string `json:"status"`
}

type CreateAddCoOwnerEventRequest struct {
	NewCoOwner string `json:"newCoOwner"`
	JWT        string `json:"jwt"`
}

type CreateAddCoOwnerEventResponse struct {
	EventId        string `json:"eventId"`
	TransactionHash string `json:"transactionHash"`
	Status         string `json:"status"`
}

type CreateBurnTokenEventRequest struct {
	Amount string `json:"amount"`
	JWT    string `json:"jwt"`
}

type CreateBurnTokenEventResponse struct {
	EventId        string `json:"eventId"`
	TransactionHash string `json:"transactionHash"`
	Status         string `json:"status"`
}

type CreatePauseEventRequest struct {
	JWT string `json:"jwt"`
}

type CreatePauseEventResponse struct {
	EventId        string `json:"eventId"`
	TransactionHash string `json:"transactionHash"`
	Status         string `json:"status"`
}

type CreateUnpauseEventRequest struct {
	JWT string `json:"jwt"`
}

type CreateUnpauseEventResponse struct {
	EventId        string `json:"eventId"`
	TransactionHash string `json:"transactionHash"`
	Status         string `json:"status"`
}