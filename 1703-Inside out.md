**INSIDE OUT**


Departamento(idDepartamento, nome)<br>
ProcessadorTexto(idProcessadorTexto, desc)<br>
Projeto(idProjeto, desc)<br>
Tipo(idTipo,desc)<br>
Empregado(idEmpregado, nome, cpf, tipo, crea(null),gerente)<br>
	tipo referencia Tipo<br>
	gerente referencia Empregado<br>
Secretaria_ProcessadorTexto(secretaria, processador)<br>
	secretaria referencia Empregado<br>
	processador referencia ProcessadorTexto<br>
Engenheiro_Projeto(engenheiro, projeto)<br>
	engenheiro referencia Empregado<br>
	projeto referencia Projeto<br>
	
