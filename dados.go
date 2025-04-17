package main

import "fmt"

func main(){
	alunoIdade := make(map[string]int)
    alunoIdade["Lucas"] = 16
    alunoIdade["Leonardo"] = 16
    alunoIdade["Rafael"] = 16
    alunoIdade["Bruno"] = 16
    fmt.Println("Idade do Lucas", alunoIdade["Lucas"])
	notasAlunos := map[string]float64{
		"Lucas" : 10,
		"Leonardo" : 9.9,
		"Rafael" : 7,
		"Bruno" : 3.5,
	}

	for k,v := range notasAlunos{
		fmt.Printf("%s tirou nota %.1f \n", k, v)
	}


}