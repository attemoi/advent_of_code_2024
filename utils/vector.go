package utils

type Vector struct {
	X, Y int
}

func (v Vector) Rotate90DegreesClockwise() Vector {
	return Vector{-v.Y, v.X}
}

func (v Vector) Add(other Vector) Vector {
	return Vector{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vector) Substract(other Vector) Vector {
	return Vector{X: v.X - other.X, Y: v.Y - other.Y}
}

func (v Vector) Multiply(multiplier int) Vector {
	return Vector{X: v.X * multiplier, Y: v.Y * multiplier}
}
