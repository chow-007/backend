package controllers

import (
	"backend/configs"
	"backend/models"
	"backend/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteRight(ctx *gin.Context)  {
	id, _ := strconv.ParseUint(ctx.Param("rightId"), 10, 16)
	if err := services.Service.DeleteRightById(uint16(id)); err != nil{
		returnMsg(ctx, configs.ERROR_DATABASE, nil, "")
		return
	}
	returnMsg(ctx, 200, nil, "")
	return
}

func GetRights(ctx *gin.Context)  {
	rights := services.Service.GetRightList()
	returnMsg(ctx, 200, rights, "")
	return
}

func GetRightsTree(ctx *gin.Context)  {
	rights := services.Service.GetRightList()

	// 一级权限Map
	topMap := make(map[uint16]models.Permission)
	// 二级权限临时存储Map
	secondMap := make(map[uint16]models.Permission)
	// 三级权限临时存储Map
	threeList := make([]models.Permission, 0)
	for _, rightItem := range rights {
		if rightItem.Level == 0 {
			topMap[rightItem.ID] = rightItem // 一级权限
			continue
		} else if rightItem.Level == 1 {
			secondMap[rightItem.ID] = rightItem // 二级权限
		} else {
			threeList = append(threeList, rightItem)

		}

	}

	// 把三级权限放入二级权限的children列表中
	for _, p := range threeList{
		secondP, ok := secondMap[p.PsPid] // 如果是三级权限，找到对应的父级权限（二级权限）
		if ok {
			children := secondP.Children
			children = append(children, p) // 三级权限放入二级权限的children列表中
			secondP.Children = children
			secondMap[p.PsPid] = secondP
		}
	}

	// 把二级权限放入一级权限的children列表中
	for _, p := range secondMap {
		topP, ok := topMap[p.PsPid]
		if ok {
			children := topP.Children
			children = append(children, p)
			topP.Children = children
			topMap[p.PsPid] = topP
		}
	}
	// 把一级权限转换成列表
	res := make([]models.Permission, 0)
	for _, v := range topMap {
		res = append(res, v)
	}

	returnMsg(ctx, 200, res, "")
	return
}
