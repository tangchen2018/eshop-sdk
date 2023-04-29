package main

import (
	"fmt"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/tiktok"
)

func main() {
	api := tiktok.New(
		new(model.Setting).
			SetShopId("******").
			SetKey("******").
			SetSecret("******").
			SetAccessToken(`******`),
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
