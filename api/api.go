package api

import (
	"github.com/gin-gonic/gin"
	"wb/repository"
)

type Api struct {
	g  *gin.Engine
	db repository.Database
}

func New(db repository.Database) Api {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	return Api{
		g:  router,
		db: db,
	}
}

func (r *Api) Run() error {

	r.g.Use(gin.Logger())
	r.g.Use(gin.ErrorLogger())
	r.g.Use(gin.Recovery())
	r.g.Any("/", r.HandlerHome)
	r.g.Any("/order/:uid", r.HandlerOrderByUID)
	r.g.LoadHTMLFiles("index.html")

	return r.g.Run("localhost:8080")
}
