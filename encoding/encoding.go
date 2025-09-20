package encoding

import (
	"bytecaster/cli"
	"encoding/base32"
	"encoding/base64"
	"log"
	"net"
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
	case cli.OptEncodingBase32:
		enc.base32()
	case cli.OptEncodingBase64:
		enc.base64()
	case cli.OptEncodingIPv4:
		enc.ipv4()
	case cli.OptEncodingMAC:
		enc.mac()
	default:
		log.Fatal("Unknown encoding")
	}

	return enc.output
}

func (e *encoder) base32() {
	/*
		IyBCeXRlQ2FzdGVyCgpTd2lzcy1rbmlmZSBmaWxlI...
	*/
	encoded := base32.StdEncoding.EncodeToString(e.input)
	e.output = []byte(encoded)
}

func (e *encoder) base64() {
	/*
		IyBCeXRlQ2FzdGVyCgpTd2lzcy1rbmlmZSBmaWxlI...
	*/
	encoded := base64.StdEncoding.EncodeToString(e.input)
	e.output = []byte(encoded)
}

func (e *encoder) ipv4() {
	/*
		IPv4-fuscation:
		"252.72.131.228\0" + "252.72.131.228\0" + "252.72.131.228\0"
	*/
	NUM_OF_BYTES := 4

	var result []byte
	rest := len(e.input) % NUM_OF_BYTES

	for i := 0; i < len(e.input)-rest; i += NUM_OF_BYTES {
		ip := net.IP(e.input[i : i+NUM_OF_BYTES]).String()
		bytes := append([]byte(ip), 0x00)
		e.output = append(e.output, bytes...)
	}

	if rest == 0 {
		e.output = result
		return
	}

	length := len(e.input)
	lastBytes := e.input[length-(NUM_OF_BYTES-rest):]

	for range rest {
		lastBytes = append(lastBytes, 255)
	}

	lastIp := net.IP(lastBytes).String()
	bytes := append([]byte(lastIp), 0x00)

	e.output = append(e.output, bytes...)
}

func (e *encoder) mac() {
	/*
		MAC-fuscation:
		"fc-48-83-e4-f0-e8\0" + "fc-48-83-e4-f0-e8\0" + "fc-48-83-e4-f0-e8\0"
	*/
	NUM_OF_BYTES := 6

	var result []byte
	rest := len(e.input) % NUM_OF_BYTES

	for i := 0; i < len(e.input)-rest; i += NUM_OF_BYTES {
		ip := net.HardwareAddr(e.input[i : i+NUM_OF_BYTES]).String()
		bytes := append([]byte(ip), 0x00)
		e.output = append(e.output, bytes...)
	}

	if rest == 0 {
		e.output = result
		return
	}

	length := len(e.input)
	lastBytes := e.input[length-(NUM_OF_BYTES-rest):]

	for range rest {
		lastBytes = append(lastBytes, 255)
	}

	lastIp := net.HardwareAddr(lastBytes).String()
	bytes := append([]byte(lastIp), 0x00)

	e.output = append(e.output, bytes...)
}
