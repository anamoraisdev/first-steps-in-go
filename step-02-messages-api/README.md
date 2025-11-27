📌 Introdução

Nessa atividade, trabalhei com JSON dentro de uma API — tanto para enviar quanto para receber dados.

Ao longo da atividade, tive contato com:

-criação de structs com tags JSON,
-leitura do corpo da requisição usando json.NewDecoder,
-envio de respostas em JSON usando json.NewEncoder,
-criação de handlers HTTP para GET e POST,
-tratamento básico de erros ao decodificar JSON,
-testes rápidos usando curl.

🚀 Funcionalidades

🔹 GET /mensagem
Retorna uma mensagem fixa em JSON.

🔹 POST /mensagem
Recebe um JSON enviado pelo cliente e devolve um status + a mensagem recebida.

🏃 Como Rodar
go run .


Acesse:
http://localhost:8080/mensagem

🧪 Exemplos (curl)
GET
curl http://localhost:8080/mensagem

POST
curl -X POST http://localhost:8080/mensagem \
  -H "Content-Type: application/json" \
  -d '{"mensagem":"Olá API!"}'
