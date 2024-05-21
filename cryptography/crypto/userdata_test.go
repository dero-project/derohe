package crypto

import (
	"math/big"
	"testing"
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

// TODO FIXME
func TestSharedKey(t *testing.T) {
	r := RandomScalar()
	private_key := RandomScalarBNRed()
	public_key := GPoint.ScalarMult(private_key)

	encrypted_key := GenerateSharedSecret(r, public_key.G1())

	big_key := new(big.Int).SetBytes(encrypted_key[:])
	decrypted_r := new(big.Int).Mul(private_key.BigInt(), big_key)

	r_bytes := r.Bytes()
	decrypted_r_bytes := decrypted_r.Bytes()
	for i := range r_bytes {
		if r_bytes[i] != decrypted_r_bytes[i] {
			t.Fatalf("error on shared key, index: %d", i)
		}
	}
}
