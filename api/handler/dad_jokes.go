package api

import (
	"context"
	"errors"
	"hearxtest/model"
	"hearxtest/pkg/dadjoke"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HandlerBackends) GetDadJokeHandler(ctx *gin.Context) {

	c := context.Background()
	joke, err := dadjoke.GetRandom(c, h.DadJokeRepository)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &joke)
}

func (h HandlerBackends) GetJokesPageHandler(ctx *gin.Context) {

	p := ctx.Param("page")
	r := ctx.DefaultQuery("records", "0")

	page, err := strconv.Atoi(p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid page number")})
		return
	}

	records, err := strconv.Atoi(r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid record quantity")})
		return
	}
	c := context.Background()

	joke_page, err := dadjoke.GetPage(c, h.DadJokeRepository, page, records)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &joke_page)
}

func (h HandlerBackends) SubmitJokesHandler(ctx *gin.Context) {
	var jokes []model.DadJoke
	if err := ctx.ShouldBindJSON(&jokes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c := context.Background()

	err := dadjoke.Save(c, h.DadJokeRepository, &jokes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, "success")
}
