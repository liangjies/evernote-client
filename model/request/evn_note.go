package request

// 用户修改密码
type ChangePassword struct {
	OldPass string `json:"oldPass"` // 旧密码
	NewPass string `json:"newPass"` // 新密码
}
