package jwt

type Config struct {
	JwtAuth struct {
		AccessSecret string // 生成jwt token的密钥，最简单的方式可以使用一个uuid值
		AccessExpire int64  // jwt token有效期，单位：秒
	}
}
