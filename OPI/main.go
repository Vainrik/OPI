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

	// изменение номер один

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
	//sdfsdf
	// sdfsd
	//ffff
	//sdfsdf
	//sdfsd
	matrix = swapColumns(matrix, 0, 2)
	fmt.Println("Перестановка столбцов 1 и 3:")
	printMatrix(matrix)

	matrix = swapRows(matrix, 1, 2)
	fmt.Println("Перестановка строки 2 и 3:")
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

// sdfsdf
// sdfsdf
// sdf
// Функция для перестановки двух строк матрицы
func swapRows(matrix [4][4]int, i, j int) [4][4]int {
	for k := 0; k < 4; k++ {
		matrix[i][k], matrix[j][k] = matrix[j][k], matrix[i][k]
	}
	return matrix
}

func swapColumns(matrix [4][4]int, i, j int) [4][4]int {
	for k := 0; k < 4; k++ {
		matrix[k][i], matrix[k][j] = matrix[k][j], matrix[k][i]
	}
	return matrix
}

// sdf
// fffff
// sdfsdf
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
