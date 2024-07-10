package middlewares

import (
	"crm-glonass/api/components"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := components.GenerateBaseResponseWithError(nil, false, components.InternalError, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := components.GenerateBaseResponseWithAnyError(nil, false, components.InternalError, err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
