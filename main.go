package main

import (
	"bytecaster/cli"
	"bytecaster/encoding"
	"bytecaster/encryption"
	"bytecaster/output"
	"fmt"
	"log"
	"os"
)

func main() {
	flags := cli.ParseCli()

	fileInfo, err := os.Stat(flags.Input)
	if err != nil {
		log.Fatal("Error stating file:", err)
	}

	fileSize := fileInfo.Size()
	data := make([]byte, fileSize)

	file, err := os.Open(flags.Input)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	_, err = file.Read(data)
	if err != nil && err.Error() != "EOF" {
		log.Fatal("Error reading file:", err)
	}

	// Encryption
	if flags.EncryptionEnabled {
		data = encryption.EncryptData(data, flags.EncryptionAlg, flags.EncryptionKey)
	}

	// Encoding
	if flags.EncodingEnabled {
		data = encoding.EncodeData(data, flags.Encoding)
	}

	// Output format
	output.Output(data, flags.OutputFormat)

	fmt.Println()

	os.Exit(0)
}
