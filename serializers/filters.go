package serializers

import (
	"strconv"
)

type CreateUserFilter struct {
	UserName  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
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
