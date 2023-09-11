package main

import "fmt"

type SnakeGame struct {
	width, height int
	queue, food [][]int
	score int
	number int
}

func Constructor353(width int, height int, food [][]int) SnakeGame {
	var queue [][]int
	queue = append(queue, []int{0, 0})
	return SnakeGame{width: width, height: height, queue: queue, food: food, score: 0, number: 0}
}

func (s *SnakeGame) move(direction string) int {
	head := s.queue[0]
	x, y := head[0], head[1]
	switch direction {
	case "U":
		x--
	case "D":
		x++
	case "L":
		y--
	case "R":
		y++
	default:
	}

	if x < 0 || y < 0 || y >= s.width || x >= s.height {
		return -1
	}

	if s.number < len(s.food) && x == s.food[s.number][0] && y == s.food[s.number][1] {
		s.score++
		s.number++
	} else {
		s.queue = s.queue[:len(s.queue) - 1]
	}

	for i := range s.queue {
		if x == s.queue[i][0] && y == s.queue[i][1] {
			return -1
		}
	}

	s.queue = append([][]int{{x, y}}, s.queue...)
	return s.score
}

func main() {
	width, height, food := 3, 2, [][]int{{1,2}, {0,1}}
	s := Constructor353(width, height, food)
	score := s.move("R")
	fmt.Println(score)
	score = s.move("D")
	fmt.Println(score)
	score = s.move("R")
	fmt.Println(score)
	score = s.move("U")
	fmt.Println(score)
	score = s.move("L")
	fmt.Println(score)
	score = s.move("U")
	fmt.Println(score)
}
