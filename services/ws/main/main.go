package main

import (
	"log"
	mba "mb/adapters"
	"net/http"
	"os"
	wsHandlers "ws/handlers"
)

const (
	ENV_API_KEY           = "API_KEY"
	ENV_URL_MESSAGE_QUEUE = "URL_RABBITMQ"
)

func main() {
	log.Println("[INIT] Iniciando servidor...")

	// 🔹 Carregar variáveis de ambiente
	API_KEY := os.Getenv(ENV_API_KEY)
	URL_MESSAGE_QUEUE := os.Getenv(ENV_URL_MESSAGE_QUEUE)

	if API_KEY == "" {
		log.Fatal("[ERRO] Variável de ambiente API_KEY não encontrada")
	}

	if URL_MESSAGE_QUEUE == "" {
		log.Fatal("[ERRO] Variável de ambiente URL_MESSAGE_QUEUE não encontrada")
	}

	log.Printf("[OK] Variáveis carregadas: API_KEY=***, URL_MESSAGE_QUEUE=%s", URL_MESSAGE_QUEUE)

	// 🔹 Inicializar RabbitMQ
	rmqBroker := mba.RabbitMQBroker{}
	rmqHandler, err := rmqBroker.Run(URL_MESSAGE_QUEUE)
	if err != nil {
		log.Fatalf("[ERRO] Falha ao conectar ao RabbitMQ: %v", err)
	}
	log.Println("[OK] Conexão com RabbitMQ estabelecida")

	// 🔹 Configurar rota WebSocket
	routeUp := wsHandlers.RouteUpgrade{
		API_KEY:   &API_KEY,
		HandlerMB: rmqHandler,
	}
	http.HandleFunc("/upgrade", routeUp.Handler)
	log.Println("[ROTA] Rota /upgrade configurada")

	// 🔹 Iniciar servidor HTTP
	port := ":8080"
	log.Printf("[SERVIDOR] Escutando na porta %s...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("[ERRO] Falha ao iniciar servidor: %v", err)
	}
}
