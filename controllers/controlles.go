package controlles

import (
	"net/http"

	"github.com/Led-New/Gin-Go/data-base"
	"github.com/Led-New/Gin-Go/models"
	"github.com/gin-gonic/gin"
)

func REtonarAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}
func NameSpace(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Olá, " + nome + ", tudo bem?",
	})
}
func CriarAlunos(c *gin.Context) {
	var alunos models.Aluno
	if err := c.ShouldBindBodyWithJSON(&alunos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&alunos)
	c.JSON(http.StatusOK, alunos)
}
func BuscaIdAlunos(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "ID não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
func DeletaAlunos(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com o sucesso"})
}
func EditaAlunos(c *gin.Context){
	var alunos models.Aluno
	id := c.Params.ByName("id")	
	database.DB.First(&alunos, id)

	if err:= c.ShouldBindJSON(&alunos); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error()})
			return
	}
	database.DB.Model(&alunos).UpdateColumns(alunos)
	c.JSON(http.StatusOK, alunos)
}
func BuscaCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Params.ByName("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "cpf não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
