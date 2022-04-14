package utils

var (
	NoteVerify      = Rules{"Title": {NotEmpty()}, "Content": {NotEmpty()}}
	NoteTitleVerify = Rules{"Title": {NotEmpty()}}
	NoteBookVerify  = Rules{"Title": {NotEmpty()}}
	LoginVerify     = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}

	IdVerify             = Rules{"ID": {NotEmpty()}}
	RegisterVerify       = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify       = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
)
