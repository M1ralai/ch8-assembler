package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func write(instr string) {
	fmt.Println(instr)
	value, err := strconv.ParseUint(instr, 16, 16)
	fmt.Println(value)
	if err != nil {
		log.Fatal(err)
	}
	high := byte((value & 0xFF00) >> 8)
	low := byte(value & 0x00FF)
	fmt.Println(high, low)
	file, err := os.OpenFile("out/output.ch8", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err := file.Write([]byte{high, low}); err != nil {
		log.Fatal(err)
	}
}

func JP(line string) {
	s := strings.Split(line, " ")
	var search string
	for i := range s[1] {
		if s[1][i] != 0 {
			search += string(s[1][i])
		}
	}
	write("1" + search)
}

func CALL(line string) {
	s := strings.Split(line, " ")
	var search string
	for i := range s[1] {
		if s[1][i] != 0 {
			search += string(s[1][i])
		}
	}
	write("2" + search)
}

func SE(line string) {
	s := strings.Split(line, " ")
	var r1, r2 string
	for i := range s {
		fmt.Println(s[i])
	}
	for i := range s[1] {
		if s[1][i] != 0 {
			r1 += string(s[1][i])
		}
	}
	for i := range s[2] {
		if s[2][i] != 0 {
			r2 += string(s[2][i])
		}
	}
	if r2[0] != 'V' {
		write("3" + string(r1[1]) + r2)
	} else {
		write("5" + string(r1[1]) + string(r2[1]) + "0")
	}
}

func SNE(line string) {
	s := strings.Split(line, " ")
	var r1, r2 string
	for i := range s {
		fmt.Println(s[i])
	}
	for i := range s[1] {
		if s[1][i] != 0 {
			r1 += string(s[1][i])
		}
	}
	for i := range s[2] {
		if s[2][i] != 0 {
			r2 += string(s[2][i])
		}
	}
	if r2[0] != 'V' {
		write("4" + string(r1[1]) + r2)
	} else {
		write("9" + string(r1[1]) + string(r2[1]) + "0")
	}
}
func ALU(line string, op string) {
	s := strings.Split(line, " ")
	var n string
	switch op {
	case "OR":
		n = "1"
	case "AND":
		n = "2"
	case "XOR":
		n = "3"
	case "SUB":
		n = "5"
	case "SHR":
		n = "6"
	case "SUBN":
		n = "7"
	case "SHL":
		n = "E"
	}
	write("8" + string(s[1][1]) + string(s[2][1]) + n)
}

func ADD(line string) {
	s := strings.Split(line, " ")
	if string(s[1][0]) == "V" && string(s[2][0]) == "V" {
		write("8" + string(s[1][1]) + string(s[2][1]) + "4")
	} else if string(s[1][0]) == "I" && string(s[2][0]) == "V" {
		write("F" + string(s[2][1]) + "1E")
	} else {
		var search string
		for i := range s[2] {
			if s[2][i] != 0 {
				search += string(s[2][i])
			}
		}
		write("7" + string(s[1][1]) + search)
	}
}

func DRW(line string) {
	s := strings.Split(line, " ")
	var search string
	for i := range s[3] {
		if s[3][i] != 0 {
			search += string(s[3][i])
		}
	}
	write("D" + string(s[1][1]) + string(s[2][1]) + search)
}

func RND(line string) {
	s := strings.Split(line, " ")
	var search string
	for i := range s[2] {
		if s[2][i] != 0 {
			search += string(s[2][i])
		}
	}
	write("C" + string(s[1][1]) + search)
}

func SKP(line string) {
	s := strings.Split(line, " ")
	var search string
	for i := range s[1] {
		if s[1][i] != 0 {
			search += string(s[1][i])
		}
	}
	write("E" + string(search[1]) + "9E")
}

func SKNP(line string) {
	s := strings.Split(line, " ")
	var search string
	for i := range s[1] {
		if s[1][i] != 0 {
			search += string(s[1][i])
		}
	}
	write("E" + string(search[1]) + "A1")
}

func instructor(line string) {
	s := strings.Split(line, " ")
	var search string
	fmt.Println(s)
	for i := range s[0] {
		if s[0][i] != 0 {
			search += string(s[0][i])
		}
	}
	switch search {
	case "LD":

	case "CLS":
		write("00E0")
	case "RET":
		write("00EE")
	case "JP":
		JP(line)
	case "CALL":
		CALL(line)
	case "SE":
		SE(line)
	case "SNE":
		SNE(line)
	case "OR":
		ALU(line, "OR")
	case "AND":
		ALU(line, "AND")
	case "XOR":
		ALU(line, "XOR")
	case "ADD":
		ADD(line)
	case "SUB":
		ALU(line, "SUB")
	case "SHR":
		ALU(line, "SHR")
	case "SUBN":
		ALU(line, "SUBN")
	case "SHL":
		ALU(line, "SHL")
	case "RND":
		RND(line)
	case "DRW":
		DRW(line)
	case "SKP":
		SKP(line)
	case "SKNP":
		SKNP(line)
	default:
		fmt.Println("i didnt seen this boy in my entire life")
	}
}

func fetch(asm string) {
	s := strings.Split(asm, "\n")
	for i := range s {
		if s[i] != "" {
			instructor(s[i])
		}
	}
}

func read(p string) string {
	file, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := make([]byte, 4086)
	if _, err := file.Read(data); err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func reset() {
	os.Remove("out/output.ch8")
}

func main() {
	os.Chdir("cmd")
	reset()
	os.Create("out/output.ch8")
	asm := read("asm/first.asm")
	fetch(asm)
}
