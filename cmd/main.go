package main

import (
	"fmt"
	"github.com/nemopss/physics-engine/engine/core"
)

func main() {
	// Создаём физический мир
	world := core.NewWorld(
		core.Vector{X: 0, Y: -9.8}, // Гравитация вниз
		1.0,                        // Шаг времени (1 секунда)
	)

	// Добавляем объекты в мир
	body1 := core.NewBody(core.Vector{X: 0, Y: 20}, core.Vector{X: 0, Y: 0}, 10.0) // Объект на высоте 20

	world.AddBody(body1)

	// Выводим начальное состояние
	fmt.Println("Step 0:")
	fmt.Println("Body 1 Position:", body1.Position, "Velocity:", body1.Velocity)

	// Обновляем мир пошагово и выводим результат после каждого шага
	for step := 1; step <= 4; step++ {
		world.Update()
		fmt.Printf("\nStep %d:\n", step)
		fmt.Println("Body 1 Position:", body1.Position, "Velocity:", body1.Velocity)
	}

}
