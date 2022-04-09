package models

import (
	"github.com/merico-dev/lake/models/common"
	"time"
)

type TapdIssueCommit struct {
	SourceId        uint64     `gorm:"primaryKey"`
	ID              uint64     `gorm:"primaryKey;type:BIGINT(100)" json:"id"`
	WorkspaceId     uint64     `json:"workspace_id"`
	UserID          string     `gorm:"type:varchar(255)"`
	UserName        string     `gorm:"type:varchar(255)"`
	HookUserName    string     `gorm:"type:varchar(255)"`
	CommitID        string     `gorm:"type:varchar(255)"`
	Message         string     `gorm:"type:varchar(255)"`
	Path            string     `gorm:"type:varchar(255)"`
	WebURL          string     `gorm:"type:varchar(255)"`
	HookProjectName string     `gorm:"type:varchar(255)"`
	CommitTime      *time.Time `json:"commit_time"`
	Created         *time.Time `json:"created"`
	Ref             string     `gorm:"type:varchar(255)"`
	RefStatus       string     `gorm:"type:varchar(255)"`
	GitEnv          string     `gorm:"type:varchar(255)"`
	FileCommit      string     `gorm:"type:varchar(255)"`
	IssueId         uint64
	IssueType       string `gorm:"type:varchar(255)"`
	common.NoPKModel
}

type TapdIssueCommitApiRes struct {
	ID              string `json:"id"`
	UserName        string `json:"user_name"`
	UserID          string `json:"user_id"`
	HookUserName    string `json:"hook_user_name"`
	CommitID        string `json:"commit_id"`
	WorkspaceID     string `json:"workspace_id"`
	Message         string `json:"message"`
	Path            string `json:"path"`
	WebURL          string `json:"web_url"`
	HookProjectName string `json:"hook_project_name"`
	CommitTime      string `json:"commit_time"`
	Created         string `json:"created"`
	Ref             string `json:"ref"`
	RefStatus       string `json:"ref_status"`
	GitEnv          string `json:"git_env"`
	FileCommit      string `json:"file_commit"`
}

type IssueTypeAndId struct {
	Type    string
	IssueId uint64
}

func (TapdIssueCommit) TableName() string {
	return "_tool_tapd_issue_commits"
}
