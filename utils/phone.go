package utils

import "regexp"

func IsValidPhoneNumber(phoneNumber string) bool {
	pattern := `^0\d{9}$`
	match, _ := regexp.MatchString(pattern, phoneNumber)
	return match
}
