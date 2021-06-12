package coffeHandler

import (
	"github.com/gin-gonic/gin"
	coffeDomain "github.com/rayzalzero/go-sukha/src/domain/coffe"
	"github.com/rayzalzero/go-sukha/src/helpers"
)

type AppHandler struct {
	Entity coffeDomain.Entity
}

func InitCoffeHandler(r *gin.RouterGroup, c coffeDomain.Entity) {
	handler := &AppHandler{
		Entity: c,
	}

	coffe := r.Group("/coffe")
	{
		coffe.GET("/list", handler.GetCoffeList)
	}
}

func (a *AppHandler) Home(c *gin.Context) {
	params := map[string]interface{}{
		"payload": gin.H{"message": "OK", "version": "2"},
		"meta":    gin.H{"message": "OK"},
	}
	c.JSON(200, helpers.OutputAPIResponseWithPayload(params))
}
