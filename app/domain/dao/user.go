package dao

type UserLogs struct {
	ID            int    `gorm:"column:id; primary_key; not null" json:"id"`
	UserName      string `gorm:"column:name" json:"user_name"`
	OperationType int    `gorm:"column:operation_type" json:"operation_type"`
	Status        int    `gorm:"column:status" json:"status"`
	RoleID        int    `gorm:"column:role_id;not null" json:"role_id"`
	Role          Role   `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	BaseModel
}
