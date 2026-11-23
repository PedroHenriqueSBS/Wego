package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"wego-back/action"
	"wego-back/database"
	"wego-back/sistema"
)

// Struct Pergunta

type ValidacaoSimples struct {
	Token string `json:"token"`
	Tipo  int    `json:"tipo"`
}

type PerguntaCriarConta struct {
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
	CPF       string `json:"cpf"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
	Tipo      int    `json:"tipo"`
}

type PerguntaEntrarConta struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type PerguntaMensagens struct {
	ValidacaoSimples
	IDConversa int `json:"id-conversa"`
}

type PerguntaCriarLoacal struct {
	ValidacaoSimples
	Nome   string `json:"nome"`
	Tipo   int    `json:"tipo"`
	Cidade string `json:"cidade"`
}

type PerguntaCriarRota struct {
	ValidacaoSimples
	LocalInicial int `json:"local-inicial"`
	LocalFinal   int `json:"local-final"`
	Distancia    int `json:"distancia"`
	TempoH       int `json:"tempo-h"`
}

type PerguntaCriarTransporte struct {
	ValidacaoSimples
	IdAutomovel  int    `json:"id-automovel"`
	UserReceber  int    `json:"user-receber"`
	UserEntregar int    `json:"user-entregar"`
	Data         string `json:"data"`
	IdRota       int    `json:"id-rota"`
}

type PerguntaCriarPacote struct {
	ValidacaoSimples
	Nome         string  `json:"nome"`
	Decricao     string  `json:"descricao"`
	Peso         float64 `json:"peso"`
	Preco        float64 `json:"preco"`
	IdTransporte int     `json:"id-transporte"`
}

type PerguntaCriarAutomovel struct {
	ValidacaoSimples
	Tipo   int    `json:"tipo"`
	Nome   string `json:"nome"`
	Modelo string `json:"modelo"`
	Ano    int    `json:"ano"`
}

// Struct Resposta

type RespostaSimples struct {
	Status string `json:"status"`
}

type RespostaEncomenda struct {
	RespostaSimples
	Transporte []database.Transporte `json:"transporte"`
}

type RespostaUsuarios struct {
	RespostaSimples
	Usuarios []database.User `json:"usuarios"`
}

type RespostaVeiculos struct {
	RespostaSimples
	Automoveis []database.Automovel `json:"automoveis"`
}

type RespostaLocais struct {
	RespostaSimples
	Locais []database.Local `json:"locais"`
}

type RespostaCriarEntrarConta struct {
	RespostaSimples
	ValidacaoSimples
}

type RespostaMensagens struct {
	RespostaSimples
	Mensagens []database.Mensagem `json:"mensagens"`
}

type RespostaChat struct {
	RespostaSimples
	Chat []database.Conversa `json:"chat"`
}

// Base

func Home(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada ValidacaoSimples
	var dadosSaida RespostaSimples

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "Home", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	IsValido := sistema.ValidarToken(dadosEntrada.Token)
	if IsValido == false {
		action.Erro("controllers", "Validar Token False", errors.New("Token Inválido"))
		dadosSaida.Status = "Erro"
	} else {
		dadosSaida.Status = "Ok"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}

func Encomenda(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada ValidacaoSimples
	var dadosSaida RespostaEncomenda
	token := dadosEntrada.Token

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "Encomenda", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	IsValido := sistema.ValidarToken(token)
	if IsValido == false {
		action.Erro("controllers", "Validar Token False", errors.New("Token Inválido"))
		dadosSaida.Status = "Erro"
	} else {
		dadosSaida.Status = "Ok"
		dadosSaida.Transporte = database.ReadAllTransportes()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)

}
func Veiculo(w http.ResponseWriter, r *http.Request) {

	var dadosEntrada ValidacaoSimples
	var dadosSaida RespostaVeiculos
	token := dadosEntrada.Token

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "Veiculo", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	IsValido := sistema.ValidarToken(token)
	if IsValido == false {
		action.Erro("controllers", "Validar Token False", errors.New("Token Inválido"))
		dadosSaida.Status = "Erro"
	} else {
		dadosSaida.Status = "Ok"
		dadosSaida.Automoveis = database.ReadAllAutomoveis()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}

//API

func Usuario(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada ValidacaoSimples
	var dadosSaida RespostaUsuarios
	token := dadosEntrada.Token

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "Usuarios", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	IsValido := sistema.ValidarToken(token)
	if IsValido == false {
		action.Erro("controllers", "Validar Token False", errors.New("Token Inválido"))
		dadosSaida.Status = "Erro"
	} else {
		dadosSaida.Status = "Ok"
		dadosSaida.Usuarios = database.ReadAllUsers()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}

func Locais(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada ValidacaoSimples
	var dadosSaida RespostaLocais
	token := dadosEntrada.Token

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "Locais", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	IsValido := sistema.ValidarToken(token)
	if IsValido == false {
		action.Erro("controllers", "Validar Token False", errors.New("Token Inválido"))
		dadosSaida.Status = "Erro"
	} else {
		dadosSaida.Status = "Ok"
		dadosSaida.Locais = database.ReadAllLocais()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}

func CriarConta(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada PerguntaCriarConta
	var dadosSaida RespostaCriarEntrarConta

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "CriarConta", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	dadosNovoUsuario, err := sistema.ValidarCriarConta(dadosEntrada.Nome, dadosEntrada.Descricao, dadosEntrada.Senha, dadosEntrada.CPF, dadosEntrada.Email, dadosEntrada.Tipo)

	if err != nil {
		action.Erro("controllers", "Validar Criar Conta", err)
		dadosSaida.Status = "Erro"
	} else {
		dadosSaida.Status = "Ok"
		dadosSaida.Tipo = dadosNovoUsuario.Tipo
		dadosSaida.Token = dadosNovoUsuario.Token
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}

func EntrarConta(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada PerguntaEntrarConta
	var dadosSaida RespostaCriarEntrarConta

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "EntrarConta", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	dadosUsuario, err := sistema.CheckPerson(dadosEntrada.Email, dadosEntrada.Senha)

	if err != nil {
		action.Erro("controllers", "Validar Entar Conta", err)
		dadosSaida.Status = "Erro"
	} else {
		dadosSaida.Status = "Ok"
		dadosSaida.Tipo = dadosUsuario.Tipo
		dadosSaida.Token = dadosUsuario.Token
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}

func Mensagem(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada ValidacaoSimples
	var dadosSaida RespostaChat
	token := dadosEntrada.Token

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "Mensagem", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	IsValido := sistema.ValidarToken(token)
	if IsValido == false {
		action.Erro("controllers", "Validar Token False", errors.New("Token Inválido"))
		dadosSaida.Status = "Erro"
	} else {
		dadosSaida.Status = "Ok"
		tokenInfo := database.ReadToken(token)
		dadosSaida.Chat = database.AllConversasUser(tokenInfo.IdUser)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}

func MensagemInterna(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada PerguntaMensagens
	var dadosSaida RespostaMensagens
	token := dadosEntrada.Token

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "Mensagens Conversa", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	IsValido := sistema.ValidarToken(token)
	if IsValido == false {
		action.Erro("controllers", "Validar Token False", errors.New("Token Inválido"))
		dadosSaida.Status = "Erro"
	} else {
		dadosSaida.Status = "Ok"
		dadosSaida.Mensagens = database.MensagensConversa(dadosEntrada.IDConversa)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}

func CriarPacote(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada PerguntaCriarPacote
	var dadosSaida RespostaSimples
	token := dadosEntrada.Token

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "Receber Criar Pacote", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	IsValido := sistema.ValidarToken(token)
	if IsValido == false {
		action.Erro("controllers", "Validar Token False", errors.New("Token Inválido"))
		dadosSaida.Status = "Erro"
	} else {
		err = database.CreatePacote(dadosEntrada.Nome, dadosEntrada.Decricao, float32(dadosEntrada.Peso), float32(dadosEntrada.Preco), dadosEntrada.IdTransporte, 1)
		if err != nil {
			dadosSaida.Status = "Erro"
			action.Erro("controllers", "Criar Pacote", err)
		} else {
			dadosSaida.Status = "Ok"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}

func CriarRota(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada PerguntaCriarRota
	var dadosSaida RespostaSimples
	token := dadosEntrada.Token

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "Receber Criar Rota", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	IsValido := sistema.ValidarToken(token)
	if IsValido == false {
		action.Erro("controllers", "Validar Token False", errors.New("Token Inválido"))
		dadosSaida.Status = "Erro"
	} else {
		err = database.CreateRota(dadosEntrada.LocalInicial, dadosEntrada.LocalFinal, dadosEntrada.Distancia, dadosEntrada.TempoH, 1)
		if err != nil {
			dadosSaida.Status = "Erro"
			action.Erro("controllers", "Criar Rota", err)
		} else {
			dadosSaida.Status = "Ok"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}

func CriarTransporte(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada PerguntaCriarTransporte
	var dadosSaida RespostaSimples
	token := dadosEntrada.Token

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "Receber Criar Transporte", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	IsValido := sistema.ValidarToken(token)
	if IsValido == false {
		action.Erro("controllers", "Validar Token False", errors.New("Token Inválido"))
		dadosSaida.Status = "Erro"
	} else {
		err = database.CreateTransporte(dadosEntrada.IdAutomovel, dadosEntrada.UserReceber, dadosEntrada.UserEntregar, dadosEntrada.Data, dadosEntrada.IdRota, 1)
		if err != nil {
			dadosSaida.Status = "Erro Edit"
			action.Erro("controllers", "Criar Transporte", err)
		} else {
			dadosSaida.Status = "Ok"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}

func CriarAutomovel(w http.ResponseWriter, r *http.Request) {
	var dadosEntrada PerguntaCriarAutomovel
	var dadosSaida RespostaSimples
	token := dadosEntrada.Token

	err := json.NewDecoder(r.Body).Decode(&dadosEntrada)
	if err != nil {
		action.Erro("controllers", "Receber Criar Automovel", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	IsValido := sistema.ValidarToken(token)
	if IsValido == false {
		action.Erro("controllers", "Validar Token False", errors.New("Token Inválido"))
		dadosSaida.Status = "Erro"
	} else {
		err = database.CreateAutomovel(dadosEntrada.Tipo, dadosEntrada.Nome, dadosEntrada.Modelo, dadosEntrada.Ano, 1)
		if err != nil {
			dadosSaida.Status = "Erro"
			action.Erro("controllers", "Criar Automovel", err)
		} else {
			dadosSaida.Status = "Ok"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dadosSaida)
}
