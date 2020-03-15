package controllers

import (
	"backend/configs"
	"github.com/gin-gonic/gin"
)

func GetAlarmData(ctx *gin.Context) {
	alarmList := make([]configs.Alarm, 0)
	for _, v := range configs.AlarmCache{
		alarmList = append(alarmList, v)
	}
	res := map[string]interface{}{
		"total":   len(alarmList),
		"alarms":   alarmList,
	}
	returnMsg(ctx, 200, res, "success")
}