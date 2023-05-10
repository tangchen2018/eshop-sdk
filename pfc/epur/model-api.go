package epur

type GetTokenResponse struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	ClientId         string `json:"client_id"`
	Scope            string `json:"scope"`
	Openid           string `json:"openid"`
}
