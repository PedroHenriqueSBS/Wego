package action

import (
	"fmt"
	"os"
)

func Erro(area string, detalhe string, erro error) {
	file, err := os.OpenFile("arquivo.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
	}
	defer file.Close()
	resposta := "=============== Aviso ===============\nNa área [ " + area + " ] em [ " + detalhe + " ] - Erro: [ " + erro.Error() + "\n=============== Aviso ===============\n"

	_, err = file.WriteString(resposta)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
	}
}
