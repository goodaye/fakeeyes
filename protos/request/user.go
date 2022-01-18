package request

// 用户登陆
type UserLogin struct {
	Name string `json:"name"`
}

// 用户登陆
type UserSignIn struct {
	Name string `json:"name"`
}

// 用户注册
type UserSignUp struct {
	Name string `json:"name"`
}

//用户登陆检查头
type UserLoginHeader struct {
	Token string `json:"token"`
}
