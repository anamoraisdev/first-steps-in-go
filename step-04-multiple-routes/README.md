📌 Introdução

Esta atividade foi feita para praticar o básico do backend em Go — aquele “feijão com arroz” das APIs.
A ideia é montar algumas rotas simples, tudo rodando localmente e armazenado apenas em memória mesmo.

Ao longo da atividade, tive contato com:

- manipulação de slices para salvar e remover dados,
- uso de structs com tags JSON,
- criação de handlers HTTP,
- tratamento básico de erros,
- testes rápidos usando curl.

Caso não se sinta confiante em prosseguir, você pode começar pelas atividades anteriores desse repositório.

🚀 Funcionalidades

A atividade implementa rotas para lidar com usuários, produtos e pedidos, todas usando HTTP e JSON:

🔹 Usuários

GET /users/ — lista usuários

POST /users/ — cria usuário

DELETE /users/{id} — remove usuário

🔹 Produtos

GET /products/ — lista produtos

POST /products/ — cria produto

DELETE /products/{id} — remove produto

🔹 Pedidos

GET /orders/ — lista pedidos

POST /orders/ — cria pedido, calcula total e valida saldo


🏃 Como Rodar

Na pasta da atividade:

go run .

O servidor estará disponível em:

http://localhost:8080

🧪 Exemplo de Teste (curl)
Criar um usuário
curl -X POST http://localhost:8080/users/ \
  -H "Content-Type: application/json" \
  -d '{"name":"Ana", "balance":100}'
