package jwt

type GenerateTokenReq struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

type GenerateTokenResp struct {
	AccessToken  string `json:"accessToken,omitempty" desc:"token值"`
	AccessExpire int64  `json:"accessExpire,omitempty" desc:"过期时间"`
	RefreshAfter int64  `json:"refreshAfter,omitempty" desc:"刷新时间"`
}
