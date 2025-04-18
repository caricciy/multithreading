# Descrição do Projeto

Este projeto é um exemplo de aplicação cliente-servidor em Golang que demonstra a comunicação entre um cliente e duas APIs de busca de endereços (ViaCEP e BrasilAPI) para um dado CEP.

A aplicação faz requisições HTTP para as APIs e exibe o endereço encontrado no console, utilizando concorrência com goroutines e canais para buscar as informações de maneira paralela. O projeto também utiliza o contexto (context) para gerenciar o tempo de execução das requisições e garantir que o processo não ultrapasse o limite de tempo definido (1 segundo).

>Nota: Este projeto é uma demonstração e não segue todas as boas práticas arquiteturais.

## Funcionalidade

- O programa recebe um CEP fixo (`01153000`), e simultaneamente realiza requisições para duas APIs externas (ViaCEP e BrasilAPI).
- O programa usa goroutines para fazer as requisições de forma concorrente, o que permite que as duas APIs sejam chamadas ao mesmo tempo.
- A resposta mais rápida é exibida, e caso haja erro ou timeout, uma mensagem é registrada.

## Estrutura do Código

1. **Tipos de Respostas (`ViaCEPResponse` e `BrasilAPIResponse`)**:
   - `ViaCEPResponse`: Contém os campos de resposta da API ViaCEP.
   - `BrasilAPIResponse`: Contém os campos de resposta da API BrasilAPI.

2. **Funções Principais**:
   - `buscarEnderecoViaCEP`: Realiza a requisição para a API ViaCEP.
   - `buscarEnderecoBrasilAPI`: Realiza a requisição para a API BrasilAPI.
   - `showResponse`: Exibe os dados recebidos com base no tipo de resposta (ViaCEP ou BrasilAPI).

3. **Controle de Tempo (Timeout)**:
   - O tempo de execução das requisições é controlado por um contexto com timeout de 1 segundo, e a resposta é registrada se dentro desse tempo.

## Como Rodar

1. Clone o repositório ou copie o código para sua máquina.
2. Certifique-se de que o Go está instalado. Caso contrário, siga as instruções de instalação no [site oficial do Go](https://golang.org/doc/install).
3. Execute o seguinte comando para rodar o projeto:

```bash
   go run cmd/main.go
```