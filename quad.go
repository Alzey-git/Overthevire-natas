package main

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {

	var wight = x //Длинна прямоугольника
	var hight = y //Высота прямоугольника

	for i := 0; i < hight; i++ { //Строки
		for j := 0; j < wight; j++ { //Столбцы
			if i == 0 && j == 0 || i == 0 && j == wight-1 || i == hight-1 && j == 0 || i == hight-1 && j == wight-1 { //Грани прямоугольника 0
				z01.PrintRune('0')
			} else if i == 0 || i == hight-1 {
				z01.PrintRune('-')
			} else if j == 0 || j == wight-1 {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n') //Переход на новую строку
	}
}

func QuadB(x, y int) {
	var wight = x //Длинна прямоугольника
	var hight = y //Высота прямоугольника

	for i := 0; i < hight; i++ { //Строки
		for j := 0; j < wight; j++ { //Столбцы
			if i == 0 && j == 0 { //Грани прямоугольника /\
				z01.PrintRune(47)
			} else if i == 0 && j == wight-1 {
				z01.PrintRune(92)
			} else if i == hight-1 && j == 0 {
				z01.PrintRune(92)
			} else if i == hight-1 && j == wight-1 {
				z01.PrintRune(47)
			} else if i == 0 || i == hight-1 {
				z01.PrintRune('*')
			} else if j == 0 || j == wight-1 {
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n') //Переход на новую строку
	}
}

func QuadC(x, y int) {
	var wight = x //Длинна прямоугольника
	var hight = y //Высота прямоугольника

	for i := 0; i < hight; i++ { //Строки
		for j := 0; j < wight; j++ { //Столбцы
			if i == 0 && j == 0 { //Грани прямоугольника /\
				z01.PrintRune('A')
			} else if i == 0 && j == wight-1 {
				z01.PrintRune('A')
			} else if i == hight-1 && j == 0 {
				z01.PrintRune('C')
			} else if i == hight-1 && j == wight-1 {
				z01.PrintRune('C')
			} else if i == 0 || i == hight-1 {
				z01.PrintRune('B')
			} else if j == 0 || j == wight-1 {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n') //Переход на новую строку
	}
}

func main() {
	QuadC(10, 10)
	QuadA(10, 10)
	QuadB(10, 10)
}
