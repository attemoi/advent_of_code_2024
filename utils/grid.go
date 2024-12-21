package utils

import "strings"

type Grid[T comparable] struct {
	entities map[Vector]T
	width    int
	height   int
}

func NewGrid[T comparable]() *Grid[T] {
	return &Grid[T]{
		entities: make(map[Vector]T),
		width:    0,
		height:   0,
	}
}

func ParseGrid[T comparable](input string, convert func(c rune) (T, bool)) *Grid[T] {
	grid := NewGrid[T]()
	lines := strings.Split(input, "\n")
	grid.height = len(lines)
	grid.width = len(lines[0])

	for j, l := range lines {
		for i, c := range l {
			if entity, found := convert(c); found {
				grid.Set(Vector{X: i, Y: j}, entity)
			}
		}
	}
	return grid
}

func (g *Grid[T]) Set(pos Vector, value T) {
	g.entities[pos] = value
}

func (g *Grid[T]) SetZeroVal(pos Vector) {
	var value T
	g.entities[pos] = value
}

func (g *Grid[T]) Get(pos Vector) (T, bool) {
	entity, exists := g.entities[pos]
	return entity, exists
}

func (g *Grid[T]) Delete(pos Vector) {
	delete(g.entities, pos)
}

func (g *Grid[T]) IsInBounds(pos Vector) bool {
	return pos.X >= 0 && pos.X < g.width && pos.Y >= 0 && pos.Y < g.height
}

func (g *Grid[T]) GroupByValue() map[T][]Vector {
	grouped := make(map[T][]Vector)
	for key, value := range g.entities {
		grouped[value] = append(grouped[value], key)
	}
	return grouped
}
