// Package gtp
// @author tabuyos
// @since 2023/7/18
// @description gtp
package gtp

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"metis/model/page"
	"metis/model/reply"
	"metis/service/gtp"
	"metis/util/logger"
	"net/http"
)

// List 列表查询
func List(ctx *gin.Context) {
	pageInfo := page.New()
	gtpService := gtp.New()
	useLogger := logger.UseLogger()

	err := ctx.ShouldBindJSON(pageInfo)
	if err != nil {
		useLogger.Error(err.Error(), zap.Error(err))
		return
	}

	useLogger.Info(pageInfo.String())
	gtps := gtpService.FindAllWithPage(pageInfo)

	ctx.JSON(http.StatusOK, reply.OkData(gtps))
}

// Detail 详情
func Detail(ctx *gin.Context) {}

// Create 新增
func Create(ctx *gin.Context) {}

// Modify 修改
func Modify(ctx *gin.Context) {}

// Delete 删除
func Delete(ctx *gin.Context) {}
