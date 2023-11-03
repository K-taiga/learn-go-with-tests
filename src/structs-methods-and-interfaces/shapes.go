package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// goのinterfaceは渡したタイプがインターフェースが要求するものと一致するか暗黙的に解決される
type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Rectangleの構造体にメソッドを定義 レシーバー名は構造体の最初の文字が慣例
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circleの構造体にメソッドを定義
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
