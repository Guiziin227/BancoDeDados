## GLOSSÁRIO<BR>

### 1) DB E A RELAÇÃO C/ SISTEMAS GERENCIADOR DO DB (SGBD)

- Um **banco de dados** é um conjunto organizado de informações armazenadas de forma estruturada. O *SGBD* (Sistema de Gerenciamento de Banco de Dados) é o software responsável por gerenciar e manipular esses dados. Ele facilita operações como consultas, inserções e manutenção da integridade e segurança dos dados. 

A relação entre o banco de dados e o SGBD é a seguinte:

- Banco de Dados (BD): Ele é o repositório onde as informações são armazenadas. Ele pode ser apenas uma coleção de dados, ou pode ter uma estrutura mais complexa dependendo do modelo (relacional, NoSQL, orientado a objetos, etc.).

- Sistema de Gerenciamento de Banco de Dados (SGBD): O SGBD é o software que oferece uma interface entre os usuários e o banco de dados. Ele fornece ferramentas para manipulação dos dados, como realizar consultas (usando SQL, por exemplo), garantir que os dados sejam armazenados de forma eficiente e segura, além de suportar a realização de operações de backup e recuperação.

### 2) MODELO RELACIONAL DE DB

- O **Modelo Relacional de Banco de Dados** organiza dados em tabelas chamadas relações, compostas por linhas (tuplas) e colunas (atributos). Cada tupla representa uma instância de dados, e cada atributo define uma propriedade desses dados. As relações entre diferentes tabelas são estabelecidas por meio de chaves primárias e estrangeiras, garantindo a integridade referencial. Esse modelo permite a manipulação eficiente de dados através de operações como seleção, projeção e junção, fundamentadas na álgebra relacional.

### 3) MICRO SERVIÇO

- Os **microserviços** são uma abordagem arquitetural que divide aplicações em serviços independentes, cada um responsável por uma função específica. Essa estrutura permite que equipes pequenas desenvolvam, testem e implantem serviços de forma autônoma, promovendo maior agilidade e escalabilidade. Cada microserviço comunica-se com os demais por meio de APIs bem definidas, facilitando a integração e manutenção do sistema como um todo. Essa arquitetura é especialmente vantajosa em ambientes de desenvolvimento ágeis e em nuvem, onde a flexibilidade e a escalabilidade são essenciais.

<HR>
<HR>

## Modelagem de um Sistema<BR>
Modelar um sistema envolve criar uma representação abstrata de como o sistema de software funcionará. Pense nisso como esboçar um projeto para uma casa—é sobre entender os componentes e como eles irão interagir. Em computação:
- **Objetivo**: Ajuda os desenvolvedores a visualizar a arquitetura do sistema, fluxos de trabalho e as relações entre diferentes partes (exemplo: interfaces de usuário, servidores, APIs).
- **Como é feito**: Isso geralmente envolve diagramas ou ferramentas como UML (Linguagem de Modelagem Unificada) para mapear coisas como:
  - **Casos de uso**: O que o sistema faz (exemplo: "fazer login", "processar pagamento").
  - **Diagramas de classe**: Os objetos ou entidades (exemplo: "Usuário", "Pedido") e suas relações.
  - **Diagramas de sequência**: O fluxo passo a passo das ações (exemplo: "Usuário clica no botão → Servidor valida → Banco de dados atualiza").
- **Exemplo**: Se você está criando um aplicativo de e-commerce, a modelagem pode mostrar como o cliente, o carrinho de compras e o gateway de pagamento se conectam e o que acontece quando alguém finaliza uma compra.

### Design de um Sistema
O design de um sistema pega o modelo e o transforma em um plano detalhado para a implementação. É menos abstrato e mais sobre o "como"—especificando tecnologias, padrões e regras.
- **Objetivo**: Definir os detalhes—como quais linguagens de programação, frameworks ou hardware você usará, e como o sistema lidará com carga, segurança ou erros.
- **Aspectos principais**:
  - **Arquitetura**: Será um aplicativo monolítico ou microservices? Cliente-servidor ou peer-to-peer?
  - **Componentes**: Dividindo-o em módulos (exemplo: front-end, back-end, banco de dados).
  - **Desempenho**: Qual deve ser a velocidade de resposta? Quantos usuários ele pode suportar?
- **Exemplo**: Para o aplicativo de e-commerce, o design pode especificar o uso de Python para o back-end, React para o front-end, e AWS para hospedagem, com foco em proteger dados de pagamento.

### Modelagem de Banco de Dados
Modelar um banco de dados é sobre planejar como os dados serão organizados, armazenados e acessados. É como projetar o sistema de arquivamento de uma grande biblioteca.
- **Objetivo**: Garantir que os dados sejam estruturados de maneira lógica, evitando duplicação e podendo ser acessados de forma eficiente.
- **Como é feito**: Isso normalmente envolve:
  - **Diagramas de Entidade-Relacionamento (ERD)**: Mostrando tabelas (exemplo: "Clientes", "Pedidos") e como elas se conectam (exemplo: um Cliente tem muitos Pedidos).
  - **Normalização**: Regras para eliminar redundâncias (exemplo: não armazenar o nome de um cliente em cada linha de pedido).
  - **Chaves**: Chaves primárias (IDs únicos) e chaves estrangeiras (links entre tabelas).
- **Exemplo**: Para o aplicativo de e-commerce, você modelaria uma tabela "Usuários" com colunas como ID, nome e email, e uma tabela "Pedidos" ligada à tabela "Usuários" via um ID de usuário.

### Design de Banco de Dados
O design de um banco de dados parte do modelo, especificando os detalhes técnicos para implementação.
- **Objetivo**: Decidir o tipo de banco de dados (exemplo: relacional como MySQL ou NoSQL como MongoDB), definir esquemas exatos e otimizar o desempenho.
- **Aspectos principais**:
  - **Esquema**: Estruturas exatas de tabelas, tipos de dados (exemplo: inteiro, string) e restrições (exemplo: "email deve ser único").
  - **Índices**: Para acelerar buscas (exemplo: indexar a coluna "ID de usuário").
  - **Escalabilidade**: Vai suportar 1.000 usuários ou 1 milhão? Deve shard ou replicar?
- **Exemplo**: Você pode projetar o banco de dados de e-commerce no PostgreSQL com uma tabela "Usuários" (ID: inteiro, nome: varchar(50), email: varchar(100)) e índices nos campos mais pesquisados, como o email.

### Como Eles Se Encaixam
- **Modelagem/design de sistemas**: Foca na visão geral—como o aplicativo ou software funciona como um todo.
- **Modelagem/design de banco de dados**: Foca no armazenamento e recuperação dos dados necessários para o sistema funcionar de maneira eficiente.
- **Na prática**: Você pode modelar um sistema primeiro (exemplo: "O cliente faz um pedido"), depois modelar o banco de dados para suportar isso (exemplo: "Tabela de pedidos vinculada aos clientes"), e então projetar ambos com escolhas tecnológicas específicas.

Resumindo, a modelagem é sobre planejamento e abstração, enquanto o design é sobre detalhes e execução. Juntos, eles garantem que você não está apenas programando de forma aleatória, mas criando algo que funcione bem e dure.
![proj](https://github.com/user-attachments/assets/2446be52-1da1-4757-8bba-24f67b96463e)


