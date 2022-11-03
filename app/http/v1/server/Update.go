package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pro/app/common/response"
	"pro/app/model"
	"pro/app/models"
)

func Update(c *gin.Context) {
	if err := c.ShouldBind(&models.UpdateSQLInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "wrong parameter")
	}
	dataSource := c.PostForm("dataSource")
	updatesJson := c.PostForm("updates")
	conditionJson := c.PostForm("condition")

	updates := model.ConvertStringToMap(updatesJson)
	condition := model.ConvertStringToMap(conditionJson)

	//execute sql query
	err := model.UpdateSQL(dataSource, updates, condition)
	if err != nil {
		response.Error(c, "Fail to delete query with error:"+err.Error())
		return
	}
	responseData := &models.UpdateSQLOutput{
		ResponseCode: 200,
	}
	response.Success(c, "Success", responseData)
}
