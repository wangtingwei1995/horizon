package models

import (
	"g.hz.netease.com/horizon/pkg/server/global"
)

type EnvironmentRegion struct {
	global.Model

	EnvironmentName string
	RegionName      string
	Disabled        bool
	CreatedBy       uint
	UpdatedBy       uint
}
