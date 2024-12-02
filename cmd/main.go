package main

import (
	"fmt"
	"github.com/nemopss/physics-engine/engine/core"
)

func main() {
	// Создаем два вектора
	v1 := core.Vector{X: 3, Y: 4}
	v2 := core.Vector{X: 1, Y: 2}

	// Пример операций
	fmt.Println("Vector 1:", v1)
	fmt.Println("Vector 2:", v2)
	fmt.Println("Addition:", v1.Add(v2))
	fmt.Println("Subtraction:", v1.Sub(v2))
	fmt.Println("Length of Vector 1:", v1.Length())
	fmt.Println("Normalized Vector 1:", v1.Normalize())
	fmt.Println("Dot Product:", v1.Dot(v2))
}
