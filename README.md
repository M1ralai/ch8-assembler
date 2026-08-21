# CHIP-8 Assembler Experiment

A small CHIP-8 assembler experiment written in Go. The project was built to practice instruction parsing, bit-level opcode encoding, and binary output rather than to provide a complete toolchain.

## What it demonstrates

- Splitting assembly source into instructions and operands
- Dispatching mnemonics to encoding functions
- Building 16-bit CHIP-8 opcodes from register and immediate fields
- Writing each opcode as high and low bytes to a `.ch8` file

## Parsing and encoding

`cmd/main.go` reads the input file line by line, selects an encoder from the first token, and constructs the four-hex-digit opcode for that instruction. The encoded value is parsed as a 16-bit integer and emitted in big-endian byte order.

The current dispatcher includes `CLS`, `RET`, `JP`, `CALL`, `SE`, `SNE`, `OR`, `AND`, `XOR`, `ADD`, `SUB`, `SHR`, `SUBN`, `SHL`, `RND`, `DRW`, `SKP`, and `SKNP`. `LD` is recognized but not implemented.

This repository contains the assembler only. The separate [chip-8 repository](https://github.com/M1ralai/chip-8) contains an independent Rust emulator experiment.

## Build and run

Requirements: Go and Make.

```bash
make build
make run
```

The current entry point uses fixed project-relative paths:

- input: `cmd/asm/first.asm`
- output: `cmd/out/output.ch8`

It can also be run directly:

```bash
go run ./cmd
```

## Limitations

- The assembler has no lexer, labels, symbols, macros, or command-line file arguments.
- Operand validation and error reporting are minimal.
- The instruction set is incomplete, including the unimplemented `LD` family.
- Parsing assumes a narrow whitespace and operand format.
- There is no automated conformance test suite.
