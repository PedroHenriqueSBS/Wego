package database

import (
	"database/sql"
	"wego-back/action"

	_ "modernc.org/sqlite"
)

type User struct {
	ID        int
	Nome      string
	Descricao string
	CPF       string
	Email     string
	Senha     string
	Tipo      int
	Status    int
}

type Automovel struct {
	ID     int
	Tipo   int
	Nome   string
	Modelo string
	Ano    int
	Status int
}

type Local struct {
	ID     int
	Nome   string
	Tipo   int
	Cidade string
	Status int
}

type Rota struct {
	ID           int
	LocalInicial int
	LocalFinal   int
	Distancia    int
	TempoH       int
	Status       int
}

type Transporte struct {
	ID           int
	IdAutomovel  int
	UserReceber  int
	UserEntregar int
	Data         string
	IdRota       int
	Status       int
}

type Pacote struct {
	ID           int
	Nome         string
	Decricao     string
	Peso         float64
	Preco        float64
	IdTransporte int
	Status       int
}

type Mensagem struct {
	ID         int
	IDConversa int
	User       int
	Conteudo   string
}

type Conversa struct {
	ID           int
	PrimeiroUser int
	SegundoUser  int
}

type Token struct {
	Token  string
	IdUser int
	Hora   int
	DMA    string
}

func InitDatabase() {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nome TEXT, 
			descricao TEXT,
			cpf TEXT,
			email TEXT,
			senha TEXT,
			tipo INTEGER,
			status INTEGER
		);
		CREATE TABLE IF NOT EXISTS automovel(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			tipo INTEGER,
			nome TEXT,
			modelo TEXT,
			ano INTEGER,
			status INTEGER
		);
		CREATE TABLE IF NOT EXISTS local(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nome TEXT,
			tipo INTEGER,
			cidade TEXT,
			status INTEGER
		);
		CREATE TABLE IF NOT EXISTS rota(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			'local-inicial' INTEGER,
			'local-final' INTEGER,
			distancia INTEGER,
			'tempo-h' INTEGER,
			status INTEGER,
			FOREIGN KEY ("local-inicial") REFERENCES local ("id"),
			FOREIGN KEY ("local-final") REFERENCES local ("id")
		);
		CREATE TABLE IF NOT EXISTS transporte(
			id INTERGER PRIMARY KEY AUTOINCREMENT,
			'id-automovel' INTEGER,
			'user-receber' INTEGER,
			'user-entregar' INTEGER,
			data TEXT,
			'id-rota' INTEGER,
			status INTEGER,
			FOREIGN KEY ("id-automovel") REFERENCES automovel ("id"),
			FOREIGN KEY ("user-receber") REFERENCES user ("id"),
			FOREIGN KEY ("user-entregar") REFERENCES user ("id"),
			FOREIGN KEY ("id-rota") REFERENCES rota ("id")
		);
		CREATE TABLE IF NOT EXISTS pacote(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nome TEXT,
			decricao TEXT,
			peso REAL,
			preco REAL,
			'id-transporte' INTEGER,
			status INTEGER,
			FOREIGN KEY ("id-transporte") REFERENCES transporte ("id")
		);
		CREATE TABLE IF NOT EXISTS conversa(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			'user-primeiro' INTEGER,
			'user-segundo' INTEGER,
			FOREIGN KEY ("user-primeiro") REFERENCES user ("id"),
			FOREIGN KEY ("user-segundo") REFERENCES user ("id")
		);
		CREATE TABLE IF NOT EXISTS mensagem(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			'id-conversa' INTEGER,
			user INTEGER,
			conteudo TEXT,
			FOREIGN KEY ("user") REFERENCES user ("id"),
			FOREIGN KEY ("id-conversa") REFERENCES conversa ("id")			
		);
		CREATE TABLE IF NOT EXISTS token(
			token PRIMARY KEY,
			user INTEGER,
			hora INTEGER,
			dma TEXT,
			FOREIGN KEY ("user") REFERENCES user ("id")
		);
	`)
}

func CreateUser(nome string, descricao string, cpf string, email string, senha string, tipo int) {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO user(nome, descricao, cpf, email, senha, tipo, status) VALUES (?,?,?,?,?,?,?)", nome, descricao, cpf, email, senha, tipo, 1)
	if err != nil {
		action.Erro("db", "Create User", err)
	}
}

func CreateAutomovel(tipo int, nome string, modelo string, ano int, status int) error {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO autmovel(tipo, nome, modelo, ano, status) VALUES (?,?,?,?,?)", tipo, nome, modelo, ano, status)
	if err != nil {
		action.Erro("db", "Create Automovel", err)
		return err
	}
	return nil
}

func CreatePacote(nome string, descricao string, peso float32, preco float32, idTransporte int, status int) error {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO pacote(nome, descricao, peso, preco, 'id-transporte', status) VALUES (?,?,?,?,?,?)", nome, descricao, peso, preco, idTransporte, status)
	if err != nil {
		action.Erro("db", "Create Pacote", err)
		return err
	}
	return nil
}

func CreateLocal(nome string, tipo int, cidade string, status int) error {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO local(nome, tipo, cidade, status) VALUES (?,?,?,?)", nome, tipo, cidade, status)
	if err != nil {
		action.Erro("db", "Create Local", err)
		return err
	}
	return nil
}

func CreateRota(localInicial int, localFinal int, distancia int, tempoHora int, status int) error {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO rota('local-inicial', 'local-final', distancia, 'tempo-h', status) VALUES (?,?,?,?,?)", localInicial, localFinal, distancia, tempoHora, status)
	if err != nil {
		action.Erro("db", "Create Rota", err)
		return err
	}
	return nil
}

func CreateTransporte(idAutomovel int, userReceber int, userEntregar int, data string, idRota int, status int) error {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO transporte('id-automovel', 'user-receber', 'user-entregar', data, id-rota, status) VALUES (?,?,?,?,?,?)", idAutomovel, userReceber, userEntregar, data, idRota, status)
	if err != nil {
		action.Erro("db", "Create Transporte", err)
		return err
	}
	return nil
}

func CreateConversa(userPrimeiro int, userSegundo int) error {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO conversa('user-primiero', 'user-segundo') VALUES (?,?)", userPrimeiro, userSegundo)
	if err != nil {
		action.Erro("db", "Create Conversa", err)
		return err
	}
	return nil
}

func CreateMensagem(idConversa int, user string, conteudo string) error {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO mensagem('id-conversa', user, conteudo) VALUES (?,?,?,?)", idConversa, user, conteudo)
	if err != nil {
		action.Erro("db", "Create Mensagem", err)
		return err
	}
	return nil
}

func CreateToken(idUser int, token string, atualHora int, atualDMA string) {

	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO mensagem(token, user, hora, dma) VALUES (?,?,?,?)", token, idUser, atualHora, atualDMA)
	if err != nil {
		action.Erro("db", "Create Mensagem", err)
	}
}

func UpdateUser(id int, nome string, descricao string, email string, tipo int, status int) {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE user SET nome = ?, descricao = ?, email = ?, tipo = ?, status = ? WHERE id = ?`, nome, descricao, email, tipo, status, id)
	if err != nil {
		action.Erro("db", "Update User", err)
	}
}

func UpdateAutomovel(id int, status int) {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE automovel SET status = ? WHERE id = ?`, status, id)
	if err != nil {
		action.Erro("db", "Update Automovel", err)
	}
}

func UpdateLocal(id int, nome string, tipo int, cidade string, status int) {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE local SET nome = ?, tipo = ?, cidade = ?, status = ? WHERE id = ?`, nome, tipo, cidade, status, id)
	if err != nil {
		action.Erro("db", "Update Local", err)
	}
}

func UpdateRota(id int, distancia int, tempoHora int, status int) {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE rota SET distancia = ?, "tempo-h" = ?, status = ? WHERE id = ?`, distancia, tempoHora, status, id)
	if err != nil {
		action.Erro("db", "Update Rota", err)
	}
}

func UpdateTransporte(id int, idAutomovel int, userReceber int, userEntregar int, data string, idRota int, status int) {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE transporte SET "id-automovel" = ?, "user-receber" = ?, "user-entregar" = ?, data = ?, "id-rota" = ?, status = ? WHERE id = ?`, idAutomovel, userReceber, userEntregar, data, idRota, status, id)
	if err != nil {
		action.Erro("db", "Update Transporte", err)
	}
}

func UpdateToken(token string, hora int, dma string) {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE user SET hora = ?, dma = ? WHERE token = ?`, hora, dma, token)
	if err != nil {
		action.Erro("db", "Update Token", err)
	}
}

func UpdatePacote(id int, nome string, decricao string, peso float64, preco float64, idTransporte int, status int) {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE pacote SET nome = ?, decricao = ?, peso = ?, preco = ?, "id-transporte" = ?, status = ? WHERE id = ?`, nome, decricao, peso, preco, idTransporte, status, id)
	if err != nil {
		action.Erro("db", "Update Pacote", err)
	}
}

func ReadAllUsers() []User {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, nome, descricao, cpf, email, senha, tipo, status FROM user`)
	if err != nil {
		action.Erro("db", "Read All Users", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Nome, &u.Descricao, &u.CPF, &u.Email, &u.Senha, &u.Tipo, &u.Status)
		if err != nil {
			action.Erro("db", "Scan User", err)
		}
		users = append(users, u)
	}

	return users
}

func ReadAllAutomoveis() []Automovel {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, tipo, nome, modelo, ano, status FROM automovel`)
	if err != nil {
		action.Erro("db", "Read All Automoveis", err)
	}
	defer rows.Close()

	var automoveis []Automovel
	for rows.Next() {
		var a Automovel
		err = rows.Scan(&a.ID, &a.Tipo, &a.Nome, &a.Modelo, &a.Ano, &a.Status)
		if err != nil {
			action.Erro("db", "Scan Automovel", err)
		}
		automoveis = append(automoveis, a)
	}

	return automoveis
}

func ReadAllLocais() []Local {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, nome, tipo, cidade, status FROM local`)
	if err != nil {
		action.Erro("db", "Read All Locais", err)
	}
	defer rows.Close()

	var locais []Local
	for rows.Next() {
		var l Local
		err = rows.Scan(&l.ID, &l.Nome, &l.Tipo, &l.Cidade, &l.Status)
		if err != nil {
			action.Erro("db", "Scan Local", err)
		}
		locais = append(locais, l)
	}

	return locais
}

func ReadAllRotas() []Rota {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, "local-inicial", "local-final", distancia, "tempo-h", status FROM rota`)
	if err != nil {
		action.Erro("db", "Read All Rotas", err)
	}
	defer rows.Close()

	var rotas []Rota
	for rows.Next() {
		var r Rota
		err = rows.Scan(&r.ID, &r.LocalInicial, &r.LocalFinal, &r.Distancia, &r.TempoH, &r.Status)
		if err != nil {
			action.Erro("db", "Scan Rota", err)
		}
		rotas = append(rotas, r)
	}

	return rotas
}

func ReadAllTransportes() []Transporte {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, "id-automovel", "user-receber", "user-entregar", data, "id-rota", status FROM transporte`)
	if err != nil {
		action.Erro("db", "Read All Transportes", err)
	}
	defer rows.Close()

	var transportes []Transporte
	for rows.Next() {
		var t Transporte
		err = rows.Scan(&t.ID, &t.IdAutomovel, &t.UserReceber, &t.UserEntregar, &t.Data, &t.IdRota, &t.Status)
		if err != nil {
			action.Erro("db", "Scan Transporte", err)
		}
		transportes = append(transportes, t)
	}

	return transportes
}

func ReadAllPacotes() []Pacote {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, nome, decricao, peso, preco, "id-transporte", status FROM pacote`)
	if err != nil {
		action.Erro("db", "Read All Pacotes", err)
	}
	defer rows.Close()

	var pacotes []Pacote
	for rows.Next() {
		var p Pacote
		err = rows.Scan(&p.ID, &p.Nome, &p.Decricao, &p.Peso, &p.Preco, &p.IdTransporte, &p.Status)
		if err != nil {
			action.Erro("db", "Scan Pacote", err)
		}
		pacotes = append(pacotes, p)
	}

	return pacotes
}

func ReadToken(token string) Token {
	tokenAtual := Token{Token: token}

	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	err = db.QueryRow(`SELECT user, hora, dma FROM token WHERE token = ?`, token).Scan(&tokenAtual.IdUser, &tokenAtual.Hora, &tokenAtual.DMA)
	if err != nil {
		action.Erro("db", "Read Token", err)
	}
	return tokenAtual
}

func ExisteToken(token string) bool {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	var temToken string
	err = db.QueryRow(`SELECT COALESCE(token, "0") FROM token  WHERE token = ?`, token).Scan(&temToken)
	if err != nil {
		action.Erro("db", "Find Numero Mensagem", err)
	}
	if temToken != "0" {
		return true
	}

	return false
}

func DataAutomovel(IDAutomovel int) Automovel {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	automovel := Automovel{ID: IDAutomovel}
	err = db.QueryRow("SELECT tipo, nome, modelo, ano FROM automovel WHERE id = ?", IDAutomovel).Scan(&automovel.Tipo, &automovel.Nome, &automovel.Modelo, &automovel.Ano)
	if err != nil {
		action.Erro("db", "Find DataAutomovel", err)
	}

	return automovel
}

func DataUser(IDUser int) User {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	usuario := User{ID: IDUser}
	err = db.QueryRow("SELECT nome, descricao, cpf, email, tipo, status FROM automovel WHERE id = ?", IDUser).Scan(&usuario.Nome, &usuario.Descricao, &usuario.CPF, &usuario.Tipo, &usuario.Status)
	if err != nil {
		action.Erro("db", "Find DataAutomovel", err)
	}

	return usuario
}

func AllConversasUser(IDUser int) []Conversa {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	var conversas []Conversa
	rows, err := db.Query("SELECT id, 'user-primeiro', 'user-segundo' FROM  conversa WHERE 'user-primeiro' = ? OR 'user-segundo' = ? ", IDUser, IDUser)
	if err != nil {
		action.Erro("db", "Read All Conversas User", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Conversa
		err = rows.Scan(&c.ID, &c.PrimeiroUser, &c.SegundoUser)
		if err != nil {
			action.Erro("db", "Scan Conversa", err)
		}
		conversas = append(conversas, c)
	}

	return conversas
}

func MensagensConversa(IDConversa int) []Mensagem {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		action.Erro("db", "Open DB", err)
	}
	defer db.Close()

	var mensagens []Mensagem
	rows, err := db.Query("SELECT id, user, conteudo FROM mensagem WHERE 'id-conversa' = ?", IDConversa)
	if err != nil {
		action.Erro("db", "Read All Mensagens Conversa", err)
	}
	defer rows.Close()

	for rows.Next() {
		var m Mensagem
		err = rows.Scan(&m.ID, &m.User, &m.Conteudo)
		if err != nil {
			action.Erro("db", "Scan Conversa", err)
		}
		mensagens = append(mensagens, m)
	}

	return mensagens
}
