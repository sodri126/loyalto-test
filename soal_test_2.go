package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	MinDice = 1
	MaxDice = 6
)

type Player struct {
	Dice     []int
	TotalOne int
	Point    int
}

func main() {
	n := 3 // total of players
	m := 4 // total of dices

	game := make([]Player, n)

	fmt.Println("Pemain: ", n, "Dadu: ", m)
	fmt.Println("=====================")

	for round := 1; ; round++ {
		totalPlayerExceededDice := 0
		fmt.Printf("Giliran %d lempar dadu:\n", round)
		fmt.Println("=====================")

		// random dice each player
		for i := 0; i < n; i++ {
			game[i].TotalOne = 0
			if round == 1 {
				game[i].Dice = make([]int, m)
			}
			for j := 0; j < len(game[i].Dice); j++ {
				game[i].Dice[j] = randomNo(MinDice, MaxDice)
				if game[i].Dice[j] == 1 {
					game[i].TotalOne++
				}
			}

			formattedDice := fmt.Sprintf("%v", game[i].Dice)
			fmt.Printf("Pemain #%d (%d): %s\n", i+1, game[i].Point, strings.Trim(formattedDice, "[]"))
		}
		fmt.Println("Setelah evaluasi:")

		// evaluation
		for i := 0; i < n; i++ {
			deletedIndex := make([]int, 0)
			for j := 0; j < len(game[i].Dice); j++ {
				if game[i].Dice[j] == 6 {
					game[i].Point++
					deletedIndex = append(deletedIndex, j)
				} else if game[i].Dice[j] == 1 && game[i].TotalOne != 0 {
					move := i + 1
					if move > n-1 {
						move = 0
					}
					game[i].TotalOne--
					game[move].Dice = append(game[move].Dice, game[i].Dice[j])
					deletedIndex = append(deletedIndex, j)
				}
			}

			for k := 0; k < len(deletedIndex); k++ {
				game[i].Dice = deleteSlice(game[i].Dice, deletedIndex[k]-k)
			}

			if len(game[i].Dice) == 0 {
				totalPlayerExceededDice++
			}
		}

		// output evaluation
		for i := 0; i < len(game); i++ {
			formattedDice := fmt.Sprintf("%v", game[i].Dice)
			fmt.Printf("Pemain #%d (%d): %s\n", i+1, game[i].Point, strings.Trim(formattedDice, "[]"))
		}

		fmt.Println("=====================")
		if n-totalPlayerExceededDice == 1 {
			break
		}
	}

	max := 0
	playerIndexHighScore := 0
	for i := 0; i < len(game); i++ {
		if game[i].Point > max {
			max = game[i].Point
			playerIndexHighScore = i + 1
		}
	}

	fmt.Printf("Score tertinggi: %d, Pada pemain #%d", max, playerIndexHighScore)
}

func deleteSlice(data []int, index int) []int {
	return append(data[:index], data[index+1:]...)
}

func randomNo(start, end int) (res int) {
	time.Sleep(10 * time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	res = rand.Intn(end-start+1) + start
	return
}
