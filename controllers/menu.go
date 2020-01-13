package controllers

import (
	"backend/models"
	"backend/services"
	"github.com/gin-gonic/gin"
	"sort"
)

func GetMenuList(ctx *gin.Context) {
	menus := services.Service.GetMenuList()
	topMenusMap := make(map[uint16]models.Permission)
	topMenusList := models.PermissionSlice{}
	for _, menu := range menus { // 循环所有权限
		if menu.PsPid == 0 { // 没有父级权限，说明是一级权限（菜单）
			topMenusMap[menu.ID] = models.Permission{
				ID:       menu.ID,
				AuthName:   menu.AuthName,
				PsPid:    menu.PsPid,
				Path:      menu.Path,
				Level:    menu.Level,
				Order:    menu.Order,
				Children: []models.Permission{},
			}
			continue
		}
	}
	for _, menu := range menus {
		topMenu, ok := topMenusMap[menu.PsPid] // 找到子级菜单的父菜单，放入父菜单的children里面
		if ok {
			children := topMenu.Children
			children = append(children, menu)
			topMenu.Children = children
			topMenusMap[menu.PsPid] = topMenu
		}
	}

	for _, v := range topMenusMap {
		// 对子菜单进行排序
		sort.Sort(v.Children)
		topMenusList = append(topMenusList, v)
	}
	sort.Sort(topMenusList)

	returnMsg(ctx, 200, topMenusList, "")
	return
}
