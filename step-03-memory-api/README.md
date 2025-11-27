📌 Introdução

Nesta atividade pratiquei o armazenamento de dados diretamente na memória, usando slices para simular um “banco” simples.

Ao longo da atividade, tive contato com:

 - criação de uma struct com ID,
 - geração incremental de IDs,
 - implementação de handlers para GET e POST,
 - uso de slices para armazenar dados em memória,
 - testes rápidos usando curl.

🚀 Funcionalidades

🔹 GET /messages
Retorna todas as mensagens armazenadas no slice.

🔹 POST /messages
Cria uma nova mensagem, gera um ID automaticamente e retorna o objeto criado.

🏃 Como Rodar
go run .

Acesse:
http://localhost:8080/messages

🧪 Exemplos (curl)
Criar mensagem
curl -X POST http://localhost:8080/messages \
  -H "Content-Type: application/json" \
  -d '{"content":"Primeira mensagem"}'

Listar mensagens
curl http://localhost:8080/messages
