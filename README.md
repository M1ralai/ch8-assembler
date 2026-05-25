# ch8-assembler

i wrote this while learning Go.

which is kind of funny, because at the same time i was also trying to write a CHIP-8 emulator in Rust.

so apparently my brilliant learning path was:

> learn Rust by writing a virtual machine  
> learn Go by writing an assembler for that virtual machine

normal people probably build a todo app.

i tried to make fake machine code.

---

## what this is

this is a tiny experimental assembler for CHIP-8 assembly.

it reads simple assembly-like instructions and writes the corresponding CHIP-8 binary output into a `.ch8` file.

the idea was simple:

```text
JP 200
ADD V1 05
DRW V0 V1 5
```

turns into raw bytes that a CHIP-8 emulator can load and execute.

it is not a serious assembler.
it is not complete.
it is not clean.
it is not production-ready.

but it was one of those projects where i started understanding that source code is just a nicer lie we tell ourselves before everything becomes bytes.

---

## why i made this

while writing the CHIP-8 emulator, i kept looking at opcodes like:

```text
0x600A
0x7101
0xA2F0
```

and at some point i wanted to go in the other direction.

not:

```text
bytes -> instruction
```

but:

```text
instruction -> bytes
```

that is basically what an assembler does.

so this became the other half of the emulator project.

the Rust project was me asking:

> how does a machine execute instructions?

this Go project was me asking:

> how do instructions become machine-readable bytes in the first place?

and honestly, that pair taught me more than a lot of “beginner projects” ever could.

---

## what it does

the assembler reads an assembly file, parses each line, maps the instruction to a CHIP-8 opcode, and writes the result as binary output.

roughly:

```text
read assembly file
→ split into lines
→ parse instruction
→ generate opcode
→ split opcode into high byte / low byte
→ write bytes into output.ch8
```

the core idea is this:

```go
high := byte((value & 0xFF00) >> 8)
low := byte(value & 0x00FF)
file.Write([]byte{high, low})
```

that part still feels cool.

because it is the exact moment where a human-readable instruction becomes something a virtual CPU can actually fetch from memory.

---

## supported-ish instructions

some CHIP-8 instructions were implemented or partially implemented:

- `CLS`
- `RET`
- `JP`
- `CALL`
- `SE`
- `SNE`
- `OR`
- `AND`
- `XOR`
- `ADD`
- `SUB`
- `SHR`
- `SUBN`
- `SHL`
- `RND`
- `DRW`
- `SKP`
- `SKNP`

some parts are unfinished, especially `LD`.

some parsing is extremely fragile.

some code is very clearly written by someone who was learning the language and the concept at the same time.

which is exactly what happened.

---

## running

```bash
go run ./cmd
```

the program expects an input assembly file around:

```text
cmd/asm/first.asm
```

and writes the output binary to:

```text
cmd/out/output.ch8
```

---

## project status

abandoned / deprecated / emotionally preserved.

i am not maintaining this as a real assembler.

this repository is more like a small fossil from the period where i was trying to understand:

- Go syntax
- file IO
- binary output
- instruction encoding
- assemblers
- emulators
- CHIP-8
- why computers are just cursed byte machines

there are better ways to write this.

there are cleaner ways to parse assembly.

there are proper assembler architectures with tokenizers, parsers, labels, symbol tables, error reporting, and multiple passes.

this is not that.

this is just me poking the machine until it answered.

---

## the funny part

this project exists because i was writing a CHIP-8 emulator in Rust and then decided:

> wait, i also need my own assembler, right?

no, i absolutely did not need my own assembler.

but writing it made the emulator feel more real.

suddenly the loop was complete:

```text
assembly code
→ assembler
→ .ch8 binary
→ emulator memory
→ fetch/decode/execute
```

that is a beautiful little pipeline.

even if the code is ugly.

---

## future

maybe one day i will rewrite this properly with:

- real lexer
- real parser
- label support
- better error messages
- instruction validation
- proper CLI arguments
- tests
- multiple output formats

or maybe i will leave it exactly like this.

a weird little Go project sitting next to a weird little Rust emulator.

two unfinished projects accidentally teaching me how software becomes bytes, and how bytes become behavior.
