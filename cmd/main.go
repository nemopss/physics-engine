package main

import (
	"fmt"
	"github.com/nemopss/physics-engine/engine/core"
)

func main() {
	// Создаем физическое тело (Body)
	body := core.NewBody(
		core.Vector{X: 0, Y: 0}, // Начальная позиция
		core.Vector{X: 0, Y: 0}, // Начальная скорость
		5.0,                     // Масса
	)

	// Применяем силу (гравитация)
	gravity := core.Vector{X: 0, Y: -9.8}   // Ускорение свободного падения
	body.ApplyForce(gravity.Mul(body.Mass)) // F = m * g

	// Обновляем состояние объекта на 1 секунду
	deltaTime := 1.0
	body.Update(deltaTime)

	// Печатаем итоговую позицию и скорость
	fmt.Println("Position after 1 second:", body.Position)
	fmt.Println("Velocity after 1 second:", body.Velocity)

}
