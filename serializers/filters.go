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

// 额外的解析
//func (f *HistoryRequest) Parse() []string {
//	// 去重
//	return utils.SliceDuplicateRemoval(f.TagKeys)
//}

//// 额外的解析
//func (f *HistoryRequest) Parse() (topics []string, offset, limit int64, err error) {
//	//offset, err = strconv.ParseInt(f.Offset, 0, 0)
//	//limit, err = strconv.ParseInt(f.Limit, 0, 0)
//	//err = jsoniter.UnmarshalFromString(f.TagKeys, &topics)
//	topics = f.TagKeys
//
//	// 去重
//	//topics = strings.Split(f.TagKeys, ",")
//	topics = utils.SliceDuplicateRemoval(topics)
//
//	return
//}
