package main

import (
	"math/rand"
	"strconv"
)

/*
* x x x x x
* x x x x x
* x x ? x x
* x x x x x
* x x x x x
 */

type BingoCard struct {
	Numbers []int
}

// for i := range slice {
//   j := rand.Intn(i + 1)
//   slice[i], slice[j] = slice[j], slice[i]
// }

// var b_row = [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
// var i_row = [15]int{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
// var n_row = [15]int{31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45}
// var g_row = [15]int{46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}
// var o_row = [15]int{61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75}

func CreateNewCard() BingoCard {
	b := BingoCard{Numbers: make([]int, 25)}
	for col := range 5 {
		options := rand.Perm(15)
		for row := range 5 {
			if col == 2 && row == 2 {
				b.Numbers[row*5+col] = 0
				continue
			}
			b.Numbers[row*5+col] = (options[row] + 1) + 15*col
		}
	}
	return b
}

func (b *BingoCard) DrawBall(ball int) {
	for idx, val := range b.Numbers {
		if val == ball {
			b.Numbers[idx] = 0
			break
		}
	}
}

func (b BingoCard) HasWonHorizontal() bool {
	var won bool
	for row := range 5 {
		won = true
		for col := range 5 {
			if b.Numbers[row*5+col] != 0 {
				won = false
				break
			}
		}
		if won {
			break
		}
	}
	return won
}
func (b BingoCard) HasWonVertical() bool {
	var won bool
	for col := range 5 {
		won = true
		for row := range 5 {
			if b.Numbers[row*5+col] != 0 {
				won = false
				break
			}
		}
		if won {
			break
		}
	}
	return won
}

func (b BingoCard) ToString() string {
	str := ""
	for idx, a := range b.Numbers {
		if idx != 0 {
			if idx%5 == 0 {
				str += "\n"
			} else {
				str += " "
			}
		}
		str += strconv.Itoa(a)
		if a < 10 {
			str += " "
		}
	}
	return str
}
