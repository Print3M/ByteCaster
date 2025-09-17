package encoding

import (
	"bin2code/cli"
	"encoding/base64"
	"log"
)

type encoder struct {
	input  []byte
	output []byte
}

func EncodeData(data []byte, encoding string) []byte {
	enc := encoder{
		input: data,
	}

	switch encoding {
	case cli.OptEncodingBase64:
		enc.base64()
	default:
		log.Fatal("Unknown encoding")
	}

	return enc.output
}

func (e *encoder) base64() {
	encoded := base64.StdEncoding.EncodeToString(e.input)
	e.output = []byte(encoded)
}

// MAC address encoding
// IP address encoding
// ASM code encoding
