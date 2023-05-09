package main

import (
	"encoding/json"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/store"
	"log"
)

func main() {
	s1 := store.New()
	s1.Listen()

	if err := s1.AddJob(&store.Token{
		Id: "",
		CallBack: func(e *store.Event) {
			tmp, _ := json.Marshal(&e.Token)
			log.Println(e.Success, e.Msg, string(tmp))
		},
		Refresh: store.Refresh{
			Key:               "",
			Secret:            "",
			PlatformCode:      model.PFC_TIKTOK,
			RefreshToken:      "",
			AccessTokenExpire: 1684159792,
		},
	}); err != nil {
		panic(err.Error())
	}

	<-make(chan bool)
}
