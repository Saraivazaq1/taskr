// Arquivo para definição das rotas para utilização do CRUD

package routes

import (
	"strconv"
	"taskr/backend/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigurarRotas(r *gin.Engine, g *controllers.GerenciadorTarefas) {

	// Rotas

	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	r.GET("/listar", func(c *gin.Context) {
		g.ListarTarefas(c)
	})

	r.POST("/criar", func(c *gin.Context) {
		g.CriarTarefa(c)
	})

	r.DELETE("/deletar", func(c *gin.Context) {
		query := c.Query("id")
		id, _ := strconv.Atoi(query)
		g.DeletarTarefa(c, id)
	})

	r.PUT("/editar", func(c *gin.Context) {
		queryId := c.Query("id")
		descricao := c.Query("descricao")
		id, _ := strconv.Atoi(queryId)
		g.AtualizarTarefas(c, id, descricao)
	})

}
