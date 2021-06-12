package coffeHandler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rayzalzero/go-sukha/src/helpers"
)

func (a *AppHandler) GetCoffeList(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	search := c.Query("search")
	// filter := c.Query("filter")
	sort := c.Query("sort")

	if page == 0 {
		page = helpers.DefaultPage
	}

	limit, offset := helpers.PaginationPageOffset(page, limit)
	data, count, err := a.Entity.GetListCoffe(c, offset, limit, search, sort)

	if err != nil {
		statusCode = 400
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	pagination := helpers.PaginationRes(page, count, limit)
	params := map[string]interface{}{
		"payload": data,
		"meta":    pagination,
	}

	c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(params))
}
