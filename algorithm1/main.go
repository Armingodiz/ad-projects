package main

import (
	"fmt"
	"math/rand"
)

func main() {
	boxes := make([]Box, 0)
	for i := 0; i < 100; i++ {
		w := rand.Intn(100)
		boxes = append(boxes, NewBox(w))
	}
	tower := Tower{
		Boxes: boxes,
	}
	startList := make([]Box, 0)
	build(boxes, startList, startList)
	fmt.Println(tower.Sorted)
}

type color int

const (
	color1 color = iota
	color2
	color3
	color4
	color5
	color6
)

type Box struct {
	Weigth     int
	Sides      []int
	ButtomSide int
}

type Tower struct {
	Boxes  []Box
	Sorted []Box
}

func canPlace(tower []Box, b Box) bool {
	if len(tower) == 0 {
		return true
	}
	if tower[len(tower)-1].Weigth > b.Weigth && b.Sides[b.ButtomSide] == tower[len(tower)-1].Sides[len(tower)-1-tower[len(tower)-1].ButtomSide] {
		return true
	}
	return false
}

func build(boxes, sorted, currentTower []Box) []Box {
	fmt.Println(sorted)
	if len(currentTower) > len(sorted) {
		sorted = getCopy(currentTower)
	}
	if len(sorted) == len(boxes) {
		return sorted
	}
	for _, box := range boxes {
		firstIndex := box.ButtomSide
		for buttIndex := 0; buttIndex < 6; buttIndex++ {
			box.ButtomSide = buttIndex
			if canPlace(currentTower, box) {
				currentTower = append(currentTower, box)
				sorted = build(boxes, sorted, currentTower)
			} else {
				box.ButtomSide = firstIndex
			}
		}
	}
	return sorted
}

func NewBox(weigth int) (b Box) {
	b.Weigth = weigth
	b.Sides = []int{0, 1, 2, 3, 4, 5}
	b.ButtomSide = 5
	return
}

func (b Box) isLighter(d Box) bool {
	return b.Weigth < d.Weigth
}

func getCopy(boxes []Box) []Box {
	cp := make([]Box, 0)
	for _, box := range boxes {
		b := NewBox(box.Weigth)
		b.Sides = box.Sides
		b.ButtomSide = box.ButtomSide
		cp = append(cp, b)
	}
	return cp
}
