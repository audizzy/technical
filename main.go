package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var numPlayers, numDice int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&numPlayers)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&numDice)

	playerDice := make(map[int]int)

	for i := 0; i < numPlayers; i++ {
		playerDice[i] = numDice
	}

	playerDice1 := make(map[int][]int)
	point := make(map[int]int)
	totalPlayerDice := len(playerDice)
	temp := []int{}
	for {

		for i := 0; i < len(playerDice); i++ {
			var h []int
			for j := 0; j < playerDice[i]; j++ {
				dice := rand.Intn(6) + 1
				h = append(h, dice)
			}
			playerDice1[i] = h
		}
		// evaluasi
		for k := 0; k < len(playerDice); k++ {
			for l := 0; l < len(playerDice1[k]); l++ {
				if playerDice1[k][l] == 6 {
					point[k] += 1
					playerDice[k] -= 1
				}
				if playerDice1[k][l] == 1 {
					if k == len(playerDice)-1 {
						playerDice[0] += 1
						playerDice[len(playerDice)-1] -= 1
					} else {
						playerDice[k+1] += 1
						playerDice[k] -= 1
					}
				}

			}

			if len(temp) != 0 {
				test := false
				for t := 0; t < len(temp); t++ {
					if k == temp[t] {
						test = true
					}
				}
				if !test {
					if playerDice[k] == 0 {
						totalPlayerDice -= 1
						temp = append(temp, k)
					}
				}
			} else {
				if playerDice[k] == 0 {
					totalPlayerDice -= 1
					temp = append(temp, k)
				}
			}
		}

		//cek dadu tiap orang
		fmt.Println(playerDice1)
		fmt.Println("point", point)
		fmt.Println("total player dice", playerDice)
		fmt.Println(totalPlayerDice)
		if totalPlayerDice == 1 {
			break
		}
	}

	maxPoint := 0
	winnerIndex := -1
	for index, value := range point {
		if value > maxPoint {
			maxPoint = value
			winnerIndex = index
		}
	}

	// print the index of pemenang
	fmt.Println("Pemenangnya adalah:", winnerIndex)
}
