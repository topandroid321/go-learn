package Functions

import (
	"strconv"

	"github.com/gofiber/fiber/v2/log"
)

func StoI(s string) int {
	result, error := strconv.Atoi(s)
	if error != nil {
		log.Error(error)
		return -0
	}
	return result
}

func IsEmpty(value any, defaultVal any) any {
	if value == nil || value == "" {
		return defaultVal
	}
	return value
}
