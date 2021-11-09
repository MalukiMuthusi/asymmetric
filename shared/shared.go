package shared

import (
	"log"
	"strconv"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Invalid number entered")
	}

	return i
}
