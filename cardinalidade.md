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
![jogos](https://github.com/user-attachments/assets/e020fe84-974f-44d3-998e-663a85d1a137)


Patrocinador(codPatrocinador, nome-pk, end, campo)

TipoPessoa(codTipoPessoa,descricao-pk)                   

TipoPessoa 1..1 --- 1..n Pessoa

Pessoa(cpf-pk, nome, dataNascimento, apelido, tipo, licençaTreinador, codigoConselhoPreparador, posicaoJogador, alturaJogador, licencaArbitro)
      tipo referencia TipoPessoa(codTipoPessoa)

Time(cnpj-pk, nome, cidade)

Pessoa_Time(pessoa-pk, time-pk, dataInicio-pk, dataFim)  
    treinador referencia Pessoa(cpf)
    time referencia Time(cnpj)

Jogo(codJogo, data, hora, timeCasa, timeVisitante, arbitroPrincipal)
    timeCasa referencia Time(cnpj)
    timeVisitante referencia Time(cnpj)
    arbitroPrincipal referencia Pessoa(cpf)


