package request

// 用户修改密码
type ChangePassword struct {
	OldPass string `json:"oldPass"` // 旧密码
	NewPass string `json:"newPass"` // 新密码
}

// 用户修改邮箱
type ChangeEmail struct {
	PassWord string `json:"password"` // 用户密码
	Email    string `json:"email"`    // 邮箱
}
