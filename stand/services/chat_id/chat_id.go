package chatid

//      generates a unique chat ID based on
// two user IDs using the Cantor pairing function.

func BuildChatID(x, y int) int64 {
	return int64((x + y) * (x + y + 1) / 2 + y)
}