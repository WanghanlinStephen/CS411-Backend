package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pro/app/cache"
	"pro/app/common/response"
	"pro/app/model"
	"pro/app/models"
)

//Visitor Section
func Index(c *gin.Context) {
	redis := cache.RedisInter.Get()
	r, err := redis.Do("Set", "test", 111)
	if err != nil {
		fmt.Println(err)
	}
	response.Success(c, "ok", r)
}

func Delete(c *gin.Context) {
	if err := c.ShouldBind(&models.DeleteSQLInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "wrong parameter")
	}
	dataSource := c.Query("dataSource")
	conditionJson := c.Query("condition")
	condition := model.ConvertStringToMap(conditionJson)
	//execute sql query
	err := model.DeleteSQL(c, dataSource, condition)
	if err != nil {
		response.Error(c, "Fail to delete query with error:"+err.Error())
		return
	}
}
