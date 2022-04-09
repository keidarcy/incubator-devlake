package models

import "github.com/merico-dev/lake/models/common"

type TapdUser struct {
	SourceId    uint64 `gorm:"primaryKey;type:BIGINT(20)"`
	WorkspaceId uint64 `gorm:"primaryKey;type:BIGINT(20)"`
	Name        string `gorm:"index;type:varchar(255)"`
	User        string `gorm:"primaryKey;type:varchar(255)"`
	common.NoPKModel
}

type TapdUserApiRes struct {
	User string `json:"user"`
	Name string `json:"name"`
}

func (TapdUser) TableName() string {
	return "_tool_tapd_users"
}
