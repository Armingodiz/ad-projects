package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MapSize = 8
)

var directions []int

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
	directions = []int{-1, 0, 1, 2}
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

func isSafe(board [][]int, row, col int) bool {
	for _, rowDirection := range directions {
		for _, colDirection := range directions {
			nextRow := row + rowDirection
			nextCol := col + colDirection
			if rowDirection == 0 && colDirection == 0 { // checking node itself
				if board[nextRow][nextCol] == int(Empty) {
					continue
				} else {
					return false
				}
			}
			for nextRow < MapSize && nextRow >= 0 && nextCol < MapSize && nextCol >= 0 { // checking next node is inside the board
				if board[nextRow][nextCol] == int(Queen) {
					return false
				}
				if board[nextRow][nextCol] == int(Wall) {
					break
				}
				// move to next node
				nextRow += rowDirection
				nextCol += colDirection
			}
		}
	}
	return true
}

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
