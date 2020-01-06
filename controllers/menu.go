package controllers

import (
	"backend/models"
	"backend/services"
	"github.com/gin-gonic/gin"
)

func GetMenuList(ctx *gin.Context) {
	menus := services.Service.GetMenuList()
	topMenusMap := make(map[uint16]models.Permission)
	topMenusList := make([]models.Permission, 0)
	for _, menu := range menus {
		if menu.PsPid == 0 { // 一级菜单
			topMenusMap[menu.ID] = models.Permission{
				ID:       menu.ID,
				AuthName:   menu.AuthName,
				PsPid:    menu.PsPid,
				Path:      menu.Path,
				Level:    menu.Level,
				Children: []models.Permission{},
			}
			continue
		}
	}
	for _, menu := range menus {
		topMenu, ok := topMenusMap[menu.PsPid]
		if ok {
			children := topMenu.Children
			children = append(children, menu)
			topMenu.Children = children
			topMenusMap[menu.PsPid] = topMenu
		}
	}

	for _, v := range topMenusMap {
		topMenusList = append(topMenusList, v)
	}

	returnMsg(ctx, 200, topMenusList, "")
	return
}
