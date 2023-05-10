package main

import (
	"errors"
	"fmt"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/pfc/lazada"
)

func main() {
	api := lazada.New(
		new(model.Setting).
			SetKey("115018").
			SetSecret("4XYbxMIC5wMMHVz5bLvM6siY2GKMDIYk").
			SetAuthCallbackUrl("https://dev.api.epur1688.cn/admin/v1/erp/platform/cross/auth/callback/Lazada").
			SetAccessToken(`50000001501ba1b638f2eQwdjQe0qGWCex7rUVCGlJG0eQ1kh3AqTBMvXbzFfbUr`),
	)
	GetOrderDetail(api)
}

func GetAuthUrl(api *lazada.Api) {
	result := api.GetAuthUrl("123")
	fmt.Println(result)
}

func GetToken(api *lazada.Api) {
	c := api.GetToken(model.BodyMap{"code": "0_115018_97qwKzDdqQfKuD0tYArR3llr2652"})
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(lazada.Response)
	fmt.Println(result)
}

func GetSeller(api *lazada.Api) {

	api.Setting.SetSiteNo("th")

	c := api.GetSeller(nil)
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(lazada.GetSellerResponse)
	fmt.Println(result)
}

func GetOrderDetail(api *lazada.Api) {

	api.Setting.SetSiteNo("th")

	c := api.GetOrder(model.BodyMap{"order_id": "690646034119032"})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(lazada.GetOrderResponse)
	fmt.Println(result)
}
