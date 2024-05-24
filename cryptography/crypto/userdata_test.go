package crypto

import (
	"testing"

	"github.com/deroproject/derohe/cryptography/bn256"
)

func TestEncryptDecrypt(t *testing.T) {
	var key [32]byte
	expected := []byte("Hello World")
	var payload []byte
	payload = append(payload, expected...)

	EncryptDecryptUserData(key, payload)
	EncryptDecryptUserData(key, payload)

	for i := range expected {
		if expected[i] != payload[i] {
			t.Fatal("error on encrypt/decrypt")
		}
	}
}

func TestSharedKey(t *testing.T) {
	r := RandomScalarBNRed()
	private_key := RandomScalarBNRed()
	public_key := GPoint.ScalarMult(private_key)

	// Create a curve point using PK
	encrypted_key := new(bn256.G1).ScalarMult(public_key.G1(), r.BigInt())

	// Decrypt it
	decrypted_r := DeriveKeyFromPoint(encrypted_key, private_key.BigInt())

	r_bytes := DeriveKeyFromR(r)
	for i := range r_bytes {
		if r_bytes[i] != decrypted_r[i] {
			t.Fatalf("error on shared key, index: %d", i)
		}
	}
}
