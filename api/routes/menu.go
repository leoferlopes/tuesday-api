package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leoferlopes/tuesday-api/api/types"
	"net/http"
)

func GetMenu(menu *types.Menu) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, menu)
	}
}
