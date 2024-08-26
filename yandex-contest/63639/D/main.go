package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscan(reader, &N)

	var commands string
	fmt.Fscan(reader, &commands)

	uniquePositions := make(map[int]bool)
	//var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			for _, newCommand := range "FRL" {
				if commands[i] == byte(newCommand) {
					continue
				}
				newCommands := commands[:i] + string(newCommand) + commands[i+1:]
				newPos := 0
				direction := 1 // 1 = right, 2 = left

				for _, command := range newCommands {
					switch command {
					case 'F':
						switch direction {
						case 1:
							newPos++
						case 2:
							newPos--
						}
					case 'R':
						direction = 1
					case 'L':
						direction = 2
					}
				}

				//mu.Lock()
				uniquePositions[newPos] = true
				//mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(len(uniquePositions))
}

//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//)
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	var N int
//	fmt.Fscan(reader, &N)
//
//	var commands string
//	fmt.Fscan(reader, &commands)
//
//	uniquePositions := make(map[int]bool)
//
//	for i := 0; i < N; i++ {
//		for _, newCommand := range "FRL" {
//			if commands[i] == byte(newCommand) {
//				continue
//			}
//			newCommands := commands[:i] + string(newCommand) + commands[i+1:]
//			newPos := simulate(newCommands)
//			uniquePositions[newPos] = true
//		}
//	}
//	fmt.Println(len(uniquePositions))
//}
//
//func simulate(commands string) int {
//	x := 0
//	direction := 1 // 1 = right, 2 = left
//
//	for _, command := range commands {
//		switch command {
//		case 'F':
//			switch direction {
//			case 1:
//				x++
//			case 2:
//				x--
//			}
//		case 'R':
//			direction = 1
//		case 'L':
//			direction = 2
//		}
//	}
//
//	return x
//}
