import tab from './tab'
import auth from './auth'
import cache from './cache'
import modal from './modal'
import download from './download'

export const plugins = {
  install(app, options) {
    // configure the app
    // 页签操作
    app.provide('$tab', tab)
    // 认证对象
    app.provide('$auth', auth)
    // 缓存对象
    app.provide('$cache', cache)
    // 模态框对象
    app.provide('$modal', modal)
    // 下载文件
    app.provide('$download', download)
  }
}
