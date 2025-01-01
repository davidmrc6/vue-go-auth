package utils

import "os"

func GetJWTKey() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}
