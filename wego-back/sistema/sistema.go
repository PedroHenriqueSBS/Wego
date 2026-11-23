package sistema

import (
	"errors"
	"math/rand/v2"
	"strconv"
	"strings"
	"time"
	"wego-back/action"
	"wego-back/database"
)

type Validacao struct {
	Token string
	Tipo  int
}

// DMA é Dia/Mês/Ano
func TimeNow() (int, string) {
	agora := time.Now()
	var horas int = agora.Hour()
	var dma string = agora.Format(("02/Jan/2006"))

	return horas, dma
}

func GerarToken() string {
	letras := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	numElementos := 0
	existeToken := true
	var token strings.Builder

	for existeToken == true {
		for numElementos < 3 {
			token.WriteString(letras[rand.IntN(len(letras))])

			numElementos += 1
		}
		for numElementos < 6 {
			numAleatorio := strconv.Itoa(rand.IntN(10))
			token.WriteString(numAleatorio)

			numElementos += 1
		}

		existeToken = database.ExisteToken(token.String())
	}

	return token.String()
}

func ValidarToken(token string) bool {
	structToken := database.ReadToken(token)
	atualHora, atualDMA := TimeNow()

	if (structToken.Token != token) || (atualHora != (structToken.Hora+1) && atualHora != structToken.Hora) || (structToken.DMA != atualDMA) {
		return false
	}

	AtualizarToken(token)

	return true
}

func AtualizarToken(token string) {
	atualHora, atualDMA := TimeNow()
	database.UpdateToken(token, atualHora, atualDMA)
}

func CheckPerson(email string, senha string) (Validacao, error) {
	users := database.ReadAllUsers()
	existePerson := false
	var person Validacao

	for _, usuario := range users {
		if usuario.Email == email && usuario.Senha == senha {
			existePerson = true
			person.Tipo = usuario.Tipo
		}
	}

	if existePerson == true {
		person.Token = GerarToken()
		return person, nil
	}
	person.Tipo = -1
	return person, errors.New("Pode ser que a Senha ou Email esteja errado. Ou este usuário não exista")
}

func ValidarCriarConta(nome string, descricao string, senha string, cpf string, email string, tipo int) (Validacao, error) {
	users := database.ReadAllUsers()
	podeCriar := true
	var person Validacao

	for _, usuario := range users {
		if (usuario.Email == email) || (usuario.CPF == cpf) {
			podeCriar = false
			return person, errors.New("Usuário já cadastrado tentando criar conta.")
		}
	}
	if podeCriar == true {
		database.CreateUser(nome, descricao, cpf, email, senha, tipo)
	}

	person, err := CheckPerson(email, senha)
	if err != nil {
		action.Erro("sistema", "ValidarCriarConta", err)
		return person, errors.New("Checar a criação do novo usuário.")
	}

	return person, nil
}

func CalculadoraEntrega(IDAutomovel int, peso float32, distanciaPadrao int) (float32, error) {
	distancia := float32(distanciaPadrao)
	dadosAutomoveis := [][]float32{
		{0.54, 720, 10},           //carro
		{0.44, 20, 5},             // moto
		{1.19, 1400, 15},          // van
		{1.91, 3500, 20},          // caminhão
		{125, 103000, 1000},       // avião
		{0.17, 9860, 1000},        // conteiner de navio
		{4034, 240000000, 100000}, /// navio
	}
	automovel := database.DataAutomovel(IDAutomovel)
	var kmPreco float32 = dadosAutomoveis[automovel.Tipo][0]
	var pesoMax float32 = dadosAutomoveis[automovel.Tipo][1]
	if pesoMax < peso {
		return 0.0, errors.New("Peso excedeu o limete do automóvel escolhido")
	}
	var taxaMinima float32 = dadosAutomoveis[automovel.Tipo][2]
	var taxaPeso float32 = 1.0

	if peso >= (0.75 * pesoMax) {
		taxaPeso = 1.20
	} else if peso >= (0.5 * pesoMax) {
		taxaPeso = 1.15
	} else if peso >= (0.25 * pesoMax) {
		taxaPeso = 1.1
	}

	return (((kmPreco * distancia) * taxaPeso) + taxaMinima), nil

}
