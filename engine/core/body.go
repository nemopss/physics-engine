package core

// Физический объект с массой, позицией и скоростью
type Body struct {
	Position Vector
	Velocity Vector
	Mass     float64
	Force    Vector
}

// Создание нового физического объекта
func NewBody(position, velocity Vector, mass float64) *Body {
	if mass <= 0 {
		panic("mass must be greater than zero")
	}

	return &Body{
		Position: position,
		Velocity: velocity,
		Mass:     mass,
		Force:    Vector{X: 0, Y: 0},
	}
}

// Добавление внешней силы
func (b *Body) ApplyForce(force Vector) {
	b.Force = b.Force.Add(force)
}

// Обновление состояния объекта на основе deltaTime
func (b *Body) Update(deltaTime float64) {
	// Ускорение: a = F / m
	acceleration := b.Force.Div(b.Mass)

	// Обновляем позицию: x = x + v * dt + 0.5 * a * dt^2
	b.Position = b.Position.Add(b.Velocity.Mul(deltaTime)).Add(acceleration.Mul(0.5 * deltaTime * deltaTime))

	// Обновляем скорость: v = v + a * dt
	b.Velocity = b.Velocity.Add(acceleration.Mul(deltaTime))

	// Обнуляем силу после обновления
	b.Force = Vector{X: 0, Y: 0}
}
