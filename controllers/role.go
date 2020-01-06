package controllers

import (
	"backend/configs"
	"backend/models"
	"backend/serializers"
	"backend/services"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func CreateRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, nil, err.Error())
		return
	}
	if err := services.Service.CreateRole(role); err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, "")
		return
	}
	returnMsg(ctx, 200, nil, "")
	return
}

func DeleteRole(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("roleId"), 10, 16)
	if err := services.Service.DeleteRole(uint16(id)); err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	services.Service.DeleteRole(uint16(id))
	returnMsg(ctx, 200, nil, "")
	return
}

func UpdateRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, nil, err.Error())
		return
	}
	columns := map[string]interface{}{
		"role_name": role.RoleName,
		"role_desc": role.RoleDesc,
	}

	if err := services.Service.UpdateRoleById(role.ID, columns); err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, "")
		return
	}
	returnMsg(ctx, 200, nil, "")
	return
}

func DeleteRoleRight(ctx *gin.Context) {
	var filter serializers.DeleteRoleRight
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, nil, err.Error())
		return
	}

	// 获取role得到权限
	role, err := services.Service.GetRoleById(filter.RoleId)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	psIds := strings.Split(role.PsIds, ",")
	if len(psIds) == 0 {
		returnMsg(ctx, configs.ERROR_NOT_FOUND, nil, "数据库查不到权限")
		return
	}
	// 删除权限，更新role
	psId := filter.RightIdToStr()
	for i, v := range psIds {
		if v == psId {
			psIds = append(psIds[:i], psIds[i+1:]...)
			services.Service.UpdateRoleById(role.ID, map[string]interface{}{
				"ps_ids": psIds,
			})
			break
		}
	}

	returnMsg(ctx, 200, nil, "")
	return
}

func GetRoleDetail(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("roleId"), 10, 16)
	role, err := services.Service.GetRoleById(uint16(id))
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, "")
		return
	}
	returnMsg(ctx, 200, role, "")
	return
}

func GetRoleList(ctx *gin.Context) {
	roles, err := services.Service.GetRoleList()
	if err != nil {
		returnMsg(ctx, 500, roles, "")
		return
	}

	permissions := services.Service.GetRightList()

	result := make(map[uint16]models.Permission)
	for _, p := range permissions {
		result[p.ID] = p
	}

	for i, r := range roles {
		permissionIds := strings.Split(r.PsIds, ",")

		// 一级权限Map
		topMap := make(map[uint16]models.Permission)
		// 二级权限临时存储Map
		secondMap := make(map[uint16]models.Permission)
		for _, strPId := range permissionIds {
			tmpPId, _ := strconv.ParseUint(strPId, 10, 16)
			pId := uint16(tmpPId)

			permission := result[pId]
			if permission.Level == 0 {
				topMap[pId] = permission // 一级权限
				continue
			} else if permission.Level == 1 {
				secondMap[pId] = permission // 二级权限
			} else {
				secondPermission, ok := secondMap[permission.PsPid] // 如果是三级权限，找到对应的父级权限（二级权限）
				if ok {
					threePermissionList := secondPermission.Children
					threePermissionList = append(threePermissionList, permission) // 三级权限放入二级权限的children列表中
					secondPermission.Children = threePermissionList
					secondMap[permission.PsPid] = secondPermission
				}
			}

		}

		// 把二级权限放入一级权限的children列表中
		for _, p := range secondMap {
			topPermission, ok := topMap[p.PsPid]
			if ok {
				secondPermissionList := topPermission.Children
				secondPermissionList = append(secondPermissionList, p)
				topPermission.Children = secondPermissionList
				topMap[p.PsPid] = topPermission
			}
		}
		// 把一级权限转换成列表
		permissionsList := make([]models.Permission, 0)
		for _, v := range topMap {
			permissionsList = append(permissionsList, v)
		}
		// 把权限列表赋值给role
		r.Children = permissionsList
		roles[i] = r

	}

	returnMsg(ctx, 200, roles, "")
	return
}

func SetRoleRights(ctx *gin.Context) {
	var filter serializers.SetRightsFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, nil, err.Error())
		return
	}
	if err := services.Service.UpdateRoleById(filter.Id, map[string]interface{}{
		"ps_ids": filter.PsIds,
	}); err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, nil, "")
	return
}
