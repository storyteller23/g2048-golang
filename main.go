package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

var score = 0

func main() {
	rand.Seed(time.Now().UnixNano())
	area := createArea()
	for {
		getTwoOfFour(area)
		displayArea(area)
		for {
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			switch string(char) {
			case "w":
				toUp(area)
			case "s":
				toDown(area)
			case "a":
				toLeft(area)
			case "d":
				toRight(area)
			case "q":
				return
			default:
				continue
			}
			break
		}
	}
}

func createArea() [][]int {
	res := make([][]int, 4)
	for i := 0; i < len(res); i++ {
		res[i] = append(res[i], make([]int, 4)...)
	}
	return res
}

func displayArea(area [][]int) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("o--------o--------o--------o--------o")
	for _, row := range area {
		fmt.Println("|        |        |        |        |")
		fmt.Print("|")
		for _, n := range row {
			if n == 0 {
				fmt.Print("        |")
			} else {
				fmt.Printf("%5d   |", n)
			}
		}
		fmt.Println()
		fmt.Println("|        |        |        |        |")
		fmt.Println("o--------o--------o--------o--------o")
	}
	fmt.Printf("\t\tSCORE: %d\n", score)
	fmt.Println("'w' - UP")
	fmt.Println("'s' - DOWN")
	fmt.Println("'a' - LEFT")
	fmt.Println("'d' - RIGHT")
	fmt.Println("'q' - EXIT")
}

func zerosCoord(area [][]int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area); j++ {
			if area[i][j] == 0 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func getTwoOfFour(area [][]int) {
	zeros := zerosCoord(area)
	if len(zeros) == 0 {
		return
	}
	choice := rand.Intn(len(zeros))
	if rand.Intn(2) == 0 {
		area[zeros[choice][0]][zeros[choice][1]] = 2
	} else {
		area[zeros[choice][0]][zeros[choice][1]] = 4
	}
}

func toUp(area [][]int) {
	for k := 0; k < len(area); k++ {
		for i := 0; i < len(area); i++ {
			for j := 1; j < len(area); j++ {
				if area[j-1][i] == 0 && area[j][i] != 0 {
					area[j-1][i], area[j][i] = area[j][i], area[j-1][i]
				}
			}
		}
	}
	for i := 0; i < len(area); i++ {
		for j := 1; j < len(area); j++ {
			if area[j-1][i] == area[j][i] {
				area[j-1][i] = area[j-1][i] + area[j][i]
				score += area[j-1][i]
				area[j][i] = 0
			}
		}
	}
	for k := 0; k < len(area); k++ {
		for i := 0; i < len(area); i++ {
			for j := 1; j < len(area); j++ {
				if area[j-1][i] == 0 && area[j][i] != 0 {
					area[j-1][i], area[j][i] = area[j][i], area[j-1][i]
				}
			}
		}
	}
}

func toDown(area [][]int) {
	for k := 0; k < len(area); k++ {
		for i := len(area) - 1; i >= 0; i-- {
			for j := len(area) - 2; j >= 0; j-- {
				if area[j+1][i] == 0 && area[j][i] != 0 {
					area[j+1][i], area[j][i] = area[j][i], area[j+1][i]
				}
			}
		}
	}
	for i := len(area) - 1; i >= 0; i-- {
		for j := len(area) - 2; j >= 0; j-- {
			if area[j+1][i] == area[j][i] {
				area[j+1][i] = area[j+1][i] + area[j][i]
				score += area[j+1][i]
				area[j][i] = 0
			}
		}
	}
	for k := 0; k < len(area); k++ {
		for i := len(area) - 1; i >= 0; i-- {
			for j := len(area) - 2; j >= 0; j-- {
				if area[j+1][i] == 0 && area[j][i] != 0 {
					area[j+1][i], area[j][i] = area[j][i], area[j+1][i]
				}
			}
		}
	}
}

func toRight(area [][]int) {
	for k := 0; k < len(area); k++ {
		for i := len(area) - 1; i >= 0; i-- {
			for j := len(area) - 2; j >= 0; j-- {
				if area[i][j+1] == 0 && area[i][j] != 0 {
					area[i][j+1], area[i][j] = area[i][j], area[i][j+1]
				}
			}
		}
	}
	for i := len(area) - 1; i >= 0; i-- {
		for j := len(area) - 2; j >= 0; j-- {
			if area[i][j+1] == area[i][j] {
				area[i][j+1] = area[i][j+1] + area[i][j]
				score += area[i][j+1]
				area[i][j] = 0
			}
		}
	}
	for k := 0; k < len(area); k++ {
		for i := len(area) - 1; i >= 0; i-- {
			for j := len(area) - 2; j >= 0; j-- {
				if area[i][j+1] == 0 && area[i][j] != 0 {
					area[i][j+1], area[i][j] = area[i][j], area[i][j+1]
				}
			}
		}
	}
}

func toLeft(area [][]int) {
	for k := 0; k < len(area); k++ {
		for i := 0; i < len(area); i++ {
			for j := 1; j < len(area); j++ {
				if area[i][j-1] == 0 && area[i][j] != 0 {
					area[i][j-1], area[i][j] = area[i][j], area[i][j-1]
				}
			}
		}
	}
	for i := 0; i < len(area); i++ {
		for j := 1; j < len(area); j++ {
			if area[i][j-1] == area[i][j] {
				area[i][j-1] = area[i][j-1] + area[i][j]
				score += area[i][j-1]
				area[i][j] = 0
			}
		}
	}

	for k := 0; k < len(area); k++ {
		for i := 0; i < len(area); i++ {
			for j := 1; j < len(area); j++ {
				if area[i][j-1] == 0 && area[i][j] != 0 {
					area[i][j-1], area[i][j] = area[i][j], area[i][j-1]
				}
			}
		}
	}
}
