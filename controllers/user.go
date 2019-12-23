/*
@Time : 2019/10/22 23:49
@Author : chenhao
*/
package controllers

import (
	"backend/configs"
	"backend/models"
	"backend/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

// @Summary 用户详情
// @Description 用户详情查询
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param userId path string true "user ID"
// @Success 200 {object} filters.Response {"code":200,"data":nil,"msg":""}
// @Router /user/detail/{userId} [get]
func UserDetail(ctx *gin.Context) {

	user, err := services.Service.GetUserById(ctx.Param("userId"))
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			returnMsg(ctx, 200, nil, "success")
			return
		}
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, user, "success")
}

// @Summary 更新用户
// @Description
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param user body models.User true "user entity"
// @Success 200 {object} filters.Response {"code":200,"data":nil,"msg":""}
// @Router /user [put]
func UserUpdate(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, nil, err.Error())
		return
	}
	if len(user.ID) <= 0{
		returnMsg(ctx, configs.ERROR_PARAMS, nil, "require userId")
		return
	}

	err := services.Service.UpdateUserById(user)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, map[string]interface{}{"user_id": user.ID}, "success")
}

// @Summary 创建用户
// @Description 创建用户，不用传 id 参数
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param user body models.User true "user entity"
// @Success 200 {object} filters.Response {"code":200,"data":nil,"msg":""}
// @Router /user [post]
func UserCreate(ctx *gin.Context) {
	//role, _ := ctx.Get("ROLE")
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	if user.Role == 0{
		returnMsg(ctx, configs.ERROR_PARAMS, "", "role field cannot null")
		return
	}

	if services.Service.IsExistUser(user.UserName){
		returnMsg(ctx, configs.ERROR_DATA_EXIST, "", "user name exist")
		return
	}
	if services.Service.CreateUser(user) {
		returnMsg(ctx, configs.ERROR_DATABASE, "", "insert failed")
		return
	}
	returnMsg(ctx, 200, "", "success")
}

// @Summary 删除用户
// @Description
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param userId path string true "user ID"
// @Success 200 {object} filters.Response {"code":200,"data":nil,"msg":""}
// @Router /user/{userId} [delete]
func UserDelete(ctx *gin.Context)  {
	if err := services.Service.DeleteUser(ctx.Param("userId")); err != nil{
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
// @Router /user/list/{limit}/{offset} [get]
func GetUserList(ctx *gin.Context)  {
	limit,_ := strconv.Atoi(ctx.Param("limit"))
	offset,_ := strconv.Atoi(ctx.Param("offset"))
	users := services.Service.GetUserList(limit, offset)
	returnMsg(ctx, 200, users, "")
}

// @Summary 登录
// @Description
// @Tags Login
// @Produce  json
// @Param user body models.User true "user entity"
// @Success 200 {object} filters.Response {"code":200,"data":nil,"msg":""}
// @Router /login [post]
func Login(ctx *gin.Context) {
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	user, err := services.Service.GetUser(input.UserName, input.UserPwd)
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			returnMsg(ctx, configs.ERROR_NOT_FOUND, "", "用户名或密码错误")
			return
		}
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	//token, err := auth.ObtainToken(user)
	//if err != nil{
	//	returnMsg(ctx, 200, map[string]string{"user_id":user.ID, "token":"nil"}, err.Error())
	//	return
	//}

	//returnMsg(ctx, 200, map[string]string{"user_id":user.ID, "token":token}, "")
	returnMsg(ctx, 200, map[string]string{"user_id":user.ID}, "")
}

