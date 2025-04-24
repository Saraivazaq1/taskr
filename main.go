package main

import (
	"github.com/gin-gonic/gin"
	"taskr/backend/controllers"
	"taskr/backend/routes"
)

func main() {

	gerenciador := controllers.NovoGerenciadorTarefas()

	// Instância do GIN
	r := gin.Default()
	

	routes.ConfigurarRotas(r, gerenciador)
	r.Run(":8080")
}
