package lazada

import "github.com/tangchen2018/eshop-sdk/model"

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	return &Api{Setting: setting}
}

type Client struct {
	model.Client
}

func NewClient(setting *model.Setting) *Client {
	return &Client{Client: model.Client{Setting: setting}}
}

func (p *Client) Execute() {

}
