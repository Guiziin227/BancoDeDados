package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Condominio struct {
	Id       uint
	Nome     string
	Unidades []*Unidade
}

type Unidade struct {
	Id           uint
	Nome         string
	Condominio   *Condominio
	Alugueis     []*Aluguel
	Propriedades []*Propriedade
}

type Pessoa struct {
	Id           uint
	Nome         string
	Propriedades []*Propriedade
	Aluguel      *Aluguel
}

type Propriedade struct {
	Id      uint
	Pessoa  *Pessoa
	Unidade *Unidade
}

type Aluguel struct {
	Id      uint
	Pessoa  *Pessoa
	Unidade *Unidade
	Valor   float32
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	pessoa := &Pessoa{}

	// Cadastrar condomínio
	fmt.Print("Digite o nome do condomínio: ")
	scanner.Scan()
	nomeCondominio := scanner.Text()

	condominio := &Condominio{
		Id:   1,
		Nome: nomeCondominio,
	}

	// Cadastrar unidades
	var quantidadeUnidades int
	fmt.Print("Quantas unidades deseja cadastrar? (mínimo 2): ")
	fmt.Scan(&quantidadeUnidades)

	if quantidadeUnidades < 2 {
		quantidadeUnidades = 2
	}

	//scanner.Scan()

	for i := 0; i < quantidadeUnidades; i++ {
		fmt.Printf("Digite o nome da unidade %d: ", i+1)
		scanner.Scan()
		nomeUnidade := scanner.Text()

		unidade := &Unidade{
			Id:         uint(i + 1),
			Nome:       nomeUnidade,
			Condominio: condominio,
		}

		condominio.Unidades = append(condominio.Unidades, unidade)
	}

	// Cadastrar pessoas
	var pessoas []*Pessoa

	fmt.Print("\nQuantas pessoas deseja cadastrar? ")
	var quantidadePessoas int
	fmt.Scan(&quantidadePessoas)

	for i := 0; i < quantidadePessoas; i++ {
		fmt.Printf("Digite o nome da pessoa %d: ", i+1)
		scanner.Scan()
		nomePessoa := scanner.Text()

		pessoa = &Pessoa{
			Id:   uint(i + 1),
			Nome: nomePessoa,
		}

		pessoas = append(pessoas, pessoa)
	}

	var quantasPropriedades int

	for _, pessoa := range pessoas {
		fmt.Printf("\n%s é proprietário(a) de alguma unidade? (s/n): ", pessoa.Nome)
		scanner.Scan()
		resposta := strings.ToLower(scanner.Text())

		if resposta == "s" {

			fmt.Printf("Quantas unidades %s possui?", pessoa.Nome)
			fmt.Scan(&quantasPropriedades)

			fmt.Println("Unidades disponíveis:")
			for _, unidade := range condominio.Unidades { //Mostras as propriedades dos guris
				fmt.Printf("- ID: %d, Nome: %s\n", unidade.Id, unidade.Nome)
			}

			for i := 0; i < quantasPropriedades; i++ {

				fmt.Printf("Digite o ID da unidade %d que essa pessoa possui: ", i+1)
				scanner.Scan()
				idUnidade, _ := strconv.Atoi(scanner.Text())

				// Encontrar a unidade correspondente
				var unidadePropriedade *Unidade
				for _, unidade := range condominio.Unidades {
					if unidade.Id == uint(idUnidade) {
						unidadePropriedade = unidade
						break
					}
				}

				if unidadePropriedade != nil {
					propriedade := &Propriedade{
						Id:      uint(len(unidadePropriedade.Propriedades) + 1),
						Pessoa:  pessoa,
						Unidade: unidadePropriedade,
					}

					pessoa.Propriedades = append(pessoa.Propriedades, propriedade)
					unidadePropriedade.Propriedades = append(unidadePropriedade.Propriedades, propriedade)
				}
			}
		}
	}

	// Associações de propriedades e aluguéis
	for _, pessoa := range pessoas { //3 pessoas

		fmt.Printf("\n%s está alugando alguma unidade? (s/n): ", pessoa.Nome)
		scanner.Scan()
		resposta := strings.ToLower(scanner.Text())

		if resposta == "s" {

			fmt.Println("Unidades disponíveis para aluguel:")

			var unidadesDisponiveis []*Unidade
			for _, unidade := range condominio.Unidades {
				// Verifica se a unidade tem proprietário
				if len(unidade.Propriedades) > 0 {
					isProprietario := false
					// Verifica se a pessoa já é dona da unidade
					for _, propriedade := range unidade.Propriedades {
						if propriedade.Pessoa == pessoa {
							isProprietario = true
							break
						}
					}

					// Só adiciona unidades que a pessoa não possui e que não tem aluguel registrado
					if !isProprietario && len(unidade.Alugueis) == 0 {
						unidadesDisponiveis = append(unidadesDisponiveis, unidade)
						fmt.Printf("- ID: %d, Nome: %s\n", unidade.Id, unidade.Nome)
					}
				}
			}

			// Verifica se existem unidades disponíveis para aluguel
			if len(unidadesDisponiveis) == 0 {
				fmt.Println("Não há unidades disponíveis para aluguel no momento.")
			} else {
				fmt.Print("Digite o ID da unidade que essa pessoa deseja alugar: ")
				scanner.Scan()
				idUnidade, _ := strconv.Atoi(scanner.Text())

				fmt.Print("Digite o valor do aluguel: ")
				scanner.Scan()
				valorAluguel, _ := strconv.ParseFloat(scanner.Text(), 32)

				// Encontrar a unidade correspondente
				var unidadeAluguel *Unidade
				for _, unidade := range condominio.Unidades {
					if unidade.Id == uint(idUnidade) {
						unidadeAluguel = unidade
						break
					}
				}

				if unidadeAluguel != nil {
					aluguel := &Aluguel{
						Id:      uint(len(unidadeAluguel.Alugueis) + 1),
						Pessoa:  pessoa,
						Unidade: unidadeAluguel,
						Valor:   float32(valorAluguel),
					}

					pessoa.Aluguel = aluguel
					unidadeAluguel.Alugueis = append(unidadeAluguel.Alugueis, aluguel)
				}
			}
		}

	}

	// Exibindo dados cadastrados
	fmt.Printf("\nCondomínio: %s\n", condominio.Nome)
	fmt.Println("Unidades cadastradas:")
	for _, unidade := range condominio.Unidades {

		if len(unidade.Propriedades) > 0 {
			fmt.Printf("- ID: %d, Nome: %s\n", unidade.Id, unidade.Nome)
			fmt.Println("  Proprietários:")
			for _, prop := range unidade.Propriedades {
				fmt.Printf("  - %s\n", prop.Pessoa.Nome)
			}
		}
	}

	fmt.Println("Unidades sem dono:")
	for _, unidade := range condominio.Unidades {
		if len(unidade.Propriedades) == 0 {
			fmt.Printf("ID: %d, Nome: %s\n", unidade.Id, unidade.Nome)
		}
	}

	fmt.Println("Alugueis:")
	for _, unidade := range condominio.Unidades {
		if len(unidade.Alugueis) > 0 {
			fmt.Println("  Inquilinos:")
			for _, alg := range unidade.Alugueis {
				fmt.Printf("  - %s alugou por R$ %.2f - casa: %s\n", alg.Pessoa.Nome, alg.Valor, alg.Unidade.Nome)
			}
		}
	}

}
