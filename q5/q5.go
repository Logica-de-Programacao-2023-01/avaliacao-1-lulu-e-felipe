package q5

import (
	"fmt"
	"strings"
)

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	// Seu código aqui

	var res strings.Builder
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			// Deleta vogais
			continue
		default:
			// Insere '.' antes de consoantes e converte para minúsculas se necessário
			if 'A' <= c && c <= 'Z' {
				c = c - 'A' + 'a'
			}
			res.WriteRune('.')
			res.WriteRune(c)
		}
	}
	return res.String()
}

func main() {
	var s string
	fmt.Print("Digite uma sequência de caracteres: ")
	fmt.Scan(&s)

	fmt.Println(ProcessString(s))
}

