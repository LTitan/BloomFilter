package response

import (
	"github.com/gin-gonic/gin"
)

// APP response of app
type APP struct {
	C *gin.Context
}

const (
	ParamsError = 4000
	ServeError  = 5000
	OK          = 2000
)

type ResponseCommon struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response to response
func (a *APP) Response(responseCode, internalCode int, msg string, data interface{}) {
	a.C.JSON(
		responseCode,
		ResponseCommon{
			Code:    internalCode,
			Message: msg,
			Data:    data,
		},
	)
	return
}

// ResponseWithPage .
func (a *APP) ResponseWithPage(responseCode, internalCode int, msg string, data interface{}, page interface{}) {
	a.C.JSON(
		responseCode,
		gin.H{
			"code":    internalCode,
			"message": msg,
			"data":    data,
			"page":    page,
		},
	)
	return
}
