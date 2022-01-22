package main

import (
	"fmt"
	"strconv"
	"strings"
)

var directions []int
var MapSize int

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

func init() {
	directions = []int{-1, 0, 1}
	var s int
	fmt.Println("enter the map size:")
	fmt.Scanln(&s)
	MapSize = s
}

func main() {
	board := createMap()
	printMap(board)
	res := &result{}
	isSolotion(board, 0, 0, res)
}

func isSolotion(board [][]int, startingI, startingJ int, res *result) bool {
	if res.CountQuens >= MapSize { // count of queens for this problem has been reached to max
		return true
	}
	for i := startingI; i < MapSize; i++ {
		for j := 0; j < MapSize; j++ {
			if i == startingI && j < startingJ {
				continue // to avoid visiting visited nodes
			}
			if canPut(board, i, j) {
				board[i][j] = int(Queen)
				res.CountQuens += 1
				if isSolotion(board, i, j+1, res) {
					fmt.Println("Solotion number " + strconv.Itoa(res.CountSolotions) + ":")
					res.CountSolotions += 1
					printMap(board)
					fmt.Println("##########################################################")
				}
				board[i][j] = int(Empty)
				res.CountQuens -= 1
			}
		}
	}
	return false
}

func canPut(board [][]int, row, col int) bool {
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

func createMap() [][]int {
	boardMap := make([][]int, MapSize)
	for i := range boardMap {
		boardMap[i] = make([]int, MapSize)
	}
	var w int
	fmt.Println("enter number of walls: ")
	fmt.Scanln(&w)
	for i := 0; i < w; i++ {
		var wall string
		var wall1x, wall1y int
		fmt.Println("Enter first wall position in x,y format")
		fmt.Scanln(&wall)
		pos := strings.Split(wall, ",")
		wall1x, _ = strconv.Atoi(pos[0])
		wall1y, _ = strconv.Atoi(pos[1])
		boardMap[wall1x][wall1y] = int(Wall)
	}
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
