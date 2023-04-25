package bonus

import (
	"fmt"
	"sort"
)

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {
	// Seu código aqui

	sort.Ints(barLengths) // ordena as barras em ordem crescente
	alturaMaxima := 0
	numTorres := 0

	for i := 0; i < len(barLengths); {
		h := 1
		for j := i + 1; j < len(barLengths); j++ {
			if barLengths[j] == barLengths[i] {
				h++
				i = j
			}
		}
		alturaMaxima = max(alturaMaxima, h)
		numTorres++
		i++
	}
	return alturaMaxima, numTorres
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	barras := []int{6, 5, 6, 7}
	alturaMaxima, numTorres := CalculateTowers(barras)
	fmt.Printf("%d altura máxima\n%d torres\n", alturaMaxima, numTorres)
}
