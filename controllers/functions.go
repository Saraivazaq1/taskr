// Arquivo para criação das funções básicas do controle de tarefas

package controllers

import (
	"github.com/gin-gonic/gin"
	"taskr/models"
)

// Struct com coleção de "tarefas"
type GerenciadorTarefas struct {
	Tarefas   map[int]models.Tarefa
	Ids       []int
	proximoID int
}

// Instância da struct acima
func NovoGerenciadorTarefas() *GerenciadorTarefas {
	return &GerenciadorTarefas{
		Tarefas:   make(map[int]models.Tarefa), // Uso do map para acesso simplificado de cada tarefa
		Ids:       []int{},                     // slice de IDs para organização futura
		proximoID: 0,                           // automatização de IDs
	}
}

// Início: CRUD

func (g *GerenciadorTarefas) CriarTarefa(c *gin.Context, descricao string) {
	novaTarefa := models.Tarefa{
		Id:        g.proximoID,
		Descricao: descricao,
		Feita:     false,
	}

	g.Tarefas[novaTarefa.Id] = novaTarefa
	g.Ids = append(g.Ids, novaTarefa.Id)
	g.proximoID++

	c.JSON(201, novaTarefa)

}

func (g *GerenciadorTarefas) DeletarTarefa(c *gin.Context, id int) {

	delete(g.Tarefas, id)

	for i, j := range g.Ids {
		if j == id {
			g.Ids = append(g.Ids[:i], g.Ids[i+1:]...)
			break
		}
	}

	c.JSON(200, "Deletado")
}

func (g *GerenciadorTarefas) ListarTarefas(c *gin.Context) {
	for _, j := range g.Ids {
		tarefa := g.Tarefas[j]
		c.JSON(200, tarefa)
	}

}

func (g *GerenciadorTarefas) AtualizarTarefas(c *gin.Context, id int, descricao string) {
	tarefa := g.Tarefas[id]
	tarefa.Descricao = descricao

	g.Tarefas[id] = tarefa

	c.JSON(200, tarefa)

}

// Fim: CRUD

func (g *GerenciadorTarefas) MarcarConcluída(c *gin.Context, id int) {
	tarefa := g.Tarefas[id]
	tarefa.Feita = true

	g.Tarefas[id] = tarefa

	c.JSON(200, tarefa)
}
