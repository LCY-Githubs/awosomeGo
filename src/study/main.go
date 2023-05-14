package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 20
	height = 10
)

type point struct {
	x, y int
}

type snake struct {
	head point
	tail []point
}

type game struct {
	snake    snake
	food     point
	score    int
	gameOver bool
}

func main() {
	g := newGame()
	for !g.gameOver {
		g.draw()
		g.update()
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("Game over! Your score is:", g.score)
}

func newGame() game {
	s := snake{
		head: point{x: 0, y: 0},
		tail: []point{},
	}
	f := randomPoint()
	return game{
		snake:    s,
		food:     f,
		score:    0,
		gameOver: false,
	}
}

func (g *game) draw() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == g.snake.head.x && y == g.snake.head.y {
				fmt.Print("O")
			} else if contains(g.snake.tail, point{x: x, y: y}) {
				fmt.Print("o")
			} else if x == g.food.x && y == g.food.y {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("Score:", g.score)
}

func (g *game) update() {
	// Move snake
	g.snake.tail = append([]point{g.snake.head}, g.snake.tail...)
	if len(g.snake.tail) > g.score+1 {
		g.snake.tail = g.snake.tail[:g.score+1]
	}
	switch rand.Intn(4) {
	case 0:
		g.snake.head.y--
	case 1:
		g.snake.head.x++
	case 2:
		g.snake.head.y++
	case 3:
		g.snake.head.x--
	}
	// Check for collision
	if g.snake.head.x < 0 || g.snake.head.x >= width || g.snake.head.y < 0 || g.snake.head.y >= height {
		g.gameOver = true
		return
	}
	if contains(g.snake.tail, g.snake.head) {
		g.gameOver = true
		return
	}
	if g.snake.head.x == g.food.x && g.snake.head.y == g.food.y {
		g.score++
		g.food = randomPoint()
	}
}

func randomPoint() point {
	return point{
		x: rand.Intn(width),
		y: rand.Intn(height),
	}
}

func contains(ps []point, p point) bool {
	for _, q := range ps {
		if p.x == q.x && p.y == q.y {
			return true
		}
	}
	return false
}
