package model

type Tokens []*Token

type Token struct {
	Id                 string  `json:"id"`
	AccessToken        string  `json:"accessToken"`          // 访问令牌
	PlatformCode       string  `json:"platformCode"`         // 平台代码
	Setting            Setting `json:"setting"`              // 配置
	RefreshToken       string  `json:"refreshToken"`         // 刷新令牌
	AccessTokenExpire  int64   `json:"accessTokenExpireIn"`  // 访问令牌有效期(时间戳)
	RefreshTokenExpire int64   `json:"refreshTokenExpireIn"` // 刷新令牌有效期(时间戳)
	TaskStatus         string  `json:"taskStatus"`           // 任务执行状态 0-未创建任务 1-等待执行 2-执行中
}
