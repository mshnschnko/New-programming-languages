package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadMatrixFromFile(filepath string) [][]float64 {
	matrix := make([][]float64, 0)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		nums := make([]float64, 0)
		for _, str_num := range strings.Split(row, " ") {
			num, err := strconv.ParseFloat(str_num, 64)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}
		matrix = append(matrix, nums)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	row_count := len(matrix)
	for _, row := range matrix {
		if len(row) != row_count {
			log.Fatal("Matrix musts be square")
		}
	}
	return matrix
}

func det(matrix [][]float64) float64 {
	if len(matrix) == 1 {
		return matrix[0][0]
	}
	if len(matrix) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}
	sum := 0.0

	for k := 0; k < len(matrix); k++ {
		temp_matrix := make([][]float64, len(matrix)-1)
		for i := 1; i < len(matrix); i++ {
			temp_matrix[i-1] = make([]float64, 0)
			for j := 0; j < len(matrix[i]); j++ {
				if j != k {
					temp_matrix[i-1] = append(temp_matrix[i-1], matrix[i][j])
				}
			}
		}
		factor := -1.0
		if (k+1)%2 == 0 {
			factor = 1.0
		}
		sum += factor * matrix[0][k] * det(temp_matrix)
	}
	return sum
}
