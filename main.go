package main

import "fmt"


func dadosPessoais(nome string, idade int) (int, string){
	var condicao string
	if idade >= 18 {
	condicao = "Você é maior de 18 " + nome
	} else {
	condicao = "Você é menor de 18 " + nome
	}
	return idade, condicao
}

func main(){
	idade, condicao := dadosPessoais("Lucas", 19)
	fmt.Println(idade)
	fmt.Println(condicao)

}

