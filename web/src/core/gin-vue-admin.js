/*
 * gin-vue-admin web框架组
 *
 * */
// 加载网站配置文件夹
import { register } from './global'

export default {
  install: (app) => {
    register(app)
    console.log(`
       欢迎使用 Platform-eff
       当前版本:v1.0.0
       Swagger UI地址:http://127.0.0.1:${import.meta.env.VITE_SERVER_PORT}/swagger/index.html
       前端登录访问地址:http://127.0.0.1:${import.meta.env.VITE_CLI_PORT}
    `)
  }
}
