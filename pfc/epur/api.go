package epur

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

	return fmt.Sprintf("%s%s?%s", *p.Setting.ServerUrl, AUTHSITE, model.BodyMap{}.
		Set("client_id", *p.Setting.Key).
		Set("redirect_uri", fmt.Sprintf("%s?callbackParams=%s", *p.Setting.AuthCallbackUrl, callbackParams)).
		Set("scope", "agent_biz").
		Set("response_type", "code").EncodeURLParams(),
	)
}

/*
	获取Token
*/

func (p *Api) GetToken(Body model.BodyMap) *model.Client {

	Body.Set("grantType", "authorization_code").
		Set("clientId", *p.Setting.Key)

	c := NewClient(p.Setting)
	c.SetPath(ACCESSTOKEN).
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("code"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetTokenResponse{}
	if err := c.Client.Response.To(&response); err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) RefreshToken(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(REFRESHTOKEN).
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("refreshToken"); c.Err != nil {
		return &c.Client
	}

	c.Request.Body.Set("grantType", "refresh_token").
		Set("clientId", *p.Setting.Key)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	response := GetTokenResponse{}
	if err := c.Client.Response.To(&response); err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) PayApply(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(PAYAPPLY).
		SetBody(Body)

	//if c.Err = Body.CheckEmptyError("refreshToken"); c.Err != nil {
	//	return &c.Client
	//}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	//response := GetTokenResponse{}
	//if err := c.Client.Response.To(&response); err != nil {
	//	return &c.Client
	//}
	//c.Response.Response.DataTo = response
	return &c.Client
}

func (p *Api) GetMerchant(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(MERCHANT_INFO).
		SetBody(Body)

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}
	//response := GetTokenResponse{}
	//if err := c.Client.Response.To(&response); err != nil {
	//	return &c.Client
	//}
	//c.Response.Response.DataTo = response
	return &c.Client
}
