package main

import (
	"main/main/generator"
	"main/main/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CNPJRequest struct {
	CNPJ string `json:"cnpj" binding:"required"`
}

type CPFRequest struct {
	CPF string `json:"cpf" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/generate/cnpj", func(c *gin.Context) {
		cnpj := generator.GeraCNPJ()
		c.JSON(http.StatusOK, gin.H{
			"cnpj": cnpj,
		})
	})

	r.GET("/generate/cpf", func(c *gin.Context) {
		cpf := generator.GeraCPF()
		c.JSON(http.StatusOK, gin.H{
			"cpf": cpf,
		})
	})

	r.POST("/validate/cnpj", func(c *gin.Context) {
		var req CNPJRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "CNPJ is required"})
			return
		}

		valid := validation.ValidaCNPJ(req.CNPJ)

		c.JSON(http.StatusOK, gin.H{
			"cnpj":  req.CNPJ,
			"valid": valid,
		})
	})

	r.POST("/validate/cpf", func(c *gin.Context) {
		var req CPFRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "CPF is required"})
			return
		}

		valid := validation.ValidaCPF(req.CPF)

		c.JSON(http.StatusOK, gin.H{
			"cpf":   req.CPF,
			"valid": valid,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
