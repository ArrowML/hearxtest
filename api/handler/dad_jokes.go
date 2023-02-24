package api

import (
	"hearxtest/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h HandlerBackends) GetDadJokeHandler(ctx *gin.Context) {

}

func (h HandlerBackends) GetJokesPageHandler(ctx *gin.Context) {

}

func (h HandlerBackends) SubmitJokesHandler(ctx *gin.Context) {
	var jokes []model.DadJoke
	if err := ctx.ShouldBindJSON(&jokes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusCreated, "success")
}
