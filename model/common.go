package model

import "encoding/json"

type Response struct {
	Success  bool `json:"success"`    // 是否成功
	Status   int  `json:"httpStatus"` // http状态码
	Response struct {
		RequestId string          `json:"requestId"` // 请求ID
		Code      string          `json:"code"`      // 平台code信息
		Message   string          `json:"message"`   // 平台msg信息
		Data      json.RawMessage `json:"data"`      // 数据
		DataTo    interface{}     `json:"dataTo"`    // 结构转化后的数据
	}
}

func (p *Response) ToMap() BodyMap {
	result := make(BodyMap)
	_ = json.Unmarshal(p.Response.Data, &result)
	return result
}

func (p *Response) To(row interface{}) error {
	return json.Unmarshal(p.Response.Data, row)
}
