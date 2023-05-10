package lazada

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

func (p *Api) GetAuthUrl(callbackParams string) string {

	return fmt.Sprintf("%s?%s", AuthURL, model.BodyMap{}.
		Set("response_type", "code").
		Set("client_id", *p.Setting.Key).
		Set("redirect_uri", fmt.Sprintf("%s?callbackParams=%s", *p.Setting.AuthCallbackUrl, callbackParams)).EncodeURLParams())
}

/*
	获取Token
	Url : https://open.lazada.com/apps/doc/api?path=%2Fauth%2Ftoken%2Fcreate
	Response: Response
*/

func (p *Api) GetToken(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(AccessTokenURL).
		//SetMethod(http.POST).
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("code"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	c.Client.Response.Response.Data = c.HttpReq.Result
	response := Response{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) GetSeller(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(SELLERURL).
		SetBody(Body)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetSellerResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	获取订单信息
	Url : https://open.lazada.com/apps/doc/api?path=%2Forder%2Fget
	Response: GetOrderDetailResponse
*/

func (p *Api) GetOrder(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/order/get`).
		//SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("order_id"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetOrderResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

/*
	获取订单明细
	Url : https://open.lazada.com/apps/doc/api?path=%2Forder%2Fitems%2Fget
	Response: GetOrderDetailResponse
*/

func (p *Api) GetOrderDetail(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`/order/items/get`).
		//SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("order_id"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := GetOrderDetailResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}
