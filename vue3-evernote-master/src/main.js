import { createApp } from 'vue'
import App from './App.vue'
import { store } from './store'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import '@/router/permission'
import router from '@/router/index'
// 统一导入el-icon图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'


const app = createApp(App)
// 统一注册el-icon图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.use(router)
app.use(store)
app.use(ElementPlus)
app.mount('#app')