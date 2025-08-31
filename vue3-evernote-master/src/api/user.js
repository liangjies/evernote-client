import service from '@/utils/request'

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {
  return service({
    url: '/base/login',
    method: 'post',
    data: data
  })
}

// @Summary 用户退出登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const logout = (data) => {
  return service({
    url: '/user/logout',
    method: 'post',
    data: data
  })
}

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/resige [post]
export const register = (data) => {
  return service({
    url: '/base/register',
    method: 'post',
    data: data
  })
}

// @Summary 获取验证码
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/captcha [post]
export const captcha = (data) => {
  return service({
    url: '/base/captcha',
    method: 'post',
    data: data
  })
}

// @Summary 获取验证码
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/captcha [post]
export const verifyCode = (data) => {
    return service({
      url: '/base/sendVerifyCode',
      method: 'post',
      data: data
    })
  }

// @Summary 用户更新昵称
// @Produce  application/json
// @Param data
// @Router /user/nickName [patch]
export const UpdateNickName = (data) => {
  return service({
      url: '/user/nickName',
      method: 'patch',
      data
  })
}

// @Summary 用户修改密码
// @Produce  application/json
// @Param data
// @Router /user/changePassword [patch]
export const ChangePassword = (data) => {
  return service({
      url: '/user/changePassword',
      method: 'patch',
      data
  })
}


// @Summary 用户修改邮箱
// @Produce  application/json
// @Param data
// @Router /user/email [patch]
export const UpdateEmail = (data) => {
  return service({
      url: '/user/email',
      method: 'patch',
      data
  })
}