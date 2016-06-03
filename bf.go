package main

import (
	"fmt"
)

func brainf(bfs string) {
	var pc []int
	memory := make([]byte, 1024)
	ptr := 0
	i := 0
	for i < len(bfs) {
		switch bfs[i] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			fmt.Printf("%c", memory[ptr])
		case ',':
			fmt.Scanf("%c", &memory[ptr])
		case '[':
			pc = append(pc, i)
		case ']':
			if memory[ptr] == 0 {
				pc = pc[:len(pc)-1]
			} else {
				i = pc[len(pc)-1]
			}
		default:
			break
		}
		i++
	}
}

func main() {
	helloWorld := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
	brainf(helloWorld)
}
