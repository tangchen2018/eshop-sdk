package api

import (
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/pfc/tiktok"
)

type Api interface {
	RefreshToken(Body model.BodyMap) *model.Client
}

func New(pfc string, setting *model.Setting) Api {

	var (
		api Api
	)

	switch pfc {
	case model.PFC_TIKTOK:
		api = tiktok.New(setting)
	default:
	}
	return api
}
