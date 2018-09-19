package userInfo

import (
	"golang.org/x/crypto/scrypt"
	"github.com/yushuailiu/MarsBase/pkg/config"
	"encoding/base64"
)

func HashPassword(password string) (string, error) {
	salt := config.GetConfig().Section("").Key("encryptKey").String()

	hash, err := scrypt.Key([]byte(password), []byte(salt), 1<<10, 8, 1, 64)

	return base64.StdEncoding.EncodeToString(hash), err
}
