package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

func main() {
	boxes := make([]Box, 0)
	for i := 0; i < 100; i++ {
		w := rand.Intn(100)
		boxes = append(boxes, NewBox(w))
	}
	boxes = sortBoxes(boxes)
	tower := build(boxes)
	printTower(tower)
}

type Color int

const (
	red Color = iota
	green
	white
	blue
	pink
	yellow
)

type Box struct {
	Weigth int
	Sides  []int
}

func canPlace(tower []Box, b Box) bool {
	if len(tower) == 0 {
		return true
	}
	if tower[len(tower)-1].Weigth > b.Weigth && tower[len(tower)-1].Sides[5] == b.Sides[0] {
		return true
	}
	return false
}

func build(boxes []Box) []Box {
	result := make([]Box, 0)
	for _, box := range boxes {
		//	firstIndex := box.ButtomSide
		for buttIndex := 0; buttIndex < 6; buttIndex++ {
			if canPlace(result, box) {
				result = append(result, box)
			} else {
				box.Sides = rotate(box.Sides)
			}
		}
	}
	return result
}

func NewBox(weigth int) (b Box) {
	b.Weigth = weigth
	b.Sides = []int{0, 1, 2, 3, 4, 5}
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
		cp = append(cp, b)
	}
	return cp
}

func sortBoxes(boxes []Box) []Box {
	var n = len(boxes)
	for i := 0; i < n; i++ {
		var max = i
		for j := i; j < n; j++ {
			if !boxes[j].isLighter(boxes[max]) {
				max = j
			}
		}
		boxes[i], boxes[max] = boxes[max], boxes[i]
	}
	return boxes
}

func rotate(sides []int) []int {
	r := len(sides) - 1%len(sides)
	sides = append(sides[r:], sides[:r]...)
	return sides
}

func reverse(boxes []Box) []Box {
	for i, j := 0, len(boxes)-1; i < j; i, j = i+1, j-1 {
		boxes[i], boxes[j] = boxes[j], boxes[i]
	}
	return boxes
}
func printTower(boxes []Box) {
	for _, box := range boxes {
		printSide(Color(box.Sides[0]))
		fmt.Println(box.Weigth)
		printSide(Color(box.Sides[5]))
	}
}
func printSide(c Color) {
	switch c {
	case red:
		color.Red(("#########"))
	case white:
		color.White(("#########"))
	case blue:
		color.Blue(("#########"))
	case pink:
		color.Magenta(("#########"))
	case yellow:
		color.Yellow(("#########"))
	case green:
		color.Green(("#########"))
	}
}
func mergeSort(items []Box) []Box {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}
func merge(a []Box, b []Box) []Box {
	final := []Box{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i].isLighter(b[j]) {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
