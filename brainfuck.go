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
    )

func interpreter() {
    tape := make([]byte, 30000)
    pointer := 15000
    codeptr := 0
    code := make([]byte, 0, 10000)
    one_byte := make([]byte, 1, 1)
    for code[codeptr] != 0 {
        switch code[codeptr] {
            case '>':
                pointer++
            case '<':
                pointer--
            case '+':
                tape[pointer]++
            case '-':
                tape[pointer]--
            case '.':
                fmt.Printf("%q", tape[pointer])
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
                    if !(code[codeptr] == '[' && counter == 0) {
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

func load_file(filename string) {

}

func main() {
    filename := flag.String("file", "", "Specify the codefile to use.")
    file, err := os.Open(*filename)
    if err != nil {
        return
    }

    n, err := file.Read(code)
    fmt.Printf("Read %d bytes.", n)
    if err != nil {
        return
    }

}
