package encryption

import (
	"fmt"
	"testing"
)

var data = "0chain.net rocks"
var expectedHash = "6cb51770083ba34e046bc6c953f9f05b64e16a0956d4e496758b97c9cf5687d5"

func TestHash(t *testing.T) {
	if Hash(data) != expectedHash {
		fmt.Printf("invalid hash\n")
	} else {
		fmt.Printf("hash successful\n")
	}
}

func TestGenerateKeys(t *testing.T) {
	publicKey, privateKey, err := GenerateKeys()
	fmt.Printf("keys: %v,%v, %v\n", privateKey, publicKey, err)
}

func BenchmarkGenerateKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateKeys()
	}
}

func TestSignAndVerify(t *testing.T) {
	publicKey, privateKey, err := GenerateKeys()
	signature, err := Sign(privateKey, expectedHash)
	if err != nil {
		fmt.Printf("error signing: %v\n", err)
		return
	}
	fmt.Printf("singing successful\n")
	if ok, err := Verify(publicKey, signature, expectedHash); err != nil || !ok {
		fmt.Printf("Verification failed\n")
	} else {
		fmt.Printf("Signing Verification successful\n")
	}
}

func BenchmarkSign(b *testing.B) {
	_, privateKey, err := GenerateKeys()
	if err == nil {
		return
	}
	for i := 0; i < b.N; i++ {
		Sign(privateKey, expectedHash)
	}
}

func BenchmarkVerify(b *testing.B) {
	publicKey, privateKey, err := GenerateKeys()
	signature, err := Sign(privateKey, expectedHash)
	if err != nil {
		return
	}
	for i := 0; i < b.N; i++ {
		Verify(publicKey, signature, expectedHash)
	}
}