package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Tarefa struct {
	Id        int
	Descricao string
	Feita     bool
}

// Struct com coleção de "tarefas"
type GerenciadorTarefas struct {
	Tarefas   map[int]Tarefa
	Ids       []int
	proximoID int
}

// Instância da struct acima
func novoGerenciadorTarefas() *GerenciadorTarefas {
	return &GerenciadorTarefas{
		Tarefas:   make(map[int]Tarefa), // Uso do map para acesso simplificado de cada tarefa
		Ids:       []int{}, // slice de IDs para organização futura
		proximoID: 0, // automatização de IDs
	}
}

// Início: CRUD

func (g *GerenciadorTarefas) criarTarefa(descricao string) {
	novaTarefa := Tarefa{
		Id:        g.proximoID,
		Descricao: descricao,
		Feita:     false,
	}

	g.Tarefas[novaTarefa.Id] = novaTarefa
	g.Ids = append(g.Ids, novaTarefa.Id)
	g.proximoID++

}

func (g *GerenciadorTarefas) deletarTarefa(id int) {

	delete(g.Tarefas, id)

	for i, j := range g.Ids {
		if j == id {
			g.Ids = append(g.Ids[:i], g.Ids[i+1:]...)
			break
		}
	}
}

func (g *GerenciadorTarefas) listarTarefas() {
	for _, j := range g.Ids {
		tarefa := g.Tarefas[j]
		fmt.Printf("ID: %d, Descrição: %s", tarefa.Id, tarefa.Descricao)
	}
}

func (g *GerenciadorTarefas) atualizarTarefas(id int) {
	tarefa := g.Tarefas[id]
	fmt.Print("Digite a nova descrição da tarefa: ")
	reader := bufio.NewReader(os.Stdin)
	tarefa.Descricao, _ = reader.ReadString('\n')

}

// Fim: CRUD

func main() {

	gerenciador := novoGerenciadorTarefas()
	reader := bufio.NewReader(os.Stdin)

	// Menu interativo no terminal
	for {
		fmt.Println("===== Menu de Tarefas =====")
		fmt.Println("1 - Criar nova tarefa")
		fmt.Println("2 - Listar tarefas")
		fmt.Println("3 - Atualizar tarefa")
		fmt.Println("4 - Deletar tarefa")
		fmt.Println("5 - Sair")
		fmt.Print("Escolha uma opção: ")

		entrada, _ := reader.ReadString('\n')
		opcao := strings.TrimSpace(entrada)

		switch opcao {
		case "1":
			fmt.Print("Digite a descrição da tarefa: ")
			desc, _ := reader.ReadString('\n')
			desc = strings.TrimSpace(desc)
			gerenciador.criarTarefa(desc)

		case "2":
			gerenciador.listarTarefas()

		case "3":
			gerenciador.listarTarefas()
			fmt.Print("Digite o ID da tarefa a marcar como feita: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)
			gerenciador.atualizarTarefas(id)

		case "4":
			gerenciador.listarTarefas()
			fmt.Print("Digite o ID da tarefa a ser deletada: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)
			gerenciador.deletarTarefa(id)

		case "5":
			fmt.Println("Saindo... :)")
			return

		default:
			fmt.Println("Opção inválida.\n")
		}
	}

}
