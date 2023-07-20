// Code generated by goctl. DO NOT EDIT.
package types

type RegistationReq struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Dob      int64  `json:"dob"`
}

type RegistationResp struct {
	Message string `json:"message"`
}

type LoginReq struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Message string `json:"message"`
	Token   string `json:"authToken"`
}
