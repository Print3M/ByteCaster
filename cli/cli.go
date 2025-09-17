package cli

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

const VERSION = "0.0.9"

type CliFlags struct {
	Input         string
	OutputFormat  string
	EncryptionKey string
	EncryptionAlg string
	Encoding      string
	ShowVersion   bool
}

var SupportedEncryptionAlgs = []string{"xor"}
var SupportedEncodings = []string{"base64"} // TODO: add more
var SupportedOutputFormats = []string{"c", "go", "bash", "php", "js", "rust", "hex", "raw"}

func ParseCli() *CliFlags {
	var flags CliFlags

	flag.StringVar(&flags.Input, "i", "", "")
	flag.StringVar(&flags.Input, "input", "", "")

	flag.StringVar(&flags.Encoding, "e", "", "")
	flag.StringVar(&flags.Encoding, "encoding", "", "")

	flag.StringVar(&flags.EncryptionAlg, "x", "", "")
	flag.StringVar(&flags.EncryptionAlg, "encryption-alg", "", "")

	flag.StringVar(&flags.EncryptionKey, "k", "", "")
	flag.StringVar(&flags.EncryptionKey, "encryption-key", "", "")

	flag.BoolVar(&flags.ShowVersion, "v", false, "")
	flag.BoolVar(&flags.ShowVersion, "version", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: Bin2Code -i <path> -f <string> -e <algorithm> -k <string> \n")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Printf("  %-26s %s\n", "-i, --input <path>", "Binary input file (required)")
		fmt.Printf("  %-26s %s\n", "-f, --format <string>", "Output format (required)")
		fmt.Printf("  %-26s %s\n", "-e, --encoding <string>", "Output encoding")
		fmt.Printf("  %-26s %s\n", "-e, --encryption-alg <algorithm>", "Encryption algorithm")
		fmt.Printf("  %-26s %s\n", "-k, --encryption-key <string>", "Encryption key")
		fmt.Printf("  %-26s %s\n", "-v, --version", "Show version of Bin2Code")
		fmt.Printf("  %-26s %s\n", "-h, --help", "Show this help")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println()
		fmt.Println("  Bin2Code -i shellcode.bin -f c -e xor -e StrongKey123")
		fmt.Println()
		fmt.Println("Created by Print3M (print3m.github.io)")
		fmt.Println()
	}

	flag.Parse()

	if flags.ShowVersion {
		fmt.Printf("Bin2Code %s\n", VERSION)
		os.Exit(0)
	}

	if flags.Input == "" || flags.OutputFormat == "" {
		flag.Usage()
		os.Exit(1)
	}

	flags.OutputFormat = strings.ToLower(flags.OutputFormat)
	if len(flags.OutputFormat) > 0 && slices.Contains(SupportedOutputFormats, flags.OutputFormat) {
		fmt.Fprintf(os.Stderr, "Output format '%s' not supported.", flags.OutputFormat)
		os.Exit(1)
	}

	flags.Encoding = strings.ToLower(flags.Encoding)
	if len(flags.Encoding) > 0 && slices.Contains(SupportedEncodings, flags.Encoding) {
		fmt.Fprintf(os.Stderr, "Encoding '%s' not supported.", flags.Encoding)
		os.Exit(1)
	}

	flags.EncryptionAlg = strings.ToLower(flags.EncryptionAlg)
	if len(flags.EncryptionAlg) > 0 {
		if slices.Contains(SupportedEncryptionAlgs, flags.EncryptionAlg) {
			fmt.Fprintf(os.Stderr, "Encryption algorithm '%s' not supported.", flags.EncryptionAlg)
			os.Exit(1)
		}

		if len(flags.EncryptionKey) == 0 {
			fmt.Fprintf(os.Stderr, "Encryption key not set!")
			os.Exit(1)
		}
	}

	return &flags
}
