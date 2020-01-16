package controllers

import (
	"backend/configs"
	"backend/serializers"
	"backend/services"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetContainerCpu(ctx *gin.Context) {
	var filter serializers.MonitorFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	if len(filter.Field) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "field can not empty")
		return
	}

	// 字段：usage_percent
	baseSql := "SELECT mean(%s) AS %s FROM docker_container_cpu where time > now() - %s group by time(%s),container_name"
	SQL := fmt.Sprintf(baseSql, filter.Field[0], filter.Field[0], filter.DashboardTime, filter.GetInterval())
	res, err := services.Service.Query(SQL)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, res, "")
	return
}

func GetContainerMemory(ctx *gin.Context) {
	var filter serializers.MonitorFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	if len(filter.Field) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "field can not empty")
		return
	}

	// usage
	baseSql := "SELECT mean(%s) / 1048576 AS mem_%s FROM docker_container_mem where time > now() - %s group by time(%s),container_name"
	SQL := fmt.Sprintf(baseSql, filter.Field[0], filter.Field[0], filter.DashboardTime, filter.GetInterval())
	res, err := services.Service.Query(SQL)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, res, "")
	return
}

func GetContainerNetwork(ctx *gin.Context) {
	var filter serializers.MonitorFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	if len(filter.Field) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "field can not empty")
		return
	}

	// rx_bytes  tx_bytes 接收、发送数据包
	baseSql := "SELECT mean(%s) AS net_%s FROM docker_container_net where time > now() - %s group by time(%s), container_name"
	SQL := fmt.Sprintf(baseSql, filter.Field[0], filter.Field[0], filter.DashboardTime, filter.GetInterval())
	res, err := services.Service.Query(SQL)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, res, "")
	return
}

func GetHosts(ctx *gin.Context) {
	//_, role := utils.GetUserIdAndRole(ctx)
	//if role != models.RAdmin{
	//	returnMsg(ctx, configs.ERROR_FORBIDDEN, nil, "非管理员账户禁止访问")
	//	return
	//}
	dashboardTime := ctx.DefaultQuery("dashboard_time", "10m")

	baseSql := `SHOW TAG VALUES WITH KEY = "host" WHERE TIME > now() - %s`
	SQL := fmt.Sprintf(baseSql, dashboardTime)

	tmpRes, err := services.Service.Query(SQL)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	res := make([]string, 0)
	for _, item := range tmpRes {
		for _, v := range item.Values {
			host := v[1].(string)
			if !utils.IsContainStr(res, host) {
				res = append(res, host)
			}
		}
	}
	returnMsg(ctx, 200, map[string][]string{"host": res}, "")
	return
}

func GetContainerMax(ctx *gin.Context) {
	//_, role := utils.GetUserIdAndRole(ctx)
	//if role != models.RAdmin{
	//	returnMsg(ctx, configs.ERROR_FORBIDDEN, nil, "非管理员账户禁止访问")
	//	return
	//}
	var filter serializers.MonitorFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	if len(filter.HostName) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "host_name can not null")
		return
	}
	if len(filter.Field) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "field can not empty")
		return
	}

	// n_containers 多少个containers
	baseSql := "SELECT max(%s) AS max_%s FROM docker where time > now() - %s AND host = '%s' group by time(%s), host"
	SQL := fmt.Sprintf(baseSql, filter.Field[0], filter.Field[0], filter.DashboardTime, filter.HostName, filter.GetInterval())
	res, err := services.Service.Query(SQL)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, res, "")
	return
}

func GetSystemCpu(ctx *gin.Context) {
	//_, role := utils.GetUserIdAndRole(ctx)
	//if role != models.RAdmin{
	//	returnMsg(ctx, configs.ERROR_FORBIDDEN, nil, "非管理员账户禁止访问")
	//	return
	//}
	var filter serializers.MonitorFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	if len(filter.HostName) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "host_name can not null")
		return
	}
	if len(filter.Field) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "field can not empty")
		return
	}

	// usage_idle
	baseSql := "SELECT 100-mean(%s) AS usage FROM cpu where time > now() - %s AND host = '%s' group by time(%s)"
	SQL := fmt.Sprintf(baseSql, filter.Field[0], filter.DashboardTime, filter.HostName, filter.GetInterval())
	res, err := services.Service.Query(SQL)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, res, "")
	return
}

func GetSystemMemory(ctx *gin.Context) {
	//_, role := utils.GetUserIdAndRole(ctx)
	//if role != models.RAdmin{
	//	returnMsg(ctx, configs.ERROR_FORBIDDEN, nil, "非管理员账户禁止访问")
	//	return
	//}
	var filter serializers.MonitorFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	if len(filter.HostName) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "host_name can not null")
		return
	}
	if len(filter.Field) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "field can not empty")
		return
	}

	// used
	baseSql := "SELECT mean(%s)/1073741824 AS %s FROM mem where time > now() - %s AND host = '%s' group by time(%s)"
	SQL := fmt.Sprintf(baseSql, filter.Field[0], filter.Field[0], filter.DashboardTime, filter.HostName, filter.GetInterval())
	res, err := services.Service.Query(SQL)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, res, "")
	return
}

func GetSystemDisk(ctx *gin.Context) {
	//_, role := utils.GetUserIdAndRole(ctx)
	//if role != models.RAdmin{
	//	returnMsg(ctx, configs.ERROR_FORBIDDEN, nil, "非管理员账户禁止访问")
	//	return
	//}
	var filter serializers.MonitorFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	if len(filter.HostName) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "host_name can not null")
		return
	}
	if len(filter.Field) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "field can not empty")
		return
	}

	// used_percent
	baseSql := "SELECT mean(%s) AS %s FROM disk where time > now() - %s AND host = '%s' group by time(%s), path"
	SQL := fmt.Sprintf(baseSql, filter.Field[0], filter.Field[0], filter.DashboardTime, filter.HostName, filter.GetInterval())
	res, err := services.Service.Query(SQL)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, res, "")
	return
}

func GetSystemDiskio(ctx *gin.Context) {
	//_, role := utils.GetUserIdAndRole(ctx)
	//if role != models.RAdmin{
	//	returnMsg(ctx, configs.ERROR_FORBIDDEN, nil, "非管理员账户禁止访问")
	//	return
	//}
	var filter serializers.MonitorFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	if len(filter.HostName) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "host_name can not null")
		return
	}
	if len(filter.Field) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "field can not empty")
		return
	}

	// read_bytes write_bytes 系统硬盘读写速率
	baseSql := `SELECT non_negative_derivative(max(%s), 1s)/1000000 AS %s_per_second FROM diskio where time > now() - %s AND host = '%s' group by time(%s), "name"`
	SQL := fmt.Sprintf(baseSql, filter.Field[0], filter.Field[0], filter.DashboardTime, filter.HostName, filter.GetInterval())
	res, err := services.Service.Query(SQL)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}

	//xCoordinate := make([]interface{}, 0)
	//legend := make([]interface{}, 0)
	//series := make([]map[string]interface{}, 0)
	//for ri, s := range res {
	//	tmpLengend := make([]string, 0)
	//	tagName := s.Tags["name"]
	//	for m, c := range s.Columns {
	//		if m == 0 {
	//			//xCoordinate = append(xCoordinate, s.Values[m][0])
	//			continue
	//		}
	//		tmpLengend = append(tmpLengend, tagName + "-" + c)
	//		legend = append(legend, tagName + "-" + c)
	//	}
	//	tmp := make(map[string][]interface{})
	//	for _, v := range s.Values {
	//		if ri == 0{
	//			xCoordinate = append(xCoordinate, v[0])
	//		}
	//		for i, c1 := range tmpLengend {
	//			//if i == 0 {continue}
	//			//tmp[c1] = []interface{}{v[i + 1]}
	//			data, ok := tmp[c1]
	//			if !ok {
	//				tmp[c1] = []interface{}{v[i + 1]}
	//				continue
	//			}
	//			data = append(data, v[i + 1])
	//			tmp[c1] = data
	//		}
	//	}
	//	for k,v := range tmp{
	//		seriesItem := make(map[string]interface{})
	//		seriesItem["name"] = k
	//		seriesItem["data"] = v
	//		series = append(series, seriesItem)
	//	}
	//}
	//res1 := map[string]interface{}{
	//	"legend": legend,
	//	"xCoordinate": xCoordinate,
	//	"series": series,
	//}

	returnMsg(ctx, 200, res, "")
	return
}

func GetSystemLoad(ctx *gin.Context) {
	//_, role := utils.GetUserIdAndRole(ctx)
	//if role != models.RAdmin{
	//	returnMsg(ctx, configs.ERROR_FORBIDDEN, nil, "非管理员账户禁止访问")
	//	return
	//}
	var filter serializers.MonitorFilter
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, configs.ERROR_PARAMS, "", err.Error())
		return
	}
	if len(filter.HostName) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "host_name can not null")
		return
	}
	if len(filter.Field) == 0 {
		returnMsg(ctx, configs.ERROR_PARAMS, "", "field can not empty")
		return
	}

	// load1 系统负载
	baseSql := "SELECT mean(%s) AS %s FROM system where time > now() - %s AND host = '%s' group by time(%s)"
	SQL := fmt.Sprintf(baseSql, filter.Field[0], filter.Field[0], filter.DashboardTime, filter.HostName, filter.GetInterval())
	res, err := services.Service.Query(SQL)
	if err != nil {
		returnMsg(ctx, configs.ERROR_DATABASE, nil, err.Error())
		return
	}
	returnMsg(ctx, 200, res, "")
	return
}
