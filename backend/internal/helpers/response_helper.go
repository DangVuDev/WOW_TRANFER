package helpers

import (
	"encoding/json"
	"net/http"
    "wowtoken-api/internal/models"
)


  // WriteError ghi lỗi vào response
  func WriteError(w http.ResponseWriter, status int, message string) {
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(status)
      resp := models.ErrorResponse{Status: status, Message: message}
      if err := json.NewEncoder(w).Encode(resp); err != nil {
          http.Error(w, "Lỗi mã hóa JSON", http.StatusInternalServerError)
      }
  }

  // ReadJSON đọc JSON từ request
  func ReadJSON(r *http.Request, v interface{}) error {
      return json.NewDecoder(r.Body).Decode(v)
  }

  // WriteJSON ghi JSON vào response
  func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(status)
      if err := json.NewEncoder(w).Encode(v); err != nil {
          http.Error(w, "Lỗi mã hóa JSON", http.StatusInternalServerError)
      }
  }

 