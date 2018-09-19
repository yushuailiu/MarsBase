package userInfo

import "regexp"

var emailReg = "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
var phoneReg = "^1[\\d]{10}$"
var usernameReg = "^[\\w]{2,20}$"
var passwordReg = "^([A-Z]|[a-z]|[0-9]|[-=[;,./~!@#$%^*()_+}{:?]){6,20}$"
var nicknameReg = `^[\w\x{4e00}-\x{9fa5}]{2,20}$`

func IsValidUsername(username string) bool {
	match, _ := regexp.MatchString(usernameReg, username)
	return match
}

func IsValidEmail(email string) bool {
	match, _ := regexp.MatchString(emailReg, email)
	return match
}

func IsValidNickname(nickname string) bool {
	match, _ := regexp.MatchString(nicknameReg, nickname)
	return match
}

func IsValidPhone(phone string) bool {
	match, _ := regexp.MatchString(phoneReg, phone)
	return match
}

func IsValidPassword(password string) bool {
	match, _ := regexp.MatchString(passwordReg, password)
	return match
}

func LoginType(identifier string) string {
	match, _ := regexp.MatchString(emailReg, identifier)
	if match {
		return "emial";
	}

	match, _ = regexp.MatchString(phoneReg, identifier)
	if match {
		return "phone"
	}

	match,_ = regexp.MatchString(usernameReg, identifier)
	if match {
		return "username"
	}
	return ""
}