package main
/*
 * >   increment the data pointer (to point to the next cell to the right).
 * <   decrement the data pointer (to point to the next cell to the left).
 * +   increment (increase by one) the byte at the data pointer.
 * -   decrement (decrease by one) the byte at the data pointer.
 * .   output a character, the ASCII value of which being the byte at the 
 *     data pointer.
 * ,   accept one byte of input, storing its value in the byte at the data 
 *     pointer.
 * [   if the byte at the data pointer is zero, then instead of moving the 
 *     instruction pointer forward to the next command, jump it forward to 
 *     the command after the matching ] command*.
 * ]   if the byte at the data pointer is nonzero, then instead of moving 
 *     the instruction pointer forward to the next command, jump it back 
 *     to the command after the matching [ command*.
 */

import (
	"fmt"
	"flag"
	"os"
	"io/ioutil"
)

func interpreter(code []byte) {
	tape := make([]byte, 10000)
	pointer := 5000
	codeptr := 0
	one_byte := make([]byte, 1, 1)
	for code[codeptr] != 0 {
		switch code[codeptr] {
		case '>':
			pointer++
			if pointer >= len(tape) {
				pointer = 0
			}
		case '<':
			pointer--
			if pointer < 0 {
				pointer = len(tape) - 1
			}
		case '+':
			tape[pointer]++
		case '-':
			tape[pointer]--
		case '.':
			fmt.Printf("%c", tape[pointer])
		case ',':
			os.Stdin.Read(one_byte)
			tape[pointer] = one_byte[0]
		case '[':
			if tape[pointer] == 0 {
				counter := 1
				for !(code[codeptr] == ']' && counter == 0) {
					codeptr++
					if code[codeptr] == '[' {
						counter++
					} else if code[codeptr] == ']' {
						counter--
					}
				}
			}
		case ']':
			if tape[pointer] != 0 {
				counter := 1
				for !(code[codeptr] == '[' && counter == 0) {
					codeptr--
					if code[codeptr] == ']' {
						counter++
					} else if code[codeptr] == '[' {
						counter--
					}
				}
			}
		}
		codeptr++
	}
}

func is_bf_char(item byte) bool {
	chars := [8]byte{'>', '<', '+', '-', '.', ',', '[', ']'}
	var contained bool
	for _, v := range chars {
		if item == v {
			contained = true
			break
		}
	}
	return contained
}

func main() {
	filename := flag.String("src", "", "Specify the codefile to use.")
	flag.Parse()
	source, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Print("File could not be read.")

	}
	code := make([]byte, 0)
	for _, v := range source {
		if is_bf_char(v) {
			code = append(code, v)
		}
	}
	code = append(code, 0)
	interpreter(code)
}
