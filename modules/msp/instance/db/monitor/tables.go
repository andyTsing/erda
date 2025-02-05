// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package monitor

import (
	"reflect"
	"time"

	"github.com/erda-project/erda/pkg/database/gormutil"
)

// tables name
const (
	TableMonitor = "sp_monitor"
)

// Monitor .
type Monitor struct {
	Id                 int       `gorm:"column:id;primary_key"`
	MonitorId          string    `gorm:"column:monitor_id"`
	TerminusKey        string    `gorm:"column:terminus_key"`
	TerminusKeyRuntime string    `gorm:"column:terminus_key_runtime"`
	Workspace          string    `gorm:"column:workspace"`
	RuntimeId          string    `gorm:"column:runtime_id"`
	RuntimeName        string    `gorm:"column:runtime_name"`
	ApplicationId      string    `gorm:"column:application_id"`
	ApplicationName    string    `gorm:"column:application_name"`
	ProjectId          string    `gorm:"column:project_id"`
	ProjectName        string    `gorm:"column:project_name"`
	OrgId              string    `gorm:"column:org_id"`
	OrgName            string    `gorm:"column:org_name"`
	ClusterId          string    `gorm:"column:cluster_id"`
	ClusterName        string    `gorm:"column:cluster_name"`
	Config             string    `gorm:"column:config;default:''"`
	CallbackUrl        string    `gorm:"column:callback_url"`
	Version            string    `gorm:"column:version"`
	Plan               string    `gorm:"column:plan"`
	IsDelete           int       `gorm:"column:is_delete"`
	Created            time.Time `gorm:"column:created;default:CURRENT_TIMESTAMP"`
	Updated            time.Time `gorm:"column:updated;default:CURRENT_TIMESTAMP"`
}

func (Monitor) TableName() string { return TableMonitor }

var monitorFieldColumns = gormutil.GetFieldToColumnMap(reflect.TypeOf(Monitor{}))
