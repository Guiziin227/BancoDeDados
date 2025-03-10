1 .. 1 
- A escolha é feita pela pessoa de quem referência quem...

1 .. n 
- "N" puxa.
  ![cardi](https://github.com/user-attachments/assets/886f301e-2067-453d-974c-dba8faab01ff)

n .. n
- Cria uma terceira tabela...

----EXEMPLO N..N------
<br>
Conceitual
Pessoa 0..n ------ 1..n Veiculo

Lógico
Pessoa(codPessoa,nome,..., cnh)
Veiculo(codVeiculo, placa, tipo)

Locacao(codLocacao,codPessoa,codVeiculo, dataRet, cadaDevo)
        codVeiculo referencia Veiculo
        codPessoa referencia Pessoa

----EXEMPLO 2-----
Sistema de gestão de jogos, equipes, ranking, clubes da superliga de volei)

Pessoa nome, dataNascimento, apelido, email, sexo, cpf

treinador pessoa, codConselho...
jogadores pessoa, posicao, altura...
preparador pessoa, tipo, codConselho,...

times com nome, lista de patrocinio, clube, endereço, lista de centro de treinamento
![jogos](https://github.com/user-attachments/assets/e020fe84-974f-44d3-998e-663a85d1a137)
