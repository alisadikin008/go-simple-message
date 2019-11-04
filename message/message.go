package message

import (
	messageRepo "message-api/repo"

	"github.com/gin-gonic/gin"
)

func PostMessage(contect *gin.Context) {
	message := messageRepo.Message{}
	contect.Bind(&message)
	data, err := messageRepo.PostMessage(message)
	responseData := map[string]interface{}{
		"Data": data,
	}
	if err != nil {
		responseData := map[string]interface{}{
			"Error": err.Error(),
		}
		contect.JSON(200, responseData)
		return
	}

	contect.JSON(200, responseData)
}

func AllMessage(contect *gin.Context) {
	data := messageRepo.AllMessage()
	responseData := map[string]interface{}{
		"Data": data,
	}
	contect.JSON(200, responseData)
}

func RealTime(contect *gin.Context) {

}
