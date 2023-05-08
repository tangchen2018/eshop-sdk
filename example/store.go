package main

import (
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/store"
	"log"
)

func main() {
	req := store.New()

	req.CallBack = func(e *store.Event) {
		log.Println(string(e.ResponseData))
		log.Println(e.Token, e.Success, e.Msg)
	}

	if _, err := req.Save(&store.Request{
		Key:                "",
		Secret:             "",
		PlatformCode:       model.PFC_TIKTOK,
		ShopId:             "",
		RefreshToken:       "",
		AccessToken:        "",
		AccessTokenExpire:  1684157615,
		RefreshTokenExpire: 1715053597,
	}); err != nil {
		panic(err)
	}

	req.Listen()
}
