package lazada

import "encoding/json"

const (
	// APIGatewaySG endpoint
	APIGatewaySG = "https://api.lazada.sg/rest"
	// APIGatewayMY endpoint
	APIGatewayMY = "https://api.lazada.com.my/rest"
	// APIGatewayVN endpoint
	APIGatewayVN = "https://api.lazada.vn/rest"
	// APIGatewayTH endpoint
	APIGatewayTH = "https://api.lazada.co.th/rest"
	// APIGatewayPH endpoint
	APIGatewayPH = "https://api.lazada.com.ph/rest"
	// APIGatewayID endpoint
	APIGatewayID = "https://api.lazada.co.id/rest"

	APICODEURL = "https://api.lazada.com/rest"

	APIREFRESHURL = "https://auth.lazada.com/rest"

	AuthURL = "https://auth.lazada.com/oauth/authorize"
)

const (

	//根据授权码获取访问令牌、刷新令牌
	AccessTokenURL = "/auth/token/create"

	RefreshURL = "/auth/token/refresh"

	//卖家信息
	SELLERURL = "/seller/get"

	UPLOADIMAGE = "/image/upload"
)

type Response struct {
	Code      string          `json:"code"`
	Type      string          `json:"type"`
	Message   string          `json:"message"`
	RequestID string          `json:"request_id"`
	Data      json.RawMessage `json:"data"`
	Result    json.RawMessage `json:"result"`
}
