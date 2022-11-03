package model

import (
	"fmt"
	"pro/config"
	"testing"
)

func Test_Update(t *testing.T) {
	config.Run()
	if err := Run(); err != nil {
		fmt.Println("数据库链接失败:", err)
		return
	}
	//Based on FYP.connection tables
	dataSource := "connection"
	conditions := make(map[string]string)
	conditions["destination"] = "27"
	updates := make(map[string]string)
	updates["map_id"] = "2"
	UpdateSQL(dataSource, updates, conditions)
}
