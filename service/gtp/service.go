// Package gtp
// @author tabuyos
// @since 2023/7/18
// @description gtp
package gtp

import (
	"metis/model/dto"
	"metis/model/page"
	"metis/repository/gtp"
)

type Service interface {
	FindAll() []dto.Gtp
	FindAllWithPage(pi *page.Info) []dto.Gtp
}
type service struct{}

func New() Service {
	return &service{}
}

func (receiver *service) FindAll() []dto.Gtp {
	panic("implement me")
}

func (receiver *service) FindAllWithPage(info *page.Info) []dto.Gtp {
	gtpRepository := gtp.New()
	return gtpRepository.SelectWithPage(info.Page, info.Size)
}
