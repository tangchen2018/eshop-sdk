package main

import (
	"fmt"
	"github.com/tangchen2018/eshop-sdk/model"
	tiktok2 "github.com/tangchen2018/eshop-sdk/pfc/tiktok"
)

func main() {
	api := tiktok2.New(
		new(model.Setting).
			SetShopId("").
			SetKey("").
			SetSecret("").
			SetAccessToken(``),
	)
	GetOrderDetail(api)
}

func GetOrderDetail(api *tiktok2.Api) {
	/* 获取订单明细 */
	bm := model.BodyMap{}.
		Set("order_id_list", []string{"576604241027239433"})
	c := api.GetOrderDetail(bm)
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(tiktok2.GetOrderDetailResponse)
	fmt.Println(result)
}
