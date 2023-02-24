package api

import (
	"database/sql"
	api "hearxtest/api/handler"
	"hearxtest/db"

	"github.com/gin-gonic/gin"
)

func InitAPI(dbc *sql.DB) *gin.Engine {

	r := db.PostgresDadJokeRepository{DB: dbc, TableName: "dad_jokes"}
	h := api.HandlerBackends{
		DadJokeRepository: r,
	}
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	// v1 api
	baseGroup := router.Group("/api/v1")
	publicRoutes(baseGroup, h)
	protectedRoutes(baseGroup, h)

	return router
}

func publicRoutes(sg *gin.RouterGroup, h api.HandlerBackends) {

	sg.GET("/jokes", h.GetDadJokeHandler)
	sg.GET("/jokes/:page", h.GetJokesPageHandler)
}

func protectedRoutes(sg *gin.RouterGroup, h api.HandlerBackends) {
	authorized := sg.Group("/")
	authorized.Use(AuthRequired())

	// Card number endpoints
	authorized.POST("/jokes", h.SubmitJokesHandler)

}
