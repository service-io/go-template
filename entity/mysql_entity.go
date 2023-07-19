// Package entity
// @author tabuyos
// @since 2023/6/30
// @description entity
package entity

import "time"

type CreateInfo struct {
	CreatorBy *int64     `json:"creatorBy" db_cmt:"创建人员 ID"`
	CreatedAt *time.Time `json:"createdAt" db_cmt:"创建时间"`
}

type UpdateInfo struct {
	UpdaterBy *int64     `json:"updaterBy" db_cmt:"更新人员 ID"`
	UpdatedAt *time.Time `json:"updatedAt" db_cmt:"更新时间"`
}

// Gtp gtp
type Gtp struct {
	Id          *int64     `json:"id" db_cmt:"主键"`
	Title       *string    `json:"title" db_cmt:"主题"`
	Description *string    `json:"description" db_cmt:"描述"`
	StartAt     *time.Time `json:"startAt" db_cmt:"开始时间"`
	EndAt       *time.Time `json:"endAt" db_cmt:"结束时间"`
	Status      *int8      `json:"status" db_cmt:"0: 发布 1: 暂存 2: 已结束 3: 已失效"`
	Top         *int8      `json:"top,omitempty" db_cmt:"0: 不置顶 1: 置顶"`
	CreateInfo
	UpdateInfo
}
