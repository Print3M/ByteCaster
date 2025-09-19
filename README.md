# ByteCaster

Swiss-knife file to array of bytes converter in multiple languages with built-in encryption & encoding.

- No external dependencies
- Standalone binary

// TODO: Graphic with examples of each stage

```text
  Input
    |
Encryption + Key
    |
    |
 Encoding
    |
    |
  Output
Formatting

```

## Installation

[Download the compiled binary](TODO) or compile Go source code.

## Usage

Example:

```bash
# Backdoor version.dll (proxy to absolute path)
./ByteCaster -i data.bin
```

**`-i / --input <path>`** [required]

Binary input file.

**`-f / --format <value>`** [required]

Output format of the processed data. This generates the final data as an array of bytes in the selected programming language. Output is always sent to STDOUT.

To avoid applying any formatting output, use the `raw` value.

Available values: `raw`, `hex`, `c`, `go`, `powershell`, `php`, `js`, `rust`, `csharp`

// TODO: Nim, Zig

**`-x / --enc-alg <value>` + `-k / --enc-key <string>`**

Data encryption. Both parameters, the encryption algorithm and the key (string), must be provided.

Availabe values: `xor`

**`-e / --encoding <value>`**

Data encoding. Often used as obfuscation to confuse analysis or changes in the entropy level of data.

Available values: `base64`, `ipv4`, `mac`

## Processing order

1. Encryption
2. Encoding
3. Output format

## Encryption algorithms

**`xor`** [0% overhead]

Typical simple XOR encryption (`a^b`). Each byte is XORed with the byte from the key.

## Encoding algorithms

**`base64`** [33%-37% overhead]

Standard Base64 encoding. We are using [the standard Go library functions here](https://pkg.go.dev/encoding/base64).

**`ipv4`**

This is known as the _IPv4Fuscation_ technique. Each output byte is converted to one octet in the IPv4 address as a decimal number.

Example data:

```text
{ 0xe9, 0x36, 0x17, 0xbb, 0xbd, 0x7f, 0x22, 0x10 }
```

The output (array of bytes) looks exactly like this in memory:

`233.54.23.187\0189.127.34.16\0` ...

> NOTE:
>
> - Each IP address ends with a null byte!
> - If the number of bytes is not divisible by 4, the missing bytes added to the last IP address are 255.

**`mac`**

This is known as the _MACFuscation_ technique. Each output byte is converted to one octet in the MAC address as a hexadecimal number (lowercase).

Example data:

```text
{ 0xe9, 0x36, 0x17, 0xbb, 0xbd, 0x7f, 0x22, 0x10, 0x84, 0xA7, 0x6f, 0xcc }
```

The output (array of bytes) looks exactly like this in memory:

`e9:36:17:bb:bd:7f\022:10:84:a7:6f:cc\0`

> NOTE:
>
> - Each MAC address ends with a null byte!
> - Hexadecimal numbers are lowercase.
> - If the number of bytes is not divisible by 6, the missing bytes added to the last MAC address are 255 (`ff`).

## Credits

- [HellShell](https://github.com/NUL0x4C/HellShell) - inspired me to implement _IPv4Fuscation_ and _MACFuscation_.
