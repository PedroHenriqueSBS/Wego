package routes

import (
	"net/http"
	"wego-back/routes/controllers"
)

func HandleRquest() {
	http.HandleFunc("GET /", controllers.Home)
	http.HandleFunc("GET /encomenda", controllers.Encomenda)
	http.HandleFunc("GET /veiculo", controllers.Veiculo)
	http.HandleFunc("GET /locais", controllers.Locais)
	http.HandleFunc("GET /mensagem", controllers.Mensagem)

	http.HandleFunc("GET /api/usuario", controllers.Usuario)
	http.HandleFunc("POST /api/criar-conta", controllers.CriarConta)
	http.HandleFunc("POST /api/entrar-conta", controllers.EntrarConta)
	http.HandleFunc("GET /api/mensagens-internas", controllers.MensagemInterna)
	http.HandleFunc("POST /api/criar-rota", controllers.CriarRota)
	http.HandleFunc("POST /api/criar-transporte", controllers.CriarTransporte)
	http.HandleFunc("POST /api/criar-pacote", controllers.CriarPacote)
	http.HandleFunc("POST /api/criar-automovel", controllers.CriarAutomovel)
	//http.HandleFunc(,)

	http.ListenAndServe(":8080", nil)
}
