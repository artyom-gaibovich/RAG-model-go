package dto

import "rag-model/app/domain/dao"

type GPTModel struct {
	ID          int    `gorm:"column:id; primary_key; not null" json:"id"`
	RequestText string `gorm:"column:request_text" json:"request_text"`
	dao.BaseModel
}
