package core

type World struct {
	Bodies    []*Body
	Gravity   Vector
	DeltaTime float64
}

// Физический мир, содержащий объекты и управляющий их состояниями
func NewWorld(gravity Vector, deltaTime float64) *World {
	return &World{
		Bodies:    []*Body{},
		Gravity:   gravity,
		DeltaTime: deltaTime,
	}
}

// Создание нового физического мира
func (w *World) AddBody(body *Body) {
	w.Bodies = append(w.Bodies, body)
}

// Обновление состояния всех объектов в мире
func (w *World) Update() {
	for _, body := range w.Bodies {
		// Применение графитации к каждому объекту
		body.ApplyForce(w.Gravity.Mul(body.Mass))

		// Обновление состояния каждого объекта
		body.Update(w.DeltaTime)

	}
}
