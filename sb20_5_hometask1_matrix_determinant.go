/*
	Задание 1. Подсчёт определителя матрицы

Что нужно сделать - Напишите функцию, вычисляющую определитель матрицы размером 3 × 3.
*/
package main

import (
	"fmt"
)

const rows1 = 3
const cols1 = 3

const rows2 = 2
const cols2 = 2

// Решение:
// разложим матрицу на составляющие
// исходная матрица
// a00 a01 a02
// a10 a11 a12
// a20 a21 a22
// декомпозиция необходима, чтобы разложить матрицу 3х3 на один массив(b), состоящий
// из первой строки исходной матрицы, и три матрицы 2х2(аi)
// где а1 будет иметь вид
// а10 а11
// а20 а21
// а2 будет иметь вид
// а10 а12
// а20 а22
// а3 будет иметь вид
// а11 а12
// а21 а22
//, т.о. формула рассчёта будет иметь вид
// det(m) = b[1]*(diag(a1)-divdiag(a1)) - b[2]*(diag(a2)-divdiag(a2)) + b[3]*(diag(a3)-divdiag(a3))
// где diag - это произведение диагонали от соответствующей матрицы 2х2, а divdiag - произведение "обратной" диагонали той же матрицы
// и эту формулу уже можно проитерировать в цикле

func decomposeOne(m [rows1][cols1]int) (b [cols1]int) {
	for i := 0; i < cols1; i++ {
		b[i] = m[0][i]
	}
	return b
}

func decomposeTwo(m [rows1][cols1]int) (a [rows2][cols2]int) {
	k := 0
	for i := 1; i <= cols2; i++ {
		for j := 0; j < rows2; j++ {
			a[k][j] = m[i][j]
		}
		k++
	}
	return a
}

func decomposeThree(m [rows1][cols1]int) (a [rows2][cols2]int) {
	k := 0
	for i := 1; i <= cols2; i++ {
		l := 0
		for j := 0; j < rows2; j++ {
			a[k][j] = m[i][l]
			l += 2
		}
		k++
	}
	return a
}

func decomposeFour(m [rows1][cols1]int) (a [rows2][cols2]int) {
	k := 0
	for i := 1; i <= cols2; i++ {
		l := 0
		for j := 1; j <= rows2; j++ {
			a[k][l] = m[i][j]
			l++
		}
		k++
	}
	return a
}

func diagonal(A [rows2][cols2]int) int { // формула подсчета произведения диагонали производной матрицы
	product := A[0][0] * A[1][1]
	return product
}

func diverseDiagonal(A [rows2][cols2]int) int { //формула подсчета произведения обратной дигонали производной матрицы
	product := A[0][1] * A[1][0]
	return product
}

func determinantMatrix(m [rows1][cols1]int) int {
	//циклическая формула подсчета детерминанта
	n := 1
	determinant := 0
	var a [rows2][cols2]int
	for i := 0; i < rows1; i++ {
		B := decomposeOne(m)
		if i == 0 {
			a = decomposeTwo(m)
		} else if i == 1 {
			a = decomposeThree(m)
		} else {
			a = decomposeFour(m)
		}
		determinant += B[i] * (diagonal(a) - diverseDiagonal(a)) * n
		n *= -1
	}
	return determinant
}

func main() {
	matrix := [rows1][cols1]int{
		{9, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(determinantMatrix(matrix))
}
