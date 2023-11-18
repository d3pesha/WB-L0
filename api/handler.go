package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wb/database"
	"wb/nats"
)

func (r *Api) HandlerOrderByUID(ctx *gin.Context) {
	uid := ctx.Params.ByName("uid")
	mem := database.New()

	order, found := nats.CacheInstance.Load(uid)
	if !found {

		orderDB, foundDB := mem.Load(uid)
		if !foundDB {
			ctx.HTML(http.StatusNotFound, "index.html", gin.H{
				"order": "Not found",
			})
			return
		}
		if err := r.db.Save(orderDB); err != nil {
			log.Println("Failed to save order to cache: ", err)
		}
		order = orderDB
	}
	js, err := json.MarshalIndent(order, "", " ")
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"order": "Internal Server Error",
		})
		return
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"order": string(js),
	})
}

func (r *Api) HandlerHome(ctx *gin.Context) {
	r.g.LoadHTMLFiles("index.html")
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"order": "Home page",
	})
}
