// Arquivo para criação das funções básicas do controle de tarefas

package controllers

import (
	"github.com/gin-gonic/gin"
	"slices"
	"taskr/backend/models"
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

func (g *GerenciadorTarefas) CriarTarefa(c *gin.Context) {
	var novaTarefa models.Tarefa

	// Tratamento de erros na requisição
	if err := c.ShouldBindJSON(&novaTarefa); err != nil {
		c.JSON(400, gin.H{"erro": "JSON inválido"})
		return
	}

	novaTarefa.Id = g.proximoID
	novaTarefa.Feita = false

	g.Tarefas[novaTarefa.Id] = novaTarefa
	g.Ids = append(g.Ids, novaTarefa.Id)
	g.proximoID++

	c.JSON(201, novaTarefa)

}

func (g *GerenciadorTarefas) DeletarTarefa(c *gin.Context, id int) {

	// Tratamento de erros na requisição
	if _, existe := g.Tarefas[id]; !existe {
		c.JSON(404, gin.H{"erro": "Tarefa não encontrada"})
		return
	}

	delete(g.Tarefas, id)

	for i, j := range g.Ids {
		if j == id {
			g.Ids = slices.Delete(g.Ids, i, i+1)
			break
		}
	}

	c.JSON(200, "Deletado")
}

func (g *GerenciadorTarefas) ListarTarefas(c *gin.Context) {
	var tarefas []models.Tarefa
	for _, id := range g.Ids {
		tarefas = append(tarefas, g.Tarefas[id])
	}

	c.JSON(200, tarefas)

}

func (g *GerenciadorTarefas) AtualizarTarefas(c *gin.Context, id int, descricao string) {
	tarefa := g.Tarefas[id]
	tarefa.Descricao = descricao

	g.Tarefas[id] = tarefa

	c.JSON(200, tarefa)

}

// Fim: CRUD
