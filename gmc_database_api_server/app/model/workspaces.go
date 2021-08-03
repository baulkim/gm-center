package model

import (
	"time"
)

type Workspace struct {
	Num           int       `gorm:"column:workspaceNum; primary_key" json:"workspaceNum"`
	Name          string    `gorm:"column:workspaceName; not null; default:null" json:"workspaceName"`
	Description   string    `gorm:"column:workspaceDescription; not null; default:null" json:"workspaceDescription"`
	SelectCluster string    `gorm:"column:selectCluster; not null; default:null" json:"selectCluster"`
	Owner         string    `gorm:"column:workspaceOwner; not null; default:null" json:"workspaceOwner"`
	Creator       string    `gorm:"column:workspaceCreator; not null; default:null" json:"workspaceCreator"`
	Created_at    time.Time `gorm:"column:created_at" json:"created_at"`
}

// Set Cluster table name to be `CLUSTER_INFO`
func (Workspace) TableName() string {
	return "WORKSPACE_INFO"
}
