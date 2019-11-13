package response

import (
	"github.com/gin-gonic/gin"
)

// APP response of app
type APP struct {
	C *gin.Context
}

// Response to response
func (a *APP) Response(responseCode, internalCode int, msg string, data interface{}) {
	a.C.JSON(
		responseCode,
		gin.H{
			"code":    internalCode,
			"message": msg,
			"data":    data,
		},
	)
	return
}
