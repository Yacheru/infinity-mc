package websocket

func CutMessagePrefix(bytes []byte) []byte {
	for i, b := range bytes {
		if b == '{' {
			return bytes[i:]
		}
	}

	return nil
}
