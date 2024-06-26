import Vue from 'vue'
import App from './App.vue'
// 路由
import router from './router'
// vuex
import store from './store'
// 引入iView和样式
import iView from 'iview';
import ViewUI from 'view-design'
import 'view-design/dist/styles/iview.css'
// 引入element-ui和样式
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// 引入正则表达式的工具
import Valid from './utils/valid'
// 引入DataV
import dataV from '@jiaminghi/data-view'

import './assets/public.css'

// 首页
import '@/styles/base.css'
import '@/styles/bootstrap.min.css'
import '@/styles/css2-Michromawght400500600700800_swap.css'
import '@/styles/css2-Montserratwght400500600700800_swap.css'
import '@/styles/font-awesome.min.css'

// 数据大屏
// 引入echarts
import * as echarts from 'echarts';
import './assets/less/index.less';
import img from './lib/img'
import utils from "./lib/utils";

// 使用iView
Vue.use(ViewUI);
Vue.use(iView);
// 使用element-ui
Vue.use(ElementUI);
// 使用DataV
Vue.use(dataV)
// 使用工具
Vue.use(utils)

// 把工具注入到vue实例中
Vue.prototype.$Valid = Valid;

// 设置echarts
Vue.prototype.$echarts = function (el) {
  return echarts.init(el, null, {renderer: 'svg'})
}
Vue.prototype.$images = img

Vue.config.productionTip = false

// 实现全局路由守卫
router.beforeEach((to, from, next) => {
	// 给网页标签名字
	if (to.meta.title) {
	  document.title = to.meta.title;
	}

  const token = store.state.userInfo.data.token;

  if (to.path === '/login') {
    // 如果目标路由是登录页，则直接放行
    next();
  } else if (!token) {
    // 如果token不存在，则重定向到登录页
    next('/login');
  } else {
    // 如果token存在，则正常放行
    next();
  }
})

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
