package controllers

import (
	"net/http"
	"strconv"

	"coffee-delivery/dto"
	"coffee-delivery/models"
	"coffee-delivery/services"

	"github.com/gin-gonic/gin"
)

func CreateCoffee(c *gin.Context) {
	var req dto.CreateCafeDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	tags := make([]models.Tag, 0)
	for _, tagName := range req.Tags {
		tags = append(tags, models.Tag{Nome: tagName})
	}

	cafe := models.Coffee{
		Nome:      req.Nome,
		Tipo:      req.Tipo,
		Preco:     req.Preco,
		Descricao: req.Descricao,
		Tags:      tags,
	}

	if err := services.CreateCoffee(cafe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao criar café"})
		return
	}
	c.JSON(http.StatusCreated, cafe)
}

func GetCoffees(c *gin.Context) {
	cafes, err := services.FindAllCoffees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar cafés"})
		return
	}
	c.JSON(http.StatusOK, cafes)
}

func DeleteCoffee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if err := services.DeleteCoffee(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Café não encontrado"})
		return
	}

	c.Status(http.StatusNoContent)
}
