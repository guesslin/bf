package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	srcFile := flag.String("file", "", "bf source file")
	stdin := flag.Bool("stdin", false, "Get bf code from stdin")
	flag.Parse()
	if *stdin {
		var bfs string
		fmt.Scanln(&bfs)
		brainf(bfs)
	} else {
		bfs, err := ioutil.ReadFile(*srcFile)
		if err != nil {
			fmt.Printf("BF File Read Error %v\n", err)
			return
		}
		brainf(string(bfs))
	}
}
