package dto

import "github.com/APSN4/RAG-model-go/app/domain/dao"

type GPTModel struct {
	ID           int      `gorm:"column:id; primary_key; not null" json:"id"`
	RequestTexts []string `gorm:"column:request_texts; type:string[]" json:"request_texts"`
	ModeGen      string   `gorm:"column:mode_gen; type:string" json:"mode_gen"`
	dao.BaseModel
}
