package core

import (
	"fmt"
	"math"
)

// Двумерный вектор с координатами X и Y
type Vector struct {
	X float64
	Y float64
}

// Сложение
func (v Vector) Add(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

// Вычитание
func (v Vector) Sub(other Vector) Vector {
	return Vector{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

// Умножение на скаляр
func (v Vector) Mul(scalar float64) Vector {
	return Vector{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

// Деление на скаляр
func (v Vector) Div(scalar float64) Vector {
	if scalar == 0 {
		panic("div by zero")
	}

	return Vector{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

// Модуль (длина) вектора
func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Нормализация вектора
func (v Vector) Normalize() Vector {
	length := v.Length()
	if length == 0 {
		panic("can't ormalize 0 len vector")
	}
	return v.Div(length)
}

// Скалярное произведение
func (v Vector) Dot(other Vector) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Строковое представление
func (v Vector) String() string {
	return fmt.Sprintf("Vector(X: %.2f, Y: %.2f)", v.X, v.Y)
}
