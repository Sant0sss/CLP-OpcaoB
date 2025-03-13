package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func gaussElimination(matrix [][]float64) ([]float64, error) {
	n := len(matrix)
	solution := make([]float64, n)

	
	m := make([][]float64, n)
	for i := range matrix {
		m[i] = make([]float64, len(matrix[i]))
		copy(m[i], matrix[i])
	}

	for i := 0; i < n; i++ {
	
		maxRow := i
		for j := i; j < n; j++ {
			if math.Abs(m[j][i]) > math.Abs(m[maxRow][i]) {
				maxRow = j
			}
		}
		m[i], m[maxRow] = m[maxRow], m[i]

	
		if math.Abs(m[i][i]) < 1e-10 {
			return nil, fmt.Errorf("matriz singular, sem solução única")
		}

		for j := i + 1; j < n; j++ {
			factor := m[j][i] / m[i][i]
			for k := i; k <= n; k++ {
				m[j][k] -= factor * m[i][k]
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		solution[i] = m[i][n]
		for j := i + 1; j < n; j++ {
			solution[i] -= m[i][j] * solution[j]
		}
		solution[i] /= m[i][i]
	}

	return solution, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Número de variáveis: ")
	nStr, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(strings.TrimSpace(nStr))
	if err != nil || n <= 0 {
		fmt.Println("Número inválido de variáveis")
		return
	}

	matrix := make([][]float64, n)
	fmt.Println("Digite a matriz aumentada (linha por linha):")

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(strings.TrimSpace(line))
		
		if len(fields) != n+1 {
			fmt.Println("Número inválido de elementos na linha")
			return
		}
		
		row := make([]float64, n+1)
		for j, field := range fields {
			val, err := strconv.ParseFloat(field, 64)
			if err != nil {
				fmt.Println("Valor não numérico detectado")
				return
			}
			row[j] = val
		}
		matrix[i] = row
	}

	solution, err := gaussElimination(matrix)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("\nSolução:")
	for i, val := range solution {
		fmt.Printf("x[%d] = %.6f\n", i, val)
	}
}
