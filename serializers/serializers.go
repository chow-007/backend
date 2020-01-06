package serializers

import "time"

type User struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"username, omitempty"`
	Password  string    `json:"-"`
	Email     string    `json:"email, omitempty"`
	Mobile    string    `json:"mobile, omitempty"`
	State     bool      `json:"state, omitempty"`
	CreatedAt time.Time `json:"created_at, omitempty"`
	UpdatedAt time.Time `json:"updated_at, omitempty"`
	Introduce string    `json:"introduce, omitempty"`
	RoleID    uint16    `json:"role_id, omitempty"`
}
