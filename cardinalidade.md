1 .. 1 
- A escolha é feita pela pessoa de quem referência quem...

1 .. n 
- "N" puxa.
  ![cardi](https://github.com/user-attachments/assets/886f301e-2067-453d-974c-dba8faab01ff)

n .. n
- Cria uma terceira tabela...

----EXEMPLO N..N------
Conceitual
Pessoa 0..n ------ 1..n Veiculo

Lógico
Pessoa(codPessoa,nome,..., cnh)
Veiculo(codVeiculo, placa, tipo)

Locacao(codLocacao,codPessoa,codVeiculo, dataRet, cadaDevo)
        codVeiculo referencia Veiculo
        codPessoa referencia Pessoa
