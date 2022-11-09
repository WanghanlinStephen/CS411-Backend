package server

import (
	"github.com/gin-gonic/gin"
	"pro/app/common/response"
	"pro/app/model"
	"pro/app/models"
)

func Advance(c *gin.Context) {
	number := c.Query("number")
	if number == "1" {
		err := model.AdvanceQueryOneSQL()
		if err != nil {
			response.Error(c, "Fail to execute query 1 with error:"+err.Error())
			return
		}
	} else {
		//todo:补充advance query 2
		return
	}
	responseData := &models.DeleteSQLOutput{
		ResponseCode: 200,
	}
	response.Success(c, "Success", responseData)
}
