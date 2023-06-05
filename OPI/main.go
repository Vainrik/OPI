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

	var matri [4][4]int
	for i := 0; i < 4; i++ {
		_, err := fmt.Fscan(file, &matri[i][0], &matri[i][1], &matri[i][2], &matri[i][3])
		if err != nil {
			fmt.Println("Ошибка при чтении данных из файла:", err)
			return
		}
	}

	// Вывод исходной матрицы
	fmt.Println("Исходная матрица:")
	printMatrix(matri)

	//ffff
	//sdfsdf
	//fsdfsd
	//sdgsdfsdf

	matri = swapColumns(matri, 0, 2)
	fmt.Println("Перестановка столбцов 1 и 3:")
	printMatrix(matri)

	matri = swapRows(matri, 1, 2)
	fmt.Println("Перестановка строки 2 и 3:")
	printMatrix(matri)

	// Транспонирование матрицы
	transposed := transposeMatrix(matri)

	// Вывод преобразованной матрицы
	fmt.Println("Транспонированная матрица:")
	printMatrix(transposed)
}

// Функция для вывода матрицы в консоль
func printMatrix(matri [4][4]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", matri[i][j])
		}
		fmt.Println()
	}
}

// sdfsdf
// sdfsdf
// sdf
// Функция для перестановки двух строк матрицы
func swapRows(matri [4][4]int, i, j int) [4][4]int {
	for k := 0; k < 4; k++ {
		matri[i][k], matri[j][k] = matri[j][k], matri[i][k]
	}
	return matri
}

func swapColumns(matri [4][4]int, i, j int) [4][4]int {
	for k := 0; k < 4; k++ {
		matri[k][i], matri[k][j] = matri[k][j], matri[k][i]
	}
	return matri
}

// sdf
// fffff
// sdfsdf
// Функция для транспонирования матрицы
func transposeMatrix(matri [4][4]int) [4][4]int {
	transposed := [4][4]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			transposed[i][j] = matri[j][i]
		}
	}
	return transposed
}
