package main

import (
	"advent-of-code-2024/common"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	buttonA Cord
	buttonB Cord
	prize   Cord
}

type Cord struct {
	x int
	y int
}

func main() {
	var input *[]string
	input = common.ReadInput("./input/input.txt")
	games := convertInput(*input)

	total := 0
	for _, game := range games {
		buttonAPresses, buttonBPresses := calculateButtonPresses(game.prize.x, game.prize.y, game.buttonA.x, game.buttonA.y, game.buttonB.x, game.buttonB.y)

		if buttonAPresses > 0 && buttonBPresses > 0 && buttonAPresses <= 100 && buttonBPresses <= 100 {
			total += (buttonAPresses * 3) + buttonBPresses
		}
	}

	fmt.Println(total)
}

func calculateButtonPresses(targetX, targetY, moveAX, moveAY, moveBX, moveBY int) (int, int) {
	for a := 0; a <= targetX/moveAX; a++ {
		for b := 0; b <= targetX/moveBX; b++ {
			if a*moveAX+b*moveBX == targetX && a*moveAY+b*moveBY == targetY {
				return a, b
			}
		}
	}
	return -1, -1
}

func convertInput(input []string) []Game {
	var games []Game
	var game Game

	for _, line := range input {

		if strings.HasPrefix(line, "Button A:") || strings.HasPrefix(line, "Button B:") {
			items := strings.Split(line, " ")
			item1 := items[2]
			item1 = strings.Replace(item1, "X+", "", -1)
			item1 = strings.Replace(item1, ",", "", -1)
			item2 := items[3]
			item2 = strings.Replace(item2, "Y+", "", -1)
			x, _ := strconv.ParseInt(item1, 10, 32)
			y, _ := strconv.ParseInt(item2, 10, 32)

			if strings.HasPrefix(line, "Button A:") {
				game.buttonA = Cord{int(x), int(y)}
			}
			if strings.HasPrefix(line, "Button B:") {
				game.buttonB = Cord{int(x), int(y)}
			}
		} else if strings.HasPrefix(line, "Prize:") {
			items := strings.Split(line, " ")
			item1 := items[1]
			item1 = strings.Replace(item1, "X=", "", -1)
			item1 = strings.Replace(item1, ",", "", -1)
			item2 := items[2]
			item2 = strings.Replace(item2, "Y=", "", -1)
			x, _ := strconv.ParseInt(item1, 10, 32)
			y, _ := strconv.ParseInt(item2, 10, 32)

			game.prize = Cord{int(x), int(y)}
		} else {
			games = append(games, game)
			game = Game{
				buttonA: Cord{0, 0},
				buttonB: Cord{0, 0},
				prize:   Cord{0, 0},
			}
		}
	}

	return games
}
