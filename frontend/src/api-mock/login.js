import users from '@api-client/data/users.json'
import { getToken } from '@/utils/auth'
import { resolve } from 'path'

const fetch = (data, time = 0) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(data)
    }, time)
  })
}

// 登录方法
export const login = (username, password, code, uuid) => {
  return fetch(users, 100).then(_users => {
    let user = _users.find(element => element.username === username && element.password === password);
    console.log("user get", user)
    Promise.resolve(user.token)
  })
}

// 注册方法
export const register = (data) => {
  return request({
    url: '/register',
    headers: {
      isToken: false
    },
    method: 'post',
    data: data
  })
}

export const getInfo = () => {
  let _users = fetch(users, 100)

  let token = getToken()
  let user = _users.find(element => element.token === token);
  return {
    user: user
  }
}

// 退出方法
export const logout = () => {
  return request({
    url: '/logout',
    method: 'post'
  })
}

// 获取验证码
export const getCodeImg = () => {
  return request({
    url: '/captchaImage',
    headers: {
      isToken: false
    },
    method: 'get',
    timeout: 20000
  })
}
