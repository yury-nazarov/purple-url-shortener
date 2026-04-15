package hash

import (
	"crypto/sha256"
	"fmt"
)

func Gen(userEmail string) (string, error) {
	h := sha256.New()
	h.Write([]byte(userEmail))
	hashString := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("INFO: String Sha256 success generate", hashString)
	return hashString, nil
}
