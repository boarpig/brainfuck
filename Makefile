brainfuck: brainfuck.go
	6g brainfuck.go
	6l -o brainfuck brainfuck.6

run:
	./brainfuck -src hello.bf

clean:
	rm brainfuck.6
	rm brainfuck
