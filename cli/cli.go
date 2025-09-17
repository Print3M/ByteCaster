package cli

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

const VERSION = "0.0.9"

var (
	OptEncryptionXor        = "xor"
	SupportedEncryptionAlgs = []string{
		OptEncryptionXor,
	}
	// TODO: vigenere
)

var (
	OptEncodingBase64  = "base64"
	SupportedEncodings = []string{
		OptEncodingBase64,
	}
)

var (
	OptOutputC             = "c"
	OptOutputGo            = "go"
	OptOutputPowershell    = "powershell"
	OptOutputCSharp        = "csharp"
	OptOutputPhp           = "php"
	OptOutputJs            = "js"
	OptOutputRust          = "rust"
	OptOutputHex           = "hex"
	OptOutputRaw           = "raw"
	SupportedOutputFormats = []string{
		OptOutputC, OptOutputGo, OptOutputPowershell,
		OptOutputPhp, OptOutputJs, OptOutputRust,
		OptOutputHex, OptOutputRaw, OptOutputCSharp,
	}
)

type CliFlags struct {
	Input             string
	OutputFormat      string
	EncryptionKey     string
	EncryptionAlg     string
	Encoding          string
	ShowVersion       bool
	EncryptionEnabled bool
	EncodingEnabled   bool
}

func ParseCli() *CliFlags {
	var flags CliFlags

	flag.StringVar(&flags.Input, "i", "", "")
	flag.StringVar(&flags.Input, "input", "", "")

	flag.StringVar(&flags.OutputFormat, "f", "", "")
	flag.StringVar(&flags.OutputFormat, "format", "", "")

	flag.StringVar(&flags.Encoding, "e", "", "")
	flag.StringVar(&flags.Encoding, "encoding", "", "")

	flag.StringVar(&flags.EncryptionAlg, "x", "", "")
	flag.StringVar(&flags.EncryptionAlg, "encryption-alg", "", "")

	flag.StringVar(&flags.EncryptionKey, "k", "", "")
	flag.StringVar(&flags.EncryptionKey, "encryption-key", "", "")

	flag.BoolVar(&flags.ShowVersion, "v", false, "")
	flag.BoolVar(&flags.ShowVersion, "version", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: Bin2Code -i <path> -f <value> \n")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Printf("  %-25s %s\n", "-i, --input <path>", "Binary input file (required)")
		fmt.Printf("  %-25s %s\n", "-f, --format <value>", "Output format (required):")
		fmt.Printf("  %-25s > %v\n", "", strings.Join(SupportedOutputFormats, ", "))
		fmt.Printf("  %-25s %s\n", "-e, --encoding <value>", "Output encoding:")
		fmt.Printf("  %-25s > %v\n", "", strings.Join(SupportedEncodings, ", "))
		fmt.Printf("  %-25s %s\n", "-x, --enc-alg <value>", "Encryption algorithm:")
		fmt.Printf("  %-25s > %v\n", "", strings.Join(SupportedEncryptionAlgs, ", "))
		fmt.Printf("  %-25s %s\n", "-k, --enc-key <string>", "Encryption key")
		fmt.Printf("  %-25s %s\n", "-v, --version", "Show version")
		fmt.Printf("  %-25s %s\n", "-h, --help", "Show this help")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println()
		fmt.Println("  Bin2Code -i shellcode.bin -f go -x xor -k StrongKey123 -e base64")
		fmt.Println()
		fmt.Println("Created by Print3M (print3m.github.io)")
		fmt.Println()
	}

	flag.Parse()

	if flags.ShowVersion {
		fmt.Printf("Bin2Code %s\n", VERSION)
		os.Exit(0)
	}

	// Required flags
	if flags.Input == "" || flags.OutputFormat == "" {
		flag.Usage()
		os.Exit(1)
	}

	flags.OutputFormat = strings.ToLower(flags.OutputFormat)

	if len(flags.OutputFormat) > 0 && !slices.Contains(SupportedOutputFormats, flags.OutputFormat) {
		fmt.Fprintf(os.Stderr, "Output format '%s' not supported!\n", flags.OutputFormat)
		os.Exit(1)
	}

	flags.Encoding = strings.ToLower(flags.Encoding)
	flags.EncodingEnabled = len(flags.Encoding) > 0

	if flags.EncodingEnabled && !slices.Contains(SupportedEncodings, flags.Encoding) {
		fmt.Fprintf(os.Stderr, "Encoding '%s' not supported!\n", flags.Encoding)
		os.Exit(1)
	}

	flags.EncryptionAlg = strings.ToLower(flags.EncryptionAlg)
	flags.EncryptionEnabled = len(flags.EncryptionAlg) > 0

	if flags.EncryptionEnabled {
		if !slices.Contains(SupportedEncryptionAlgs, flags.EncryptionAlg) {
			fmt.Fprintf(os.Stderr, "Encryption algorithm '%s' not supported!\n", flags.EncryptionAlg)
			os.Exit(1)
		}

		if len(flags.EncryptionKey) == 0 {
			fmt.Fprintf(os.Stderr, "Encryption key not set!\n")
			os.Exit(1)
		}
	}

	return &flags
}
