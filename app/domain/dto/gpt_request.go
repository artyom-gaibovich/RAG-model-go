package dto

import "rag-model/app/domain/dao"

type GPTModel struct {
	ID           int      `gorm:"column:id; primary_key; not null" json:"id"`
	RequestTexts []string `gorm:"column:request_texts; type:string[]" json:"request_texts"`
	dao.BaseModel
}
