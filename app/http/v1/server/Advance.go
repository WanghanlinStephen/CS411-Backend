package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"pro/app/common/response"
	"pro/app/model"
	"pro/app/models"
)

func AdvanceOne(c *gin.Context) {

	result, err := model.AdvanceQueryOneSQL()
	if err != nil {
		response.Error(c, "Fail to execute query 1 with error:"+err.Error())
		return
	}
	jsonResult, err := json.Marshal(result)
	responseData := &models.AdvanceQueryOneOutPut{
		ResponseCode:   200,
		ResponseResult: string(jsonResult),
	}
	response.Success(c, "Success", responseData)
}
