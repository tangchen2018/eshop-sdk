package tiktok

import (
	"fmt"
	"github.com/tangchen2018/eshop-sdk/model"
)

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	return &Api{Setting: setting}
}

/*
	获取授权链接
	callbackParams : 同步回调的自定义参数
*/
func (p *Api) GetAuthUrl(callbackParams string) string {

	return fmt.Sprintf("%s%s?%s", AUTHSITE, AUTH, model.BodyMap{}.
		Set("app_key", *p.Setting.Key).
		Set("state", callbackParams).EncodeURLParams())
}

/*
	获取Token
	Url : https://bytedance.feishu.cn/docs/doccnROmkE6WI9zFeJuT3DQ3YOg
	Response: GetTokenResponse
*/

func (p *Api) GetToken(Body model.BodyMap) *Client {

	Body.Set("grant_type", "authorized_code")

	c := NewClient(p.Setting)
	c.SetPath(GETACCESS).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("auth_code"); c.Err != nil {
		return c
	}

	c.Execute()
	if c.Err != nil {
		return c
	}
	response := GetTokenResponse{}
	if err := c.Client.Response.To(&response); err != nil {
		return c
	}
	c.Response.Response.DataTo = response
	return c
}

/*
	获取订单列表
	Url : https://partner.tiktokshop.com/doc/page/262815?external_id=262815
	Response: GetOrderListResponse
*/

func (p *Api) GetOrderList(Body model.BodyMap) *Client {

	var cursor *string
	c := NewClient(p.Setting)
	c.SetPath(`/api/orders/search`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("page_size"); c.Err != nil {
		return c
	}

	result := GetOrderListResponse{}

	for {

		if cursor != nil && len(*cursor) > 0 {
			c.Request.Body.Set("cursor", cursor)
		}

		cResult := getOrderListResponse{}

		c.Execute()
		if c.Err = c.Client.Response.To(&cResult); c.Err != nil {
			return c
		}

		if cResult.OrderList != nil && len(cResult.OrderList) > 0 {
			for index, _ := range cResult.OrderList {
				result.List = append(result.List, cResult.OrderList[index])
			}
		}

		if cResult.More == false {
			result.Total = cResult.Total
			break
		} else {
			cursor = &cResult.NextCursor
		}
	}

	c.Response.Response.DataTo = result

	return c
}

/*
	获取订单明细
	Url : https://partner.tiktokshop.com/doc/page/262814?external_id=262814
	Response: GetOrderDetailResponse
*/

func (p *Api) GetOrderDetail(Body model.BodyMap) *Client {

	c := NewClient(p.Setting)
	c.SetPath(`/api/orders/detail/query`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("order_id_list"); c.Err != nil {
		return c
	}

	c.Execute()

	response := GetOrderDetailResponse{}
	if err := c.Client.Response.To(&response); err != nil {
		return c
	}
	c.Response.Response.DataTo = response
	return c
}
