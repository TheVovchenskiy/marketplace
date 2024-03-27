package hash

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashPassword(password string, salt string) (hashedPassword string) {
	passwordBytes := []byte(password)

	sha512Hasher := sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)

	hashedPassword = hex.EncodeToString(sha512Hasher.Sum(nil))

	return
}

func MatchPasswords(hashedPassword, currPassword string, salt string) (ok bool) {
	return hashedPassword == HashPassword(currPassword, salt)
}
