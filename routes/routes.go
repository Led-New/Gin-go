package routes

import (
	"github.com/Led-New/Gin-Go/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controlles.REtonarAlunos)
	r.GET("/alunos/:id", controlles.BuscaIdAlunos)
	r.GET("/:nome", controlles.NameSpace)
	r.GET("/alunos/cpf/:cpf", controlles.BuscaCPF)
	r.POST("/alunos", controlles.CriarAlunos)
	r.DELETE("/alunos/:id", controlles.DeletaAlunos)
	r.PATCH("/alunos/:id", controlles.EditaAlunos)
	r.Run()
}
