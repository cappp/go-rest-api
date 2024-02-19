package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type comida struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Sobre string `json:"sobre"`
	Nota  int    `json:"nota"`
}

var comidas = []comida{
	{ID: "1", Nome: "Acarajé", Sobre: "O acarajé é uma especialidade gastronômica das culinárias africana e afro-brasileira. Trata-se de um bolinho feito de massa de feijão-fradinho, cebola e sal, e frito em azeite de dendê. No continente africano é conhecido como akara, e especificamente no norte da Nigéria é também chamado de kosai.", Nota: 10},
	{ID: "2", Nome: "Abará", Sobre: "Abará é um bolinho de feijão-fradinho moído cozido em banho-maria embrulhado em folha de bananeira. É um prato típico da culinária da África e da cozinha baiana. Também faz parte da comida ritual do candomblé.", Nota: 10},
	{ID: "3", Nome: "Pão de queijo", Sobre: "O pão de queijo é uma iguaria oriunda de Minas Gerais, muito difundida em todo o Brasil. Embora não haja registro de local e época exata de sua criação, há consenso de que tenha se originado em Minas Gerais em meados do Século XVIII.", Nota: 8},
	{ID: "4", Nome: "Feijoada", Sobre: "Feijoada é uma designação portuguesa a um prato da culinária transmontana que se popularizou também nos demais países lusófonos como Brasil, Angola, Moçambique, Timor-Leste e Macau. Consiste num guisado de feijão, normalmente com carne, e quase sempre acompanhado com arroz.", Nota: 10},
	{ID: "5", Nome: "Baião de dois", Sobre: "Baião de dois é um prato tipico da região Nordeste. Consiste num preparado de arroz e feijão, de preferência o feijão verde ou feijão novo. É frequente adicionar-se queijo coalho e nata. Não se adiciona carne-seca no Ceará.", Nota: 9},
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"mensagem": "Bem-vindo a API Rest de comida brasileira. Dê uma olhada em /comidas aí."})
	})

	router.GET("/comidas", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, comidas)
	})

	router.GET("/comidas/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		for _, comida := range comidas {
			if comida.ID == id {
				ctx.IndentedJSON(http.StatusOK, comida)
				return
			}
		}

		ctx.IndentedJSON(http.StatusNotFound, gin.H{"mensagem": "Encontrei essa comida aí não rapaz..."})
	})

	router.POST("/comidas", func(ctx *gin.Context) {
		var novaComida comida

		if err := ctx.BindJSON(&novaComida); err != nil {
			return
		}

		comidas = append(comidas, novaComida)
		ctx.IndentedJSON(http.StatusCreated, novaComida)
	})

	router.DELETE("/comidas/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		for index, comida := range comidas {
			if comida.ID == id {
				comidas = append(comidas[:index], comidas[index+1:]...)
				ctx.IndentedJSON(http.StatusOK, comidas)
				return
			}
		}

		ctx.IndentedJSON(http.StatusNotFound, gin.H{"mensagem": "Encontrei essa comida aí não rapaz..."})
	})

	router.Run()
}
