package chatid

import (
	regex "services/regex"
	"strconv"
)

func BuildChatID(xID string, yID string) (string, error) {
	ok := regex.IsStringInt(xID)
	if !ok {
		return "", nil
	}
	ok = regex.IsStringInt(yID)
	if !ok {
		return "", nil
	}

	x, _ := strconv.Atoi(xID)
	y, _ := strconv.Atoi(yID)

	chatID := cantorPairing(x, y)
	
	return strconv.Itoa(chatID), nil
}
func cantorPairing(x, y int) int {
	return (x + y) * (x + y + 1) / 2 + y
}