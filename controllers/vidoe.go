package controllers

import (
	"backend/services"
	"github.com/gin-gonic/gin"
)

func GetVideoList(ctx *gin.Context)  {
	ret := services.Service.GetVideos()
	tmpMap := make(map[string][]string)
	for _, item := range ret {
		v, ok := tmpMap[item.Code]
		if !ok {
			ls := []string{item.Url}
			tmpMap[item.Code] = ls
			continue
		}
		v = append(v, item.Url)
		tmpMap[item.Code] = v
	}
	res := make([]map[string]interface{}, 0)
	for k,v := range tmpMap{
		tmp := make(map[string]interface{})
		tmp["code"] = k
		tmp["videos"] = v
		res = append(res, tmp)
	}
	returnMsg(ctx, 200, res, "")
}
