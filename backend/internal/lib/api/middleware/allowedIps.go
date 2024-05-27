package middleware

import (
	"github.com/spf13/viper"
)

func AllowedIps(ip string) bool {
	allowedIps := map[string]string{
		"185.71.76.0/27":   "allowed",
		"185.71.77.0/27":   "allowed",
		"77.75.153.0/25":   "allowed",
		"77.75.156.11":     "allowed",
		"77.75.156.35":     "allowed",
		"77.75.154.128/25": "allowed",
		"2a02:5180::/32":   "allowed",
	}

	mode := viper.GetString("mode")

	if _, ok := allowedIps[ip]; !ok && mode == "release" {
		return false
	}

	return true
}
