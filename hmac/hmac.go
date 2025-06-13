package HMAC

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
)

func GenerateHMAC(data []byte, key []byte) []byte {
	mac := hmac.New(sha256.New, key) // 32 bytes
	mac.Write(data)
	return mac.Sum(nil)
}

func VerifyHMAC(data []byte, key []byte, hash []byte) bool {
	// Generate HMAC for the provided data using the same key
	generatedHMAC := GenerateHMAC(data, key)

	// Use subtle.ConstantTimeCompare to securely compare the two HMACs
	return subtle.ConstantTimeCompare(generatedHMAC, hash) == 1
}
