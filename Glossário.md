## GLOSSÁRIO<BR>

1) DB E A RELAÇÃO C/ SISTEMAS GERENCIADOR DO DB (SGBD)

- Um **banco de dados** é um conjunto organizado de informações armazenadas de forma estruturada. O *SGBD* (Sistema de Gerenciamento de Banco de Dados) é o software responsável por gerenciar e manipular esses dados. <br> Ele facilita operações como consultas, inserções e manutenção da integridade e segurança dos dados. 

A relação entre o banco de dados e o SGBD é a seguinte:

- Banco de Dados (BD): Ele é o repositório onde as informações são armazenadas. Ele pode ser apenas uma coleção de dados, ou pode ter uma estrutura mais complexa dependendo do modelo (relacional, NoSQL, orientado a objetos, etc.).

- Sistema de Gerenciamento de Banco de Dados (SGBD): O SGBD é o software que oferece uma interface entre os usuários e o banco de dados. Ele fornece ferramentas para manipulação dos dados, como realizar consultas (usando SQL, por exemplo), garantir que os dados sejam armazenados de forma eficiente e segura, além de suportar a realização de operações de backup e recuperação.

2) MODELO RELACIONAL DE DB

- O **Modelo Relacional de Banco de Dados** organiza dados em tabelas chamadas relações, compostas por linhas (tuplas) e colunas (atributos). Cada tupla representa uma instância de dados, e cada atributo define uma propriedade desses dados. As relações entre diferentes tabelas são estabelecidas por meio de chaves primárias e estrangeiras, garantindo a integridade referencial. Esse modelo permite a manipulação eficiente de dados através de operações como seleção, projeção e junção, fundamentadas na álgebra relacional.

3) MICRO SERVIÇO

Os **microserviços** são uma abordagem arquitetural que divide aplicações em serviços independentes, cada um responsável por uma função específica. Essa estrutura permite que equipes pequenas desenvolvam, testem e implantem serviços de forma autônoma, promovendo maior agilidade e escalabilidade. Cada microserviço comunica-se com os demais por meio de APIs bem definidas, facilitando a integração e manutenção do sistema como um todo. Essa arquitetura é especialmente vantajosa em ambientes de desenvolvimento ágeis e em nuvem, onde a flexibilidade e a escalabilidade são essenciais.
