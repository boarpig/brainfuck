brainfuck:
	6g brainfuck.go
	6l -o brainfuck brainfuck.6

run:
	./brainfuck code.txt

clean:
	rm brainfuck.6
	rm brainfuck
