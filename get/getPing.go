package get

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPing(c* gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"hello" : "gin Second",
	})
}
