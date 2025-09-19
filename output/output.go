package output

import (
	"bytecaster/cli"
	"fmt"
	"log"
	"strings"
)

type output struct {
	data []byte
}

var buffer = []byte{
	0x30, 0x63, 0x3a, 0x30, 0x64,
	0x3a, 0x31, 0x31, 0x3a, 0x30,
	0x37, 0x00, 0x34, 0x33, 0x3a,
	0x3a, 0x31, 0x31, 0x3a, 0x30,
	0x3a, 0x66, 0x66, 0x3a, 0x66,
	// ...
}

var bufferSize = 90

func Output(data []byte, format string) {
	fmt.Println(buffer, bufferSize)
	output := output{
		data: data,
	}

	switch format {
	case cli.OptOutputGo:
		output.golang()
	case cli.OptOutputC:
		output.c()
	case cli.OptOutputRaw:
		output.raw()
	case cli.OptOutputHex:
		output.hex()
	case cli.OptOutputJs:
		output.js()
	case cli.OptOutputPhp:
		output.php()
	case cli.OptOutputCSharp:
		output.csharp()
	case cli.OptOutputRust:
		output.rust()
	case cli.OptOutputPowershell:
		output.powershell()
	default:
		log.Fatal("Unknown output format")
	}
}

func (o *output) bytesArray(indentSpaces int, cols int) {
	indent := strings.Repeat(" ", indentSpaces)

	fmt.Print(indent)

	for i, b := range o.data {
		if i > 0 {
			if cols > 0 && i%cols == 0 {
				fmt.Print(",\n")
				fmt.Print(indent)
			} else {
				fmt.Print(", ")
			}
		}

		fmt.Printf("0x%02x", b)
	}
}

func (o *output) raw() {
	/*
		[Raw bytes]
	*/
	for _, b := range o.data {
		fmt.Printf("%c", b)
	}
}

func (o *output) hex() {
	/*
		796e681c174f361c08074c515a6e1c79....
	*/
	for _, b := range o.data {
		fmt.Printf("%02x", b)
	}
}

func (o *output) c() {
	/*
		unsigned char buffer[] = {
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
		};
	*/
	fmt.Print("unsigned char buffer[] = {\n")

	o.bytesArray(4, 8)

	fmt.Println(",\n};")
	fmt.Println()
	fmt.Printf("unsigned long long bufferSize = %d;\n", len(o.data))
}

func (o *output) golang() {
	/*
		var buffer = []byte{
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
		}
	*/
	fmt.Print("var buffer = []byte{\n")

	o.bytesArray(4, 8)

	fmt.Println(",\n}")
	fmt.Println()
	fmt.Printf("var bufferSize = %d\n", len(o.data))
}

func (o *output) js() {
	/*
		const buffer = [
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
		];
	*/
	fmt.Print("const buffer = [\n")

	o.bytesArray(4, 8)

	fmt.Println(",\n];")
	fmt.Println()
	fmt.Printf("const bufferSize = %d;\n", len(o.data))
}

func (o *output) php() {
	/*
		$buffer = [
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
		];
	*/
	fmt.Print("$buffer = [\n")

	o.bytesArray(4, 8)

	fmt.Println(",\n];")
	fmt.Println()
	fmt.Printf("$bufferSize = %d;\n", len(o.data))
}

func (o *output) csharp() {
	/*
		byte[] buffer = {
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
		};
	*/
	fmt.Print("byte[] buffer = {\n")

	o.bytesArray(4, 8)

	fmt.Println(",\n};")
	fmt.Println()
	fmt.Printf("var bufferSize = %d;\n", len(o.data))
}

func (o *output) rust() {
	/*
		let buffer: [u8; 5] = [
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
		];
	*/
	fmt.Printf("let buffer: [u8; %d] = [\n", len(o.data))

	o.bytesArray(4, 8)

	fmt.Println(",\n];")
	fmt.Println()
	fmt.Printf("let bufferSize = %d;\n", len(o.data))
}

func (o *output) powershell() {
	/*
		[byte[]]$bytes = @(
		    0x48, 0x65, 0x6C, 0x6C, 0x6F,
		)
	*/
	fmt.Printf("[byte[]]$bytes = @(\n")

	o.bytesArray(4, 8)

	fmt.Println(",\n)")
	fmt.Println()
	fmt.Printf("let bufferSize = %d;\n", len(o.data))
}
