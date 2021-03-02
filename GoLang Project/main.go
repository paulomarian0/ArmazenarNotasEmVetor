package main

import (
	"fmt"
	"sort"
)

/*
Faça um programa para ler a nota da prova de 15 alunos e armazene num vetor, calcule
e imprima a media geral.
*/

func main() {
	const tamanho = 15
	total := 0.0

	sliceA := []float64{}

	var vetorA [tamanho]float64

	var notaAluno float64

	for i := 0; i < tamanho; i++ {
		fmt.Printf("Digite a nota do aluno %d:", i)
		fmt.Scan(&notaAluno)

		vetorA[i] = notaAluno
		total += vetorA[i]
		sliceA = append(sliceA, vetorA[i])

	}

	media := total / tamanho

	fmt.Println("Nota de todos alunos pela ordem de digitação:", sliceA)
	fmt.Print("\n")
	sort.Float64s(sliceA)

	fmt.Println("Nota de todos os alunos ordenada de forma crescente:", sliceA)
	fmt.Print("\n")

	fmt.Printf("Media dos valores dos vetores: %.2f", media)
	fmt.Print("\n")

}
