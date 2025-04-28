package main

import (
	"taskr/backend/controllers"
	"taskr/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	gerenciador := controllers.NovoGerenciadorTarefas()

	// Instância do GIN
	router := gin.Default()
	
	// Definição de arquivos estáticos
	router.Static("/frontend", "./frontend")
	router.LoadHTMLFiles("frontend/index.html")

	router.Static("/script", "./frontend/script")
	router.Static("/style", "./frontend/style")


	routes.ConfigurarRotas(router, gerenciador)
	
	router.Run(":3001")
}
