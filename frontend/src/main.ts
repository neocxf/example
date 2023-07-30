import { createApp } from "vue";

import App from "@/App.vue";

import '@/assets/icons' // icon
import '@/permission' // permission control
// import "@/styles/index.scss";
import "uno.css";

// If you want to use ElMessage, import it.
import "element-plus/theme-chalk/src/message.scss";

const app = createApp(App);

import store from '@/store'
import router from '@/router'
import directive from '@/directive' // directive
import plugins from '@/plugins' // plugins
app.use(store)
app.use(router)
app.use(directive)
app.use(plugins)

// 全局方法挂载
import { download } from '@/utils/request'
import { getDicts } from "@/api/system/dict/data";
import { getConfigKey } from "@/api/system/config";
import { parseTime, resetForm, addDateRange, selectDictLabel, selectDictLabels, handleTree } from "@/utils/ruoyi";
app.provide("getDicts", getDicts)
app.provide("getConfigKey", getConfigKey)
app.provide("parseTime", parseTime)
app.provide("resetForm", resetForm)
app.provide("addDateRange", addDateRange)
app.provide("selectDictLabel", selectDictLabel)
app.provide("selectDictLabels", selectDictLabels)
app.provide("download", download)
app.provide("handleTree", handleTree)

// 字典数据组件
import { install } from '@/components/DictData'
import DataDict from '@/utils/dict'

app.use(DataDict, install)

// import "@/styles/element/index.scss";

import Cookies from 'js-cookie'
import ElementPlus from "element-plus";
// import all element css, uncommented next line
// import "element-plus/dist/index.css";

// or use cdn, uncomment cdn link in `index.html`
app.use(ElementPlus, { size: 'small', zIndex: 3000 });

// app.use(ElementPlus, {
//     size: Cookies.get('size') || 'medium' // set element-ui default size
// })

app.mount("#app");
