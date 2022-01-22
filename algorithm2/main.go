package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MapSize = 8
)

type Node int

const (
	Empty Node = iota
	Queen
	Wall
)

type result struct {
	CountQuens     int
	CountSolotions int
}

func main() {
	var wall string
	var Wall1x, wall1y, wall2x, wall2y int
	fmt.Println("Enter first wall position in x,y format")
	fmt.Scanln(&wall)
	pos := strings.Split(wall, ",")
	Wall1x, _ = strconv.Atoi(pos[0])
	wall1y, _ = strconv.Atoi(pos[1])
	fmt.Println("Enter second wall position in x,y format")
	fmt.Scanln(&wall)
	pos = strings.Split(wall, ",")
	wall2x, _ = strconv.Atoi(pos[0])
	wall2y, _ = strconv.Atoi(pos[1])
	board := createMap(Wall1x, wall1y, wall2x, wall2y)
	printMap(board)
}

func solve() {}
func createMap(wall1x, wall1y, wall2x, wall2y int) [][]int {
	boardMap := make([][]int, MapSize)
	for i := range boardMap {
		boardMap[i] = make([]int, MapSize)
	}
	boardMap[wall1x][wall1y] = int(Wall)
	boardMap[wall2x][wall2y] = int(Wall)
	return boardMap
}

func printMap(board [][]int) {
	for i := 0; i < MapSize; i++ {
		for j := 0; j < MapSize; j++ {
			fmt.Printf("%s ", strconv.Itoa(board[i][j]))
		}
		fmt.Println()
	}
}
