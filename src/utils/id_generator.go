package utils

import "time"

func GenerateId() int {
	return int(time.Now().UnixNano())
}
