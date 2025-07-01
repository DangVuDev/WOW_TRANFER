package helpers

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
    return uuid.New().String()
}
// IsValidEmail kiểm tra email có hợp lệ không
func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	// Biểu thức chính quy cho email
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(strings.TrimSpace(email))
}