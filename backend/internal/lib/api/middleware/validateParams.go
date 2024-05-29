package middleware

func ValidateParams(nickname, price, email, donat, duration string) (string, string, string, string, string, bool) {
	if nickname != "" && price != "" && email != "" && donat != "" && duration != "" {
		return nickname, price, email, donat, duration, true
	}

	return "", "", "", "", "", false
}
