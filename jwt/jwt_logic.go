package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"encoding/json"
	"strings"
	_ "embed"
	"github.com/zeromicro/go-zero/core/conf"
)

// 将 jwt_config.yml 文件嵌入到编译的二进制文件中，这样打包成二进制文件时才能正常使用相对路径进行调用
//go:embed jwt_config.yml
var configFileContent []byte

// @Title GenerateToken
// @Description 生成一个 Token 值，并将其封装返回
// @Author Xiaomeng.Ge
// @Date 2022-12-28 19:57:17
//
// @Param req *GenerateTokenReq
//
// @return *result.Res
func GenerateToken(userId int64, userName string) (string, error) {
	now := time.Now().Unix()

	var c Config
	// 将 configFile 写入 c
	err := conf.LoadFromYamlBytes(configFileContent, &c)
	if err != nil {
		return "", err
	}
	accessSecret := c.JwtAuth.AccessSecret
	accessExpire := c.JwtAuth.AccessExpire
	accessToken, err := getJwtToken(accessSecret, now, accessExpire, userId, userName)

	if err != nil {
		return "", err
	}

	return accessToken, err
}

// @Title getJwtToken
// @Description 得到一个 Token 值
// @Author Xiaomeng.Ge
// @Date 2022-12-28 19:57:25
//
// @Param secretKey string
// @Param iat int64
// @Param seconds int64
// @Param userId int64
// @Param userName string
//
// @return string
// @return error
func getJwtToken(secretKey string, iat int64, seconds int64, userId int64, userName string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["jwtUserId"] = userId
	claims["jwtUserName"] = userName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// @Title ParseToken
// @Description 解析 token 值，得到用户 id
// @Author Xiaomeng.Ge
// @Date 2022-12-29 19:18:39
//
// @Param tokenString string
//
// @return int64
func ParseToken(tokenString string) (int64, error) {
	arr := strings.Split(tokenString, ".")

	payload, err := jwt.DecodeSegment(arr[1])
	if err != nil {
		return 0, err
	}

	claims := make(jwt.MapClaims, 0)
	// 或
	// claims := make(map[string]interface{}, 0)
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return 0, err
	}

	userId := int64(claims["jwtUserId"].(float64))
	return userId, nil
}
