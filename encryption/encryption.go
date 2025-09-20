package encryption

import (
	"bytecaster/cli"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rc4"
	"crypto/sha256"
	"io"
	"log"
)

type encryptor struct {
	input  []byte
	output []byte
	key    []byte
}

func EncryptData(data []byte, algorithm string, key string) []byte {
	enc := encryptor{
		input: data,
		key:   []byte(key),
	}

	switch algorithm {
	case cli.OptEncryptionXOR:
		enc.xor()
	case cli.OptEncryptionAES256:
		enc.aes256()
	case cli.OptEncryptionRC4:
		enc.rc4()
	default:
		log.Fatal("Unknown encryption algorithm")
	}

	return enc.output
}

func (e *encryptor) xor() {
	keyBytes := []byte(e.key)
	e.output = make([]byte, len(e.input))

	for i := 0; i < len(e.input); i++ {
		e.output[i] = e.input[i] ^ keyBytes[i%len(keyBytes)]
	}
}

func (e *encryptor) aes256() {
	// Key derivation (32-bits)
	derivedKey := sha256.Sum256(e.key)

	block, err := aes.NewCipher(derivedKey[:])
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)

	}

	ciphertext := gcm.Seal(nil, nonce, e.input, nil)

	e.output = append(nonce, ciphertext...)
}

func (e *encryptor) rc4() {
	c, err := rc4.NewCipher(e.key)
	if err != nil {
		panic(err)
	}

	e.output = make([]byte, len(e.input))

	c.XORKeyStream(e.output, e.input)
}
