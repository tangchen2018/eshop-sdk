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
		Key:                "687brqhjn0bao",
		Secret:             "650c1fc8c387c83d86916f597b346c9f846acbe4",
		PlatformCode:       model.PFC_TIKTOK,
		ShopId:             "8646921741465782268",
		RefreshToken:       "ROW_CyQm8wAAAADfyImmCET447SzLlILQ3Un3J532nmpelblKTCVoHUqzJWDKTDQ6Ln94Ais9tLd59k",
		AccessToken:        "ROW_4uJxIgAAAACOZVwk3CoFM_RfdPJTqI4lMg_n8uzRYr2OdYU3m55LkSGoemAWYfY4WUcQW0_cyqTkVKm6_uTgg7ZEQKmJyhZ6jwU2RwzA4vlnczD4f3wuyjNZ_JfRehv4yE4vs_v6KTlCXUekhWNAicT77s6TsohGJnJM-eaUGdtHXPnS34XTG1MAwcy-Oe71wu35Fb5g0yN0e6skj7R65RZ8tBXV2Qfl",
		AccessTokenExpire:  1684157615,
		RefreshTokenExpire: 1715053597,
	}); err != nil {
		panic(err)
	}

	req.Listen()
}
