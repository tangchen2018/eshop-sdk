package main

import (
	"fmt"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/tiktok"
)

func main() {
	api := tiktok.New(
		new(model.Setting).
			SetShopId("8646921741465782268").
			SetKey("687brqhjn0bao").
			SetSecret("650c1fc8c387c83d86916f597b346c9f846acbe4").
			SetAccessToken(`ROW_qo_MqwAAAACOZVwk3CoFM_RfdPJTqI4lMg_n8uzRYr2OdYU3m55LkSGoemAWYfY4WUcQW0_cyqTkVKm6_uTgg7ZEQKmJyhZ6jwU2RwzA4vlnczD4f3wuyiCLWYLN7zINIZdjSQlBCMykILTI22Eo-rw4cCyA20Sy5FG7Yg3gqJcn-xYnKBX84Sov2mSO-t_FEt_iGBLD4ZFodcXZRNCW0ytCvLY706KZ`),
	)
	GetOrderDetail(api)
}

func GetOrderDetail(api *tiktok.Api) {
	/* 获取订单明细 */
	bm := model.BodyMap{}.
		Set("order_id_list", []string{"576604241027239433"})
	c := api.GetOrderDetail(bm)
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(tiktok.GetOrderDetailResponse)
	fmt.Println(result)
}
