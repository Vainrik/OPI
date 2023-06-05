package main

import (
	"fmt"
	"os"
)

func main() {
	// Открытие файла с данными матрицы
	file, err := os.Open("matrix.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Чтение данных из файла в матрицу
	var matrix [4][4]int
	for i := 0; i < 4; i++ {
		_, err := fmt.Fscan(file, &matrix[i][0], &matrix[i][1], &matrix[i][2], &matrix[i][3])
		if err != nil {
			fmt.Println("Ошибка при чтении данных из файла:", err)
			return
		}
	}

	// Вывод исходной матрицы
	fmt.Println("Исходная матрица:")
	printMatrix(matrix)

	// Перестановка строк
	matrix = swapRows(matrix, 1, 3)
	fmt.Println("Перестановка строки 2 и 4:")
	printMatrix(matrix)
	// Перестановка столбцов
	matrix = swapColumns(matrix, 0, 2)
	fmt.Println("Перестановка столбца 1 и 3:")
	printMatrix(matrix)
	// Транспонирование матрицы
	transposed := transposeMatrix(matrix)

	// Вывод преобразованной матрицы
	fmt.Println("Транспонированная матрица:")
	printMatrix(transposed)
}

// Функция для вывода матрицы в консоль
func printMatrix(matrix [4][4]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

// Функция для перестановки двух строк матрицы
func swapRows(matrix [4][4]int, i, j int) [4][4]int {
	for k := 0; k < 4; k++ {
		matrix[i][k], matrix[j][k] = matrix[j][k], matrix[i][k]
	}
	return matrix
}

// Функция для перестановки двух столбцов матрицы
func swapColumns(matrix [4][4]int, i, j int) [4][4]int {
	for k := 0; k < 4; k++ {
		matrix[k][i], matrix[k][j] = matrix[k][j], matrix[k][i]
	}
	return matrix
}

// Функция для транспонирования матрицы
func transposeMatrix(matrix [4][4]int) [4][4]int {
	transposed := [4][4]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			transposed[i][j] = matrix[j][i]
		}
	}
	return transposed
}
