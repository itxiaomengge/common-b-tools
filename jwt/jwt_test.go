package jwt

import (
	"testing"
	"fmt"
	"encoding/json"
	//"github.com/golang-jwt/jwt/v4"
	//"strings"
)

func Test_JWT(t *testing.T) {
	req := &GenerateTokenReq{
		UserId:   "12341243",
		UserName: "Xiaomeng.Ge",
	}

	res, _ := GenerateToken(req.UserId, req.UserName)

	json, _ := json.Marshal(res)

	fmt.Println(string(json))
}

////func Test_GetId(t *testing.T) {
////	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzIzMzQ5MTcsImlhdCI6MTY3MjI5MTcxNywiand0VXNlcklkIjoxMjM0MTI0Mywiand0VXNlck5hbWUiOiJYaWFvbWVuZy5HZSJ9.-6LpwIAT6iflxA8rLFlfvbC3i57mwoQmaBChlVqtoZM"
////
////	r, _ := http.NewRequest("GET", "", nil)
////	r.Header.Set("Authorization", token)
////
////	//ctx := new(context.Context)
////
////	//req := &UserInfoReq{}
////	req := &GenerateTokenReq{}
////	httpx.Parse(r, &req)
////
////	fmt.Println(r.Context().Value("Authorization"))
////
////	uid := GetUidFromCtx(r.Context())
////
////	fmt.Println(uid)
////}
//// CtxKeyJwtUserId get uid from ctx
////var CtxKeyJwtUserId = "jwtUserId"
////
////func GetUidFromCtx(ctx context.Context) int64 {
////	var uid int64
////	if jsonUid, ok := ctx.Value("jwtUserId").(json.Number); ok {
////		if int64Uid, err := jsonUid.Int64(); err == nil {
////			uid = int64Uid
////		} else {
////			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
////		}
////	}
////	return uid
////}
////
////type UserInfoReq struct {
////}
//
//func Test_DecodeSegment(t *testing.T) {
//	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzIzMzQ5MTcsImlhdCI6MTY3MjI5MTcxNywiand0VXNlcklkIjoxMjM0MTI0Mywiand0VXNlck5hbWUiOiJYaWFvbWVuZy5HZSJ9.-6LpwIAT6iflxA8rLFlfvbC3i57mwoQmaBChlVqtoZM"
//
//	arr := strings.Split(token, ".")
//
//	res, _ := jwt.DecodeSegment(arr[1])
//
//	fmt.Println(string(res))
//}
//
//func Test_ParseToken(t *testing.T) {
//	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzIzMzQ5MTcsImlhdCI6MTY3MjI5MTcxNywiand0VXNlcklkIjoxMjM0MTI0Mywiand0VXNlck5hbWUiOiJYaWFvbWVuZy5HZSJ9.-6LpwIAT6iflxA8rLFlfvbC3i57mwoQmaBChlVqtoZM"
//
//	userId := ParseToken(token)
//
//	fmt.Println(userId)
//}
