// Package gtp
// @author tabuyos
// @since 2023/7/19
// @description gtp
package gtp

import (
	"database/sql"
	"go.uber.org/zap"
	"metis/database"
	"metis/model/dto"
	"metis/util/logger"
)

type Repository interface {
	SelectWithPage(page uint, size uint) []dto.Gtp
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (receiver *repository) SelectWithPage(page uint, size uint) []dto.Gtp {
	var gtps []dto.Gtp
	useLogger := logger.UseLogger()
	db := database.FetchDB()
	prepare, _ := db.Prepare("select id, title, status, start_at from gtp limit ? offset ?")
	rows, err := prepare.Query(size, (page-1)*size)
	if err != nil {
		useLogger.Error(err.Error(), zap.Error(err))
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			useLogger.Error(err.Error(), zap.Error(err))
		}
	}(rows)

	for rows.Next() {
		var gtp dto.Gtp
		if err := rows.Scan(&gtp.Id, &gtp.Title, &gtp.Status, &gtp.StartAt); err != nil {
			useLogger.Error(err.Error(), zap.Error(err))
			return nil
		}
		gtps = append(gtps, gtp)
	}
	return gtps
}
