package encryption

import (
	"bin2code/cli"
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
	case cli.OptEncryptionXor:
		enc.xor()
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
