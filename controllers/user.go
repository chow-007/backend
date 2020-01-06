/*
@Time : 2019/10/22 23:49
@Author : chenhao
*/
package controllers

import (
	"backend/auth"
	"backend/configs"
	"backend/models"
	"backend/serializers"
	"backend/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

// @Summary 创建用户
// @Description 创建用户，不用传 id 参数
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param user body models.User true "user entity"
// @Success 200 {object} filters.Response {"code":200,"data":nil,"msg":""}
// @Router /user [post]
func CreateUser(ctx *gin.Context) {
	var filter serializers.CreateUserFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}

	if services.Service.IsExistUserName(filter.UserName) {
		returnMsg(ctx, configs.ERROR_DATA_EXIST, "", "user name exist")
		return
	}
	user := models.User{
		UserName:  filter.UserName,
		Password:  filter.Password,
		Email:     filter.Email,
		Mobile:    filter.Mobile,
	}
	if err := services.Service.CreateUser(user); err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, "", "insert failed")
		return
	}
	res := map[string]interface{}{
		"id":          user.ID,
		"username":    user.UserName,
		"mobile":      user.Mobile,
		"email":       user.Email,
		"role_id":     user.RoleID,
		"create_time": user.CreatedAt,
	}
	returnMsg(ctx, 200, res, "success")
}

// @Summary 用户详情
// @Description 用户详情查询
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param userId path string true "user ID"
// @Success 200 {object} filters.Response {"code":200,"data":nil,"msg":""}
// @Router /user/detail/{userId} [get]
func UserDetail(ctx *gin.Context) {
	userId, _ := strconv.ParseUint(ctx.Param("userId"), 10, 64)
	user, _ := services.Service.GetUserById(userId)
	res := map[string]interface{}{
		"id":       user.ID,
		"username": user.UserName,
		"email":    user.Email,
		"mobile":   user.Mobile,
	}
	returnMsg(ctx, 200, res, "")
	return
}

// @Summary 更新用户
// @Description
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param user body models.User true "user entity"
// @Success 200 {object} filters.Response {"code":200,"data":nil,"msg":""}
// @Router /user [put]
func UpdateUser(ctx *gin.Context) {
	var user serializers.UpdateUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, nil, err.Error())
		return
	}

	if err := services.Service.UpdateUserById(user.ID, map[string]interface{}{
		"email":  user.Email,
		"mobile": user.Mobile,
	}); err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, nil, "")
}

func UpdateUserState(ctx *gin.Context) {
	var state serializers.UpdateUser
	if err := ctx.ShouldBindJSON(&state); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, nil, err.Error())
		return
	}
	if err := services.Service.UpdateUserById(state.ID, map[string]interface{}{
		"state": state.State,
	}); err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, nil, "")
}

func UpdateUserRoleId(ctx *gin.Context) {
	var user serializers.UpdateUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, nil, err.Error())
		return
	}
	if err := services.Service.UpdateUserById(user.ID, map[string]interface{}{
		"role_id": user.RoleID,
	}); err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, nil, "")
}

// @Summary 删除用户
// @Description
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param userId path string true "user ID"
// @Success 200 {object} filters.Response {"code":200,"data":nil,"msg":""}
// @Router /user/{userId} [delete]
func DeleteUser(ctx *gin.Context) {
	if err := services.Service.DeleteUser(ctx.Param("userId")); err != nil {
		returnMsg(ctx, 200, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, nil, "")
}

// @Summary 用户列表
// @Description
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param limit path int true "pageSize"
// @Param offset path int true "pageNum"
// @Success 200 {object} filters.Response {"code":200,"data":nil,"msg":""}
// @Router /user/list [get]
func GetUserList(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("pagesize", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("pagenum", "1"))
	query := ctx.DefaultQuery("query", "")
	users, total := services.Service.GetUserList(limit, offset, query)
	userList := make([]map[string]interface{}, 0)
	for _, user := range users {
		userList = append(userList, map[string]interface{}{
			"id":          user.ID,
			"username":    user.UserName,
			"mobile":      user.Mobile,
			"email":       user.Email,
			"create_time": user.CreatedAt,
			"role_name":   user.Role.RoleName,
			"state":       user.State,
		})
	}
	res := map[string]interface{}{
		"total":   total,
		"pagenum": offset,
		"users":   userList,
	}
	returnMsg(ctx, 200, res, "")
}

// @Summary 登录
// @Description
// @Tags Login
// @Produce  json
// @Param user body models.User true "user entity"
// @Success 200 {object} filters.Response {"code":200,"data":nil,"msg":""}
// @Router /login [post]
func Login(ctx *gin.Context) {
	var filter serializers.CreateUserFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	user, err := services.Service.GetUserByNameAndPassword(filter.UserName, filter.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			returnMsg(ctx, configs.ERROR_NOT_FOUND, "", "用户名或密码错误")
			return
		}
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}

	token, err := auth.ObtainToken(user)
	if err != nil {
		returnMsg(ctx, 200, nil, err.Error())
		return
	}

	res := map[string]interface{}{
		"id":       user.ID,
		"role_id":      user.RoleID,
		"username": user.UserName,
		"mobile":   user.Mobile,
		"email":    user.Email,
		"token":    "Bearer" + token,
	}

	returnMsg(ctx, 200, res, "")
}
