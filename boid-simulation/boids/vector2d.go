type Vector2D struct {
	x float64
	y float64
}

func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{ x: v1.x + v2.x, y: v1.y + v2.y }
}

func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{ x: v1.x - v2.x, y: v1.y - v2.y }
}

func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{ x: v1.x * v2.x, y: v1.y * v2.y }
}

func (v1 Vector2D) AddV(d float64) Vector2D {
	return Vector2D{ x: v1.x + v2.x, y: v1.y + v2.y }
}
