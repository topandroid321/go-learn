package Utils

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
