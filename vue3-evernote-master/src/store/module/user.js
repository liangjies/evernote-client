import { login, logout } from '@/api/user.js'
import router from '@/router/index.js'

export const user = {
  namespaced: true,
  state: {
    userInfo: {
      uuid: '',
      nickName: '',
      userName: '',
      headerImg: '',
      email: '',
      createdAt:'',
    },
    token: ''
  },
  mutations: {
    setEmail(state, email) {
      state.userInfo.email = email
    },
    setHeaderImg(state, headerImg) {
      state.userInfo.headerImg = headerImg
    },
    setNickName(state, nickName) {
      state.userInfo.nickName = nickName
    },
    setUserInfo(state, userInfo) {
      // 这里的 `state` 对象是模块的局部状态
      state.userInfo = userInfo
    },
    setToken(state, token) {
      // 这里的 `state` 对象是模块的局部状态
      state.token = token
    },
    LoginOut(state) {
      state.userInfo = {}
      state.token = ''
      router.push({ name: 'Login', replace: true })
      // window.location.reload()
    },
  },
  actions: {
    async LoginIn({ commit, dispatch, rootGetters, getters }, loginInfo) {
      const res = await login(loginInfo)
      if (res.code === 200) {
        await commit('setUserInfo', res.data.user)
        await commit('setToken', res.data.token)
        console.log(res.data.token)
        router.push({ name: "NoteDetail", params: { id: 0 } });
        return true
      }else{
        return res
      }
    },
    async LoginOut({ commit }) {
      const res = await logout()
      if (res.code === 200) {
        commit('LoginOut')
      }
    },
  },
  getters: {
    userInfo(state) {
      return state.userInfo
    },
    token(state) {
      return state.token
    },
  },
}
