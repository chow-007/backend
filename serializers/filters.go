package serializers

import (
	"fmt"
	"strconv"
	"strings"
)

type CreateUserFilter struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}

type UpdateUser struct {
	ID     uint   `json:"id" binding:"required"`
	State  bool   `json:"state"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	RoleID uint16 `json:"role_id"`
}

type DeleteRoleRight struct {
	RoleId  uint16 `json:"role_id" binding:"required"`
	RightId int    `json:"right_id" binding:"required"`
}

func (d *DeleteRoleRight) RightIdToStr() string {
	return strconv.Itoa(d.RightId)
}

type SetRightsFilter struct {
	Id    uint16 `json:"id" binding:"required"`
	PsIds string `json:"ps_ids" binding:"required"`
}

// 请求结构体
type HistoryRequest struct {
	Fields    []string `json:"fields" binding:"required"`
	StartTime int64    `json:"start_time" binding:"required"`
	EndTime   int64    `json:"end_time" binding:"required"`
	Offset    int64    `json:"offset"`
	Limit     int64    `json:"limit"`
	Period    string   `json:"period"`
	GroupFunc string   `json:"group_func"`
}

func (h *HistoryRequest) GetSelectFields() string {
	var safetyFields []string
	for _, f := range h.Fields {
		safetyFields = append(safetyFields, fmt.Sprintf("%s(%s) AS %s", h.GroupFunc, f, f))
	}
	return strings.Join(safetyFields, ", ")
}

type MonitorFilter struct {
	//ContainerIds  []string `json:"container_ids"` // 查询的容器
	Field         []string   `json:"field" binding:"required"`         // select查询字段
	DashboardTime string   `json:"dashboard_time" binding:"required"` // 查询时长
	HostName      string   `json:"host_name"`							// 主机名称
	//Interval      string   `json:"interval"`                          // 聚合粒度
}

//func (m *MonitorFilter) GetSafetySqlOr() string {
//	var safetyFields []string
//	for _, f := range m.ContainerIds {
//		safetyFields = append(safetyFields, fmt.Sprintf("container_id = '%s'", f))
//	}
//	return fmt.Sprintf("( %s )", strings.Join(safetyFields, " OR "))
//}

// 匹配聚合粒度
func (m *MonitorFilter) GetInterval() string {
	interval := map[string]string{
		"30d": "6h",
		"7d":  "1h",
		"2d":  "30m",
		"24h": "10m",
		"12h": "5m",
		"6h":  "1m",
		"1h":  "1m",
		"15m":  "1m",
		"5m":  "1m",
	}

	v, ok := interval[m.DashboardTime]
	if !ok {
		return "1m"
	}
	return v
}
