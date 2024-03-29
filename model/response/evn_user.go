package response

import (
	"evernote-client/model"
)

type SysUserResponse struct {
	User model.EvnUser `json:"user"`
}

type LoginResponse struct {
	User      model.EvnUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
