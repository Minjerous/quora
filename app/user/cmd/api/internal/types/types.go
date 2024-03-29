// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id        int64  `json:"id"`
	UserName  string `json:"userName"`
	Avatar    string `json:"avatar"`
	Info      string `json:"info"`
	Following int64  `json:"following"`
	Follower  int64  `json:"follower"`
	Agreed    int64  `json:"agreed"`
	Liked     int64  `json:"liked"`
	Collected int64  `json:"collected"`
	Gender    int    `json:"gender"`
}

type RegisterReq struct {
	Email    string `form:"email" binding:"required"`
	UserName string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
	Code     string `form:"code"`
}

type RegisterResp struct {
	Status       int    `json:"status"`
	Msg          string `json:"msg"`
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type LoginReq struct {
	Email    string `form:"account"`
	Password string `form:"password"`
}

type LoginResp struct {
	Status       int    `json:"status"`
	Msg          string `json:"msg"`
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	Status   int    `json:"status"`
	Msg      string `json:"msg"`
	UserInfo User   `json:"userInfo"`
}

type OAuthReq struct {
	Code string `form:"code"`
}

type OAuthResp struct {
	Status       int    `json:"status"`
	Msg          string `json:"msg"`
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type SendEmailReq struct {
	Email string `form:"email"`
}

type SendEmailResp struct {
	VarifyCode string `json:"varify_code"`
}
